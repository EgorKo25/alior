package amqp

import (
	"callback_service/src/repository"
	"callback_service/src/service"
	"callback_service/src/transport"
	"context"
	"encoding/json"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Consume(amqpURL, queueName string, svc *service.CallbackService) error {
	conn, err := transport.ConnectToRabbitMQ(amqpURL)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close connection: %v", err)
		}
	}()

	ch, err := transport.SetupChannel(conn)
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

			ctx := context.Background()
			err := svc.CreateCallback(ctx, msg.Number, msg.Date.Format(time.RFC3339), msg.Name)
			if err != nil {
				log.Printf("Failed to create callback: %v", err)
			} else {
				err = Produce(amqpURL, "notify", "new callback")

				if err != nil {
					log.Printf("Failed to send notification: %v", err)
				}
			}
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
