package service

import (
	"callback_service/src/broker"
	"callback_service/src/database"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

// IBroker declare methods to interact with broker queues
type IBroker interface {
	Subscribe(ctx context.Context, queue string, handler func(ctx context.Context, delivery amqp.Delivery) error) error
	Publish(message *broker.Message) error
}

// ICallback declare interaction methods for Callback structure
type ICallback interface {
	CreateCallback(ctx context.Context, data *database.Callback) error
	GetCallback(ctx context.Context, limit int, offset int) (callback *database.Callback, err error)
	DeleteCallbackByID(ctx context.Context, id int32) error
	GetTotalCallbacks(ctx context.Context) (int, error)
}

// ILogger interface declares methods for Logger struct
type ILogger interface {
	Error(msg string, args ...interface{})
	Info(msg string, args ...interface{})
}

// CMS is a structure to store broker, db and logger instances
type CMS struct {
	Broker  IBroker
	Storage ICallback
	Logger  ILogger
}

// NewCMS is a constructor for CMS
func NewCMS(Broker IBroker, Storage ICallback, Logger ILogger) *CMS {
	return &CMS{
		Broker:  Broker,
		Storage: Storage,
		Logger:  Logger,
	}
}

// Run is a CMS method to run subscriber
func (c *CMS) Run(ctx context.Context) error {
	errCh := make(chan error, 1)

	go func() {
		if err := c.Broker.Subscribe(ctx, "ask", c.HandleMessage); err != nil {
			errCh <- err
		}
	}()

	return <-errCh
}
