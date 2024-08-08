package broker

import (
	"context"
	"log"

	"github.com/streadway/amqp"
)

type AMQPSession struct {
	*amqp.Connection
	*amqp.Channel
	exchangeName string
	exchangeType string
	routingKey   string
	queueName    string
}

func (c *AMQPSession) Close() error {
	if c.Connection == nil {
		return nil
	}

	return c.Connection.Close()
}

// Запускает горутину, которая постоянно пытается создать новое соединение и канал, возвращая их в виде сессии через канал.
func Redial(ctx context.Context, uri, exchangeName, exchangeType, routingKey, queueName string) chan chan AMQPSession {
	// Создаем канал для отправки каналов сессий
	sessions := make(chan chan AMQPSession)

	go func() {
		// Создаем канал для сессии
		session := make(chan AMQPSession)

		defer close(sessions)

		for {
			select {
			case sessions <- session: // Отправляем канал текущей сессии через канал sessions
			case <-ctx.Done(): // Проверяем, не завершен ли контекст
				log.Println("Shutting down session factory")
				return // Выходим из горутины, завершая её работу
			}

			// Пытаемся установить соединение с RabbitMQ
			conn, err := amqp.Dial(uri)
			if err != nil {
				log.Fatalf("cannot (re)dial: %v: %q", err, uri)
			}

			// Пытаемся создать канал
			ch, err := conn.Channel()
			if err != nil {
				log.Fatalf("cannot create channel: %v", err)
			}

			// Объявляем эксченж
			if err := ch.ExchangeDeclare(
				exchangeName, // Название эксченжа
				exchangeType, // Тип эксченжа
				false,        // durable
				true,         // autodelete
				false,        // internal
				false,        // noWait
				nil); err != nil {
				log.Fatalf("cannot declare fanout exchange: %v", err)
			}

			select {
			case session <- AMQPSession{conn, ch, exchangeName, exchangeType, routingKey, queueName}: // Отправляем текущую сессию в канал
			case <-ctx.Done(): // Если контекст завершен, выходим из горутины
				log.Println("shutting down new session")
				return
			}
		}
	}()

	return sessions
}

// Отправляет сообщения в RabbitMQ с возможностью автоматического восстановления соединения в случае его разрыва.
// Эта функция работает с сессиями, предоставляемыми функцией redial, и использует подтверждения публикации, чтобы гарантировать доставку сообщений.
func Publish(sessions chan chan AMQPSession, messages <-chan []byte) {
	for session := range sessions { // Получаем новую сессию от канала sessions
		var (
			running bool                              // Есть ли что отправить
			reading = messages                        // Канал из которого читаем что отправить
			pending = make(chan []byte, 1)            // Буфер
			confirm = make(chan amqp.Confirmation, 1) // Канал для получения подтверждений об отправке от RabbitMQ
		)

		pub := <-session // Получаем текущую сессию (соединение и канал)

		// Включаем подтверждения публикации для текущего канала
		if err := pub.Confirm(false); err != nil {
			log.Printf("publisher confirms not supported")
			close(confirm) // Если подтверждения не поддерживаются, закрываем канал confirm
		} else {
			pub.NotifyPublish(confirm) // Включаем получение подтверждений публикации
		}

	Publish:
		for {
			var body []byte
			select {
			case confirmed, ok := <-confirm:
				if !ok {
					break Publish // Если канал confirm закрыт, выходим из цикла и переходим к следующей сессии
				}
				if !confirmed.Ack {
					log.Printf("nack message %d, body: %q", confirmed.DeliveryTag, string(body))
				}
				reading = messages // Возобновляем чтение новых сообщений

			case body = <-pending:
				err := pub.Publish(pub.exchangeName, pub.routingKey, false, false, amqp.Publishing{
					Body: body,
				})
				// Если публикация не удалась, повторяем попытку на следующей сессии
				if err != nil {
					pending <- body
					pub.Close()
					break Publish
				}

			case body, running = <-reading:
				if !running {
					return // Если больше нет сообщений для отправки, завершаем работу функции
				}
				pending <- body // Отправляем сообщение в канал pending для публикации
				reading = nil   // Останавливаем чтение новых сообщений, пока текущее не будет опубликовано
			}
		}
	}
}

// Подписывается на сообщения из RabbitMQ.
// Эта функция создаёт очередь для потребителя, связывает её с эксченжем, и затем начинает получать сообщения из этой очереди.
// Все полученные сообщения отправляются в канал messages, чтобы их можно было дальше обрабатывать.
func Subscribe(sessions chan chan AMQPSession, messages chan<- []byte) {
	for session := range sessions {
		// Когда сессия становится доступной, она извлекается из канала
		sub := <-session
		// Создаётся очередь
		if _, err := sub.QueueDeclare(sub.queueName,
			false, // durable
			true,  // autodelete
			true,  // exclusive
			false, // noWait
			nil); err != nil {
			log.Printf("cannot consume from queue: %q, %v", sub.queueName, err)
			return
		}

		// Связывание очереди с эксчежем
		if err := sub.QueueBind(sub.queueName, sub.routingKey, sub.exchangeName, false, nil); err != nil {
			log.Printf("cannot consume without a binding to exchange: %q, %v", sub.exchangeName, err)
			return
		}

		// Консюмит сообщения из очереди
		deliveries, err := sub.Consume(sub.queueName,
			"",    // consumer
			false, // autoAck
			true,  // exclusive
			false, // noLocal
			false, // noWait
			nil)
		if err != nil {
			log.Printf("cannot consume from: %q, %v", sub.queueName, err)
			return
		}

		log.Printf("subscribed...")

		// цикл обрабатывает каждое сообщение, приходящее в очередь
		for msg := range deliveries {
			messages <- msg.Body // тело сообщения передаётся в канал messages для дальнейшей обработки

			err := sub.Ack(msg.DeliveryTag, false) // сообщение подтверждается как обработанное, чтобы оно не было доставлено другим консюмерам

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
