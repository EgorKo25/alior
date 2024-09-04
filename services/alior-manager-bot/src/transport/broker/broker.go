package broker

import (
	"context"
	"errors"
	"github.com/EgorKo25/common/broker"
	"github.com/EgorKo25/common/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IBroker interface {
	Produce(body string, msgType string) error
	Consume(ctx context.Context) (*amqp.Delivery, error)
}

type Broker struct {
	broker *broker.Broker
	logger logger.ILogger
}

func NewBroker(uri, exchangeName, exchangeKind, routingKey, queueName string, logger logger.ILogger) (*Broker, error) {
	brokerConfig := broker.NewDialConfig(uri, exchangeName, exchangeKind, routingKey, queueName)

	_, err := broker.Init(brokerConfig, logger)
	if err != nil {
		logger.Error("failed to initialize broker: %s", err)
		return nil, err
	}

	b, err := broker.GetBroker()
	if err != nil {
		logger.Error("failed to get broker: %s", err)
		return nil, err
	}

	return &Broker{
		broker: b,
		logger: logger,
	}, nil
}

func (b *Broker) Produce(body string, msgType string) error {
	message := NewMessage(body, msgType)

	err := b.broker.Publish(amqp.Publishing{
		ContentType: message.Properties.ContentType,
		Type:        message.Properties.Type,
		Body:        []byte(message.Body),
	})
	if err != nil {
		b.logger.Error("failed to publish message: %s", err)
		return err
	}
	return nil
}

func (b *Broker) Consume(ctx context.Context) (*amqp.Delivery, error) {
	delivery, err := b.broker.Consume()
	if err != nil {
		b.logger.Error("failed to consume message: %s", err)
		return nil, err
	}

	for {
		select {
		case msg := <-delivery:
			if msg.ContentType == "callback" {
				return &msg, nil
			}
			b.logger.Warn("skipped message with ContentType: %s", msg.ContentType)
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				return nil, errors.New("exit with timeout")
			}
			return nil, ctx.Err()
		}
	}
}
