package broker

import (
	"context"
	"log"
	"time"

	"github.com/streadway/amqp"
)

//go:generate go run github.com/vektra/mockery/v2@v2.44.1 --name IConnection
type IConnection interface {
	Channel() (*amqp.Channel, error)
	Close() error
}

//go:generate go run github.com/vektra/mockery/v2@v2.44.1 --name IChannel
type IChannel interface {
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	Confirm(noWait bool) error
	NotifyPublish(confirm chan amqp.Confirmation) chan amqp.Confirmation
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	Ack(tag uint64, multiple bool) error
	Close() error
}

type Session struct {
	Connection   IConnection
	Channel      IChannel
	ExchangeName string
	ExchangeKind string
	RoutingKey   string
	QueueName    string
}

func (c *Session) Close() error {
	if c.Connection == nil {
		return nil
	}
	return c.Connection.Close()
}

func Redial(ctx context.Context, uri, exchangeName, exchangeKind, routingKey, queueName string) chan Session {
	sessionChan := make(chan Session)

	go func() {
		defer close(sessionChan)

		for {
			select {
			case <-ctx.Done():
				log.Println("Shutting down session factory")
				return
			default:
				// Пытаемся установить соединение с RabbitMQ
				conn, err := amqp.Dial(uri)
				if err != nil {
					log.Fatalf("cannot (re)dial: %v: %q", err, uri)
				}

				// Пытаемся создать канал
				ch, err := conn.Channel()
				if err != nil {
					conn.Close()
					log.Fatalf("cannot create channel: %v", err)
				}

				// Объявляем exchange
				if err := ch.ExchangeDeclare(
					exchangeName,
					exchangeKind, // Тип exchange
					false,        // флаг durable
					true,         // флаг autoDelete
					false,        // флаг internal
					false,        // флаг noWait
					nil); err != nil {
					log.Printf("cannot declare exchange: %v", err)
					ch.Close()
					conn.Close()
					time.Sleep(time.Second) // Задержка перед повторной попыткой
					continue
				}

				// Отправляем текущую сессию в канал
				sessionChan <- Session{
					Connection:   conn,
					Channel:      ch,
					ExchangeName: exchangeName,
					ExchangeKind: exchangeKind,
					RoutingKey:   routingKey,
					QueueName:    queueName,
				}
			}
		}
	}()

	return sessionChan
}

func Publish(sessionChan chan Session, messages <-chan []byte) {
	for session := range sessionChan {
		var (
			running bool                              // Флаг, указывающий, есть ли сообщения для отправки
			reading = messages                        // Канал, из которого читаются сообщения
			pending = make(chan []byte, 1)            // Буфер для хранения сообщений, ожидающих отправки
			confirm = make(chan amqp.Confirmation, 1) // Канал для получения подтверждений публикации от RabbitMQ
			body    []byte                            // Переменная для хранения текущего сообщения
		)

		// Включаем подтверждения публикации для текущего канала
		if err := session.Channel.Confirm(false); err != nil {
			log.Printf("publisher confirms not supported: %v", err)
			close(confirm) // Если подтверждения не поддерживаются, закрываем канал confirm
		} else {
			session.Channel.NotifyPublish(confirm) // Включаем получение подтверждений публикации
		}

	PublishLoop:
		for {
			select {
			case confirmed, ok := <-confirm:
				if !ok {
					break PublishLoop // Если канал confirm закрыт, выходим из цикла и переходим к следующей сессии
				}
				if !confirmed.Ack {
					log.Printf("nack message %d", confirmed.DeliveryTag)
				}
				reading = messages // Возобновляем чтение новых сообщений

			case body := <-pending:
				err := session.Channel.Publish(session.ExchangeName, session.RoutingKey, false, false, amqp.Publishing{
					Body: body,
				})
				if err != nil {
					pending <- body // Если публикация не удалась, сохраняем сообщение в буфере и выходим
					session.Connection.Close()
					break PublishLoop
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

func Subscribe(sessionChan chan Session, messages chan<- []byte) {
	for session := range sessionChan {
		// Создаем очередь
		_, err := session.Channel.QueueDeclare(
			session.QueueName,
			false, // durable
			true,  // autoDelete
			true,  // exclusive
			false, // noWait
			nil)   // args
		if err != nil {
			log.Printf("cannot declare queue: %q, %v", session.QueueName, err)
			continue // Переходим к следующей сессии в случае ошибки
		}

		// Привязываем очередь к exchange
		err = session.Channel.QueueBind(
			session.QueueName,
			session.RoutingKey,
			session.ExchangeName,
			false, // noWait
			nil)   // args
		if err != nil {
			log.Printf("cannot bind queue: %q to exchange: %q, %v", session.QueueName, session.ExchangeName, err)
			continue // Переходим к следующей сессии в случае ошибки
		}

		// Подписываемся на сообщения из очереди
		deliveries, err := session.Channel.Consume(
			session.QueueName,
			"",    // consumerTag
			false, // autoAck
			true,  // exclusive
			false, // noLocal
			false, // noWait
			nil)   // args
		if err != nil {
			log.Printf("cannot consume from queue: %q, %v", session.QueueName, err)
			continue // Переходим к следующей сессии в случае ошибки
		}

		log.Printf("subscribed to queue: %q...", session.QueueName)

		// Обрабатываем сообщения из очереди
		for msg := range deliveries {
			messages <- msg.Body // Отправляем тело сообщения в канал messages

			// Подтверждаем получение сообщения
			err := session.Channel.Ack(msg.DeliveryTag, false)
			if err != nil {
				log.Printf("cannot ack message: %v", err)
			}
		}
	}
}
