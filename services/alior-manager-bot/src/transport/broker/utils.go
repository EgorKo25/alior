package broker

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func (b *Broker) CMSMessageExchange(ctx context.Context, body string, msgType string) (*amqp.Delivery, error) {
	err := b.Produce(body, msgType)
	if err != nil {
		b.logger.Error("failed to produce request", "error", err)
		return nil, err
	}
	b.logger.Info("successfully produced request: %s", body)

	time.Sleep(20 * time.Second)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	msg, err := b.Consume(ctx)
	if err != nil {
		b.logger.Error("failed to consume message", "error", err)
		return nil, err
	}

	switch msg.Type {
	case "success":
		b.logger.Info("Initial callback response", "response", msg)
		return msg, nil
	case "error":
		b.logger.Warn("Failed to get initial callback response", "error", msg)
		return msg, errors.New("failed to get initial callback response")
	default:
		return nil, errors.New("unknown response type")
	}
}
