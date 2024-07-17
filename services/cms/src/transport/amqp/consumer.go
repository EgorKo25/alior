package amqp

import (
	"callback_service/src/repository"
	"callback_service/src/service"
	"callback_service/src/transport"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func Consume(ctx context.Context, amqpURL, queueName string, svc *service.CallbackService) error {
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
			log.Printf("Received a message: %s", d.Body)

			var msg repository.Callback
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			if err := svc.CreateCallback(ctx, msg.Name, msg.Phone, msg.Type, msg.Idea); err != nil {
				log.Printf("Failed to create callback: %v", err)
			}

			if err := Produce(ctx, amqpURL, "notify", "new callback"); err != nil {
				log.Printf("Failed to send notification: %v", err)
			}
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
