package broker

import (
	"context"
	"errors"
	"fmt"

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

//go:generate go run github.com/vektra/mockery/v2@v2.44.1 --name ILogger
type ILogger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

type QueueConfig struct {
	QueueName  string     // Название очереди
	Durable    bool       // очередь будет сохраняться при перезапуске
	AutoDelete bool       // очередь не будет удалена автоматически
	Exclusive  bool       // очередь доступна для других соединений
	NoWait     bool       // ждать подтверждения от сервера
	Args       amqp.Table // дополнительные аргументы
}

type ConsumeConfig struct {
	Consumer  string     // Идентификатор потребителя (consumer tag), который будет использоваться для идентификации данного потребителя.
	AutoAck   bool       // Автоматическое подтверждение сообщений; если true, сообщения будут автоматически подтверждены при получении.
	Exclusive bool       // Эксклюзивный доступ к очереди; если true, очередь будет доступна только этому потребителю.
	NoLocal   bool       // Локальная подписка; если true, то сообщения, опубликованные в этом же соединении, не будут доставляться этому потребителю.
	NoWait    bool       // Не ожидать подтверждения от сервера RabbitMQ; если true, функция вернется немедленно.
	Args      amqp.Table // Дополнительные аргументы для настройки подписки (например, для реализации различных расширений).
}

type ExchangeConfig struct {
	ExchangeName string     // Имя Exchange, к которому подключаемся
	ExchangeKind string     // Тип Exchange, к которому подключаемся (fanout, direct, topic, headers)
	RoutingKey   string     // routing key ¯\_(ツ)_/¯
	Durable      bool       // сохранять очередь после перезапуска RabbitMQ
	AutoDelete   bool       // автоматически удалять очередь, если нет потребителей
	Internal     bool       // доступ к очереди имеет только текущее соединение
	NoWait       bool       // не ждать подтверждения от сервера RabbitMQ
	Args         amqp.Table // дополнительные аргументы для настройки очереди
}

type DialConfig struct {
	URI           string          // Ссылка на подключение к RabbitMQ
	Connect2Queue bool            // 1, если будем подключаться к очереди; 0, если только к exchange (если нужен только publish)
	ExchangeCfg   *ExchangeConfig // Конфиг для Exchange
	QueueCfg      *QueueConfig    // Конфиг для Queue
	ConsumeCfg    *ConsumeConfig  // Конфиг для consume
}

// Создает DialConfig c параметрами по умолчанию.
func NewDialConfig(uri, exchangeName, exchangeKind, routingKey, queueName string, connect2Queue bool) *DialConfig {
	// Устанавливаем значения по умолчанию для Exchange и Queue конфигураций.
	exchangeCfg := ExchangeConfig{
		ExchangeName: exchangeName, // Имя Exchange по умолчанию
		ExchangeKind: exchangeKind, // Тип Exchange по умолчанию
		RoutingKey:   routingKey,   // Пустой routing key по умолчанию
		Durable:      false,        // Exchange не сохраняется после перезапуска по умолчанию
		AutoDelete:   true,         // Exchange удаляется, если нет потребителей по умолчанию
		Internal:     false,        // Exchange доступен для других соединений по умолчанию
		NoWait:       false,        // Ожидаем подтверждения от сервера по умолчанию
		Args:         nil,          // Нет дополнительных аргументов по умолчанию
	}

	queueCfg := QueueConfig{
		QueueName:  queueName, // Имя очереди по умолчанию
		Durable:    false,     // Очередь не сохраняется после перезапуска по умолчанию
		AutoDelete: true,      // Очередь удаляется, если нет потребителей по умолчанию
		Exclusive:  false,     // Очередь доступна для других соединений по умолчанию
		NoWait:     false,     // Ожидаем подтверждения от сервера по умолчанию
		Args:       nil,       // Нет дополнительных аргументов по умолчанию
	}

	consumeCfg := ConsumeConfig{
		Consumer:  "",    // Идентификатор потребителя (consumer tag). Пустая строка означает, что RabbitMQ сгенерирует его автоматически.
		AutoAck:   false, // Автоматическое подтверждение сообщений. false означает, что сообщения должны подтверждаться вручную.
		Exclusive: false, // Эксклюзивный доступ к очереди. false позволяет другим потребителям подписываться на ту же очередь.
		NoLocal:   false, // Локальная подписка. false означает, что потребитель может получать сообщения, опубликованные в том же соединении.
		NoWait:    false, // Ожидание подтверждения от сервера. false означает, что потребитель будет ждать подтверждения от RabbitMQ.
		Args:      nil,   // Дополнительные аргументы для настройки подписки. nil означает, что дополнительных параметров нет.
	}

	return &DialConfig{
		URI:           uri,
		Connect2Queue: connect2Queue,
		ExchangeCfg:   &exchangeCfg,
		QueueCfg:      &queueCfg,
		ConsumeCfg:    &consumeCfg,
	}
}

type Session struct {
	Connection IConnection // Интерфейс подключения к RabbitMQ
	Channel    IChannel    // Канал для общения с RabbitMQ
	Cfg        DialConfig  // Конфиг для подключения
}

func (c *Session) Close() error {
	if c.Connection == nil {
		return nil
	}

	return c.Connection.Close()
}

// URI           string - Ссылка на подключение к RabbitMQ;
// ExchangeName  string - Имя Exchange, к которому подключаемся;
// ExchangeKind  string - Тип Exchange, к которому подключаемся (fanout, direct, topic, headers);
// RoutingKey    string - routing key для binding очереди к exchange;
// Connect2Queue bool   - 1, если будем подключаться к очереди; 0, если только к exchange (если нужен только publish);
// QueueName     string - Название очереди;.

// Функция, которая создает и поддерживает подключение к брокеру, согласно конфигу
// Если в конфиге Connect2Queue = 0, то подключится только к exchange, если 1, то еще и к очереди.
func DialSessionChan(ctx context.Context, Dconfig DialConfig, logger ILogger) (chan Session, error) {
	sessionChan := make(chan Session)
	errChan := make(chan error, 1)

	go func() {
		defer func() {
			close(sessionChan)
			close(errChan)
		}()

		for {
			select {
			case <-ctx.Done():
				logger.Info("Shutting down session factory")
				errChan <- errors.New("context done")

				return
			default:
				// Пытаемся установить соединение с RabbitMQ
				conn, err := amqp.Dial(Dconfig.URI)
				if err != nil {
					logger.Error(fmt.Sprintf("cannot (re)dial: %v: %q", err, Dconfig.URI))
					errChan <- err // Отправляем ошибку в канал и завершаем горутину

					return
				}

				// Пытаемся создать канал
				ch, err := conn.Channel()
				if err != nil {
					conn.Close()
					logger.Error(fmt.Sprintf("cannot create channel: %v", err))
					errChan <- err // Отправляем ошибку в канал и завершаем горутину

					return
				}

				// Объявляем exchange
				if err := ch.ExchangeDeclare(
					Dconfig.ExchangeCfg.ExchangeName,
					Dconfig.ExchangeCfg.ExchangeKind, // Тип exchange
					Dconfig.ExchangeCfg.Durable,      // флаг durable
					Dconfig.ExchangeCfg.AutoDelete,   // флаг autoDelete
					Dconfig.ExchangeCfg.Internal,     // флаг internal
					Dconfig.ExchangeCfg.NoWait,       // флаг noWait
					Dconfig.ExchangeCfg.Args); err != nil {
					ch.Close()
					conn.Close()
					logger.Error(fmt.Sprintf("cannot declare exchange: %v", err))
					errChan <- err // Отправляем ошибку в канал и завершаем горутину

					return
				}

				if Dconfig.Connect2Queue {
					if Dconfig.QueueCfg.QueueName == "" {
						logger.Error("empty queue name")
						errChan <- errors.New("empty queue name")

						return
					}

					// Создаем очередь
					_, err = ch.QueueDeclare(
						Dconfig.QueueCfg.QueueName,
						Dconfig.QueueCfg.Durable,    // durable
						Dconfig.QueueCfg.AutoDelete, // autoDelete
						Dconfig.QueueCfg.Exclusive,  // exclusive
						Dconfig.QueueCfg.NoWait,     // noWait
						Dconfig.QueueCfg.Args)       // args
					if err != nil {
						ch.Close()
						conn.Close()
						logger.Error(fmt.Sprintf("cannot declare queue: %q, %v", Dconfig.QueueCfg.QueueName, err))
						errChan <- err // Отправляем ошибку в канал и завершаем горутину

						return
					}

					// Привязываем очередь к exchange
					err = ch.QueueBind(
						Dconfig.QueueCfg.QueueName,
						Dconfig.ExchangeCfg.RoutingKey,
						Dconfig.ExchangeCfg.ExchangeName,
						Dconfig.QueueCfg.NoWait, // noWait
						Dconfig.QueueCfg.Args)   // args
					if err != nil {
						ch.Close()
						conn.Close()
						logger.Error(fmt.Sprintf("cannot bind queue: %q to exchange: %q, %v", Dconfig.QueueCfg.QueueName, Dconfig.ExchangeCfg.ExchangeName, err))
						errChan <- err // Отправляем ошибку в канал и завершаем горутину

						return
					}
				}

				// Отправляем текущую сессию в канал
				sessionChan <- Session{
					Connection: conn,
					Channel:    ch,
					Cfg:        Dconfig,
				}
			}
		}
	}()

	// Ожидаем либо получение сессий, либо ошибку
	select {
	case err := <-errChan:
		return nil, err
	case <-sessionChan:
		return sessionChan, nil
	}
}

// Постоянно принимает сообщения из messages <-chan []byte
// Отдает их на exchange, к которому подключена сессия chan Session.
func Publish(sessionChan chan Session, messages <-chan amqp.Publishing, logger ILogger) {
	for session := range sessionChan {
		var (
			reading = messages                        // Канал, из которого читаются сообщения
			pending = make(chan amqp.Publishing, 1)   // Буфер для хранения сообщений, ожидающих отправки
			confirm = make(chan amqp.Confirmation, 1) // Канал для получения подтверждений публикации от RabbitMQ
		)

		// Включаем подтверждения публикации для текущего канала
		if err := session.Channel.Confirm(false); err != nil {
			logger.Info(fmt.Sprintf("publisher confirms not supported: %v", err))
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
					logger.Warn(fmt.Sprintf("nack message %d", confirmed.DeliveryTag))
				}
				reading = messages // Возобновляем чтение новых сообщений

			case msg := <-pending:
				err := session.Channel.Publish(
					session.Cfg.ExchangeCfg.ExchangeName,
					session.Cfg.ExchangeCfg.RoutingKey,
					true,
					false,
					msg,
				)
				if err != nil {
					pending <- msg // Если публикация не удалась, сохраняем сообщение в буфере и выходим
					session.Connection.Close()
					break PublishLoop
				}

			case msg, running := <-reading:
				if !running {
					return // Если больше нет сообщений для отправки, завершаем работу функции
				}
				pending <- msg // Отправляем сообщение в канал pending для публикации
				reading = nil  // Останавливаем чтение новых сообщений, пока текущее не будет опубликовано
			}
		}
	}
}

// Постоянно принимает сообщения из очереди, к которой подключена сессия chan Session
// Отдает их в канал messages chan<- []byte.
func Subscribe(sessionChan chan Session, messages chan<- amqp.Delivery, logger ILogger) {
	for session := range sessionChan {
		if !session.Cfg.Connect2Queue {
			logger.Info("Session config Connect2Queue is false")
			continue
		}
		// Подписываемся на сообщения из очереди
		deliveries, err := session.Channel.Consume(
			session.Cfg.QueueCfg.QueueName,
			session.Cfg.ConsumeCfg.Consumer,  // consumerTag
			session.Cfg.ConsumeCfg.AutoAck,   // autoAck
			session.Cfg.ConsumeCfg.Exclusive, // exclusive
			session.Cfg.ConsumeCfg.NoLocal,   // noLocal
			session.Cfg.ConsumeCfg.NoWait,    // noWait
			session.Cfg.ConsumeCfg.Args)      // args
		if err != nil {
			logger.Error(fmt.Sprintf("cannot consume from queue: %q, %v", session.Cfg.QueueCfg.QueueName, err))
			continue // Переходим к следующей сессии в случае ошибки
		}

		logger.Info(fmt.Sprintf("subscribed to queue: %q...", session.Cfg.QueueCfg.QueueName))
		// Обрабатываем сообщения из очереди
		for msg := range deliveries {
			messages <- msg // Отправляем тело сообщения в канал messages
			// Подтверждаем получение сообщения
			err := session.Channel.Ack(msg.DeliveryTag, false)
			if err != nil {
				logger.Error(fmt.Sprintf("cannot ack message: %v", err))
			}
		}
	}
}
