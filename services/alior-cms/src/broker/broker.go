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

type IBroker interface {
	Publish(message *Message) error
	Subscribe(ctx context.Context, queue string, handler func(ctx context.Context, delivery amqp.Delivery) error) error
	Close()
}

type ConnectionManager struct {
	Url    string
	Logger ILogger
	Conn   *amqp.Connection
}

func (cm *ConnectionManager) Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial(cm.Url)
	if err != nil {
		cm.Logger.Error("failed to connect to AMQP broker: %s", err)
		return nil, err
	}
	cm.Conn = conn
	cm.Logger.Info("connected to AMQP broker")
	return conn, nil
}

func (cm *ConnectionManager) Close() error {
	if cm.Conn != nil {
		err := cm.Conn.Close()
		if err != nil {
			cm.Logger.Error("failed to close AMQP connection: %s", err)
			return err
		}
		cm.Logger.Info("AMQP connection closed")
	}
	return nil
}

type ChannelManager struct {
	Logger ILogger
}

func (cm *ChannelManager) CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()
	if err != nil {
		cm.Logger.Error("failed to create AMQP channel: %s", err)
		return nil, err
	}
	cm.Logger.Info("AMQP channel created")
	return channel, nil
}

func (cm *ChannelManager) Close(channel *amqp.Channel) error {
	if channel != nil {
		err := channel.Close()
		if err != nil {
			cm.Logger.Error("failed to close AMQP channel: %s", err)
			return err
		}
		cm.Logger.Info("AMQP channel closed")
	}
	return nil
}

type Broker struct {
	ConnManager    IConnectionManager
	ChannelManager IChannelManager
	Conn           *amqp.Connection
	Channel        *amqp.Channel
	Logger         ILogger
}

func NewBroker(Url string, logger ILogger) (*Broker, error) {
	connManager := &ConnectionManager{
		Url:    Url,
		Logger: logger,
	}

	channelManager := &ChannelManager{
		Logger: logger,
	}

	conn, err := connManager.Connect()
	if err != nil {
		return nil, err
	}

	channel, err := channelManager.CreateChannel(conn)
	if err != nil {
		err := connManager.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &Broker{
		Conn:           conn,
		Channel:        channel,
		Logger:         logger,
		ConnManager:    connManager,
		ChannelManager: channelManager,
	}, nil
}

func (b *Broker) Close() {
	if err := b.ChannelManager.Close(b.Channel); err != nil {
		b.Logger.Error("failed to close channel: %v", err)
	}
	if err := b.ConnManager.Close(); err != nil {
		b.Logger.Error("failed to close connection: %v", err)
	}
}

func (b *Broker) Publish(message *Message) error {
	return b.Channel.Publish(
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
	messages, err := b.Channel.Consume(
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
		for m := range messages {
			if err := handler(ctx, m); err != nil {
				b.Logger.Error("error handling message: %v", err)
			}
		}
	}()

	return nil
}
