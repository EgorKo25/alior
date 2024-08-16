package broker

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func DeclareQueue(ch *amqp.Channel, queueName string) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func CreateResponseMessage(delivery amqp.Delivery, responseBody string, responseType string) *RabbitMQMessage {
	msg := NewMessage(responseBody, delivery)
	msg.Properties.Type = responseType
	msg.Properties.Timestamp = time.Now()

	return msg
}
