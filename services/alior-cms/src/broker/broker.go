package broker

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ILogger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type IConnectionManager interface {
	Connect() (*amqp.Connection, error)
	Close() error
}

type IChannelManager interface {
	CreateChannel(conn *amqp.Connection) (*amqp.Channel, error)
	Close(channel *amqp.Channel) error
}

type ConnectionManager struct {
	Url    string
	logger ILogger
	conn   *amqp.Connection
}

func (cm *ConnectionManager) Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial(cm.Url)
	if err != nil {
		cm.logger.Error("failed to connect to AMQP broker", "error", err)
		return nil, err
	}
	cm.conn = conn
	cm.logger.Info("connected to AMQP broker")
	return conn, nil
}

func (cm *ConnectionManager) Close() error {
	if cm.conn != nil {
		err := cm.conn.Close()
		if err != nil {
			cm.logger.Error("failed to close AMQP connection", "error", err)
			return err
		}
		cm.logger.Info("AMQP connection closed")
	}
	return nil
}

type ChannelManager struct {
	logger ILogger
}

func (cm *ChannelManager) CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()
	if err != nil {
		cm.logger.Error("failed to create AMQP channel", "error", err)
		return nil, err
	}
	cm.logger.Info("AMQP channel created")
	return channel, nil
}

func (cm *ChannelManager) Close(channel *amqp.Channel) error {
	if channel != nil {
		err := channel.Close()
		if err != nil {
			cm.logger.Error("failed to close AMQP channel", "error", err)
			return err
		}
		cm.logger.Info("AMQP channel closed")
	}
	return nil
}

type Broker struct {
	connManager    IConnectionManager
	channelManager IChannelManager
	conn           *amqp.Connection
	channel        *amqp.Channel
	logger         ILogger
}

func NewBroker(Url string, logger ILogger) (*Broker, error) {
	connManager := &ConnectionManager{
		Url:    Url,
		logger: logger,
	}

	channelManager := &ChannelManager{
		logger: logger,
	}

	conn, err := connManager.Connect()
	if err != nil {
		return nil, err
	}

	channel, err := channelManager.CreateChannel(conn)
	if err != nil {
		connManager.Close()
		return nil, err
	}

	return &Broker{
		conn:           conn,
		channel:        channel,
		logger:         logger,
		connManager:    connManager,
		channelManager: channelManager,
	}, nil
}

func (b *Broker) Close() {
	if err := b.channelManager.Close(b.channel); err != nil {
		b.logger.Error("failed to close channel: %v", err)
	}
	if err := b.connManager.Close(); err != nil {
		b.logger.Error("failed to close connection: %v", err)
	}
}

func (b *Broker) Publish(message *Message) error {
	return b.channel.Publish(
		message.Headers.Exchange,   // exchange
		message.Headers.RoutingKey, // routing key
		false,                      // mandatory
		false,                      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message.Body),
		})
}

func (b *Broker) Subscribe(ctx context.Context, queue string, handler func(ctx context.Context, delivery amqp.Delivery) error) error {
	msgs, err := b.channel.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	go func() {
		for m := range msgs {
			if err := handler(ctx, m); err != nil {
				b.logger.Error("error handling message: %v", err)
			}
		}
	}()

	return nil
}
