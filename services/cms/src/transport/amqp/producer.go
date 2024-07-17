package amqp

import (
	"callback_service/src/transport"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func Produce(ctx context.Context, amqpURL, queueName string, body string) error {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close connection: %v", err)
		}
	}()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			log.Printf("Failed to close channel: %v", err)
		}
	}()

	q, err := transport.DeclareQueue(ch, queueName)
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
		log.Printf("Failed to produce notification")
	}

	log.Printf(" [x] Sent %s\n", body)

	return nil
}
