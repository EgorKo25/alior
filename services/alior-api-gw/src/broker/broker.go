package broker

import (
	"github.com/streadway/amqp"
)

type (
	IAmqpMessage interface {
		Create() (*amqp.Publishing, error)
	}
	Broker interface {
		Publish(msg IAmqpMessage)
	}
)
