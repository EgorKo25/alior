package service

import (
	"callback_service/src/broker"
	"callback_service/src/database"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IBroker interface {
	Subscribe(ctx context.Context, queue string, handler func(ctx context.Context, delivery amqp.Delivery) error) error
	Publish(message *broker.Message) error
}

type ICallback interface {
	CreateCallback(ctx context.Context, data *database.Callback) error
	GetCallback(ctx context.Context, limit int, offset int) (callback *database.Callback, err error)
	DeleteCallbackByID(ctx context.Context, id int32) error
	GetTotalCallbacks(ctx context.Context) (int, error)
}

type ILogger interface {
	Error(msg string, args ...interface{})
	Info(msg string, args ...interface{})
}

type CMS struct {
	Broker  IBroker
	Storage ICallback
	Logger  ILogger
}

func NewCMS(Broker IBroker, Storage ICallback, Logger ILogger) *CMS {
	return &CMS{
		Broker:  Broker,
		Storage: Storage,
		Logger:  Logger,
	}
}

func (c *CMS) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		if err := c.Broker.Subscribe(ctx, "ask", c.HandleMessage); err != nil {
			errCh <- err
		}
	}()

	return <-errCh
}
