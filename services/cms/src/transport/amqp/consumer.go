package amqp

import (
	"callback_service/src/logger"
	"callback_service/src/repository"
	"callback_service/src/service"
	"callback_service/src/transport"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IConsumer interface {
	Consume(ctx context.Context) error
}

type Consumer struct {
	amqp    string
	queue   string
	service *service.CallbackService
	logger  logger.ILogger
}

func NewConsumer(amqpURL, queueName string, svc *service.CallbackService, logger logger.ILogger) *Consumer {
	return &Consumer{
		amqp:    amqpURL,
		queue:   queueName,
		service: svc,
		logger:  logger,
	}
}

func (c Consumer) Consume(ctx context.Context) error {
	conn, err := amqp.Dial(c.amqp)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			c.logger.Error("Failed to close connection: %v", err)
		}
	}()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			c.logger.Error("Failed to close channel: %v", err)
		}
	}()

	q, err := transport.DeclareQueue(ch, c.queue)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			c.logger.Info("Received a message: %s", d.Body)

			var msg repository.Callback
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				c.logger.Error("failed to unmarshal message: %v", err)
				continue
			}

			if err := c.service.CreateCallback(ctx, msg.Name, msg.Phone, msg.Type, msg.Idea); err != nil {
				c.logger.Error("failed to create callback: %v", err)
			}

			producer := NewProducer(c.amqp, c.queue, c.logger)
			if err := producer.Produce(ctx, "new callback"); err != nil {
				c.logger.Error("failed to send notification: %v", err)
			}
		}
	}()

	c.logger.Info("Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
