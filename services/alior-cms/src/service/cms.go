package service

import (
	"callback_service/src/database"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IBroker interface {
	Consume(ctx context.Context, queueName string, handler func(context.Context, amqp.Delivery) error) error
	Produce(ctx context.Context, delivery amqp.Delivery, responseBody string, responseType string) error
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
	broker  IBroker
	storage ICallback
	logger  ILogger
}

func NewCMS(broker IBroker, storage ICallback, logger ILogger) *CMS {
	return &CMS{
		broker:  broker,
		storage: storage,
		logger:  logger,
	}
}

func (c *CMS) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		if err := c.broker.Consume(ctx, "ask", c.HandleMessage); err != nil {
			errCh <- err
		}
	}()

	return <-errCh
}
