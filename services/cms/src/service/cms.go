package service

import (
	"callback_service/src/database"
	"context"
	"encoding/json"
)

type IBroker interface {
	Consume(ctx context.Context, queueName string, handler func(context.Context, []byte) error) error
	Produce(ctx context.Context, queueName string, body []byte) error
}

type ICallback interface {
	CreateCallback(ctx context.Context, data *database.Callback) error
}

type ILogger interface {
	Error(msg string, args ...interface{})
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
	return c.broker.Consume(ctx, "create", c.HandleMessage)
}

func (c *CMS) HandleMessage(ctx context.Context, body []byte) error {
	callback, err := ConvertToRepositoryAndValidate(body)
	if err != nil {
		c.logger.Error("error during validation or conversion: %s", err.Error())
		return c.broker.Produce(ctx, "error", []byte(err.Error()))
	}

	err = c.storage.CreateCallback(ctx, callback)
	if err != nil {
		c.logger.Error("error inserting callback: %s", err.Error())
		return c.broker.Produce(ctx, "error", []byte(err.Error()))
	}

	successMsg := "Callback created successfully"
	return c.broker.Produce(ctx, "success", []byte(successMsg))
}

func ConvertToRepositoryAndValidate(callbackSrc []byte) (*database.Callback, error) {
	var callback database.Callback
	err := json.Unmarshal(callbackSrc, &callback)
	if err != nil {
		return nil, err
	}
	return &callback, nil
}
