package transport

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectToRabbitMQ(amqpURL string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func SetupChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func DeclareQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return q, err
	}
	return q, nil
}
