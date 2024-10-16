package broker

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

// ILogger local logger declare
type ILogger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// IConnectionManager to manage broker connection
type IConnectionManager interface {
	Connect() (*amqp.Connection, error)
	Close() error
}

// IChannelManager to manage broker channel
type IChannelManager interface {
	CreateChannel(conn *amqp.Connection) (*amqp.Channel, error)
	Close(channel *amqp.Channel) error
}

// IBroker to manage broker operations
type IBroker interface {
	Publish(message *Message) error
	Subscribe(ctx context.Context, queue string, handler func(ctx context.Context, delivery amqp.Delivery) error) error
	Close()
}

// ConnectionManager structure to store connection and URL to connect
type ConnectionManager struct {
	URL    string
	Logger ILogger
	Conn   *amqp.Connection
}

// Connect is a ConnectionManager method to connect broker
func (cm *ConnectionManager) Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial(cm.URL)
	if err != nil {
		cm.Logger.Error("failed to connect to AMQP broker: %s", err)
		return nil, err
	}
	cm.Conn = conn
	cm.Logger.Info("connected to AMQP broker")
	return conn, nil
}

// Close is a ConnectionManager method to close broker connection
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

// ChannelManager structure manage broker channel
type ChannelManager struct {
	Logger ILogger
}

// CreateChannel is a ChannelManager method to create broker channel
func (cm *ChannelManager) CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	channel, err := conn.Channel()
	if err != nil {
		cm.Logger.Error("failed to create AMQP channel: %s", err)
		return nil, err
	}
	cm.Logger.Info("AMQP channel created")
	return channel, nil
}

// Close is a ChannelManager method to close broker channel
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

// Broker structure to store connection, channel and their managers
type Broker struct {
	ConnManager    IConnectionManager
	ChannelManager IChannelManager
	Conn           *amqp.Connection
	Channel        *amqp.Channel
	Logger         ILogger
}

// NewBroker is a Broker constructor
func NewBroker(URL string, logger ILogger) (*Broker, error) {
	connManager := &ConnectionManager{
		URL:    URL,
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

// Close is a Broker structure method to close channel and connections
func (b *Broker) Close() {
	if err := b.ChannelManager.Close(b.Channel); err != nil {
		b.Logger.Error("failed to close channel: %v", err)
	}
	if err := b.ConnManager.Close(); err != nil {
		b.Logger.Error("failed to close connection: %v", err)
	}
}

// Publish is a Broker structure method to publish message to broker
func (b *Broker) Publish(message *Message) error {
	return b.Channel.Publish(
		message.Headers.Exchange,
		message.Headers.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message.Body),
			Headers: amqp.Table{
				"msg_type": message.Properties.Type,
			},
		})
}

// Subscribe is a Broker structure method to get messages from broker
func (b *Broker) Subscribe(ctx context.Context, queue string, handler func(ctx context.Context, delivery amqp.Delivery) error) error {
	messages, err := b.Channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
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
