package broker

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ILogger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type Broker struct {
	Url    string
	logger ILogger
}

func NewBroker(Url string, logger ILogger) *Broker {
	return &Broker{
		Url:    Url,
		logger: logger,
	}
}

func (b *Broker) Consume(ctx context.Context, queueName string, handler func(context.Context, []byte) error) error {
	conn, err := amqp.Dial(b.Url)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			b.logger.Error("failed to close connection: %v", err)
		}
	}()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			b.logger.Error("failed to close channel: %v", err)
		}
	}()

	q, err := DeclareQueue(ch, queueName)
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

	go func() {
		for d := range msgs {
			b.logger.Info("Received a message: %s", d.Body)

			if err := handler(ctx, d.Body); err != nil {
				b.logger.Error("failed to handle message: %v", err)
			}
		}
	}()

	b.logger.Info("Waiting for messages. To exit press CTRL+C")
	<-ctx.Done()

	return nil
}

func (b *Broker) Produce(ctx context.Context, queueName string, body []byte) error {
	conn, err := amqp.Dial(b.Url)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			b.logger.Error("Failed to close connection: %v", err)
		}
	}()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			b.logger.Error("Failed to close channel: %v", err)
		}
	}()

	q, err := DeclareQueue(ch, queueName)
	if err != nil {
		return err
	}

	ctxWTO, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctxWTO,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		b.logger.Error("Failed to produce message: %v", err)
		return err
	}

	b.logger.Info("Sent message: %s", body)

	return nil
}
