package amqp

import (
	"callback_service/src/logger"
	"callback_service/src/transport"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type IProducer interface {
	Produce(ctx context.Context, body string) error
}

type Producer struct {
	amqp   string
	queue  string
	logger logger.ILogger
}

func NewProducer(amqpURL, queueName string, logger logger.ILogger) *Producer {
	return &Producer{
		amqp:   amqpURL,
		queue:  queueName,
		logger: logger,
	}
}

func (p *Producer) Produce(ctx context.Context, body string) error {
	conn, err := amqp.Dial(p.amqp)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			p.logger.Error("failed to close connection: %v", err)
		}
	}()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			p.logger.Error("failed to close channel: %v", err)
		}
	}()

	q, err := transport.DeclareQueue(ch, p.queue)
	if err != nil {
		return err
	}

	ctxWTO, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctxWTO,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		p.logger.Error("failed to produce notification")
	}

	p.logger.Info(" [x] Sent %s\n", body)

	return nil
}
