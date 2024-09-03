package broker

import (
	"errors"
	"github.com/EgorKo25/common/broker"
	"github.com/EgorKo25/common/logger"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

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
		ContentType: "callback",
		Body:        []byte(message.Body),
	})
	if err != nil {
		b.logger.Error("failed to publish message: %s", err)
		return err
	}
	return nil
}

func (b *Broker) Consume(timeout time.Duration) (string, error) {
	delivery, err := b.broker.Consume()
	if err != nil {
		b.logger.Error("failed to consume message: %s", err)
		return "", err
	}

	select {
	case msg := <-delivery:
		return string(msg.Body), nil
	case <-time.After(timeout):
		return "", errors.New("exit with timeout")
	}
}
