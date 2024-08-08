package service

import (
	"callback_service/src/broker"
	"callback_service/src/database"
	"callback_service/src/logger"
	"context"
	"encoding/json"
	"fmt"
)

type CMS struct {
	broker  broker.IBroker
	storage database.ICallback
	logger  logger.ILogger
}

func NewCMS(broker broker.IBroker, storage database.ICallback, logger logger.ILogger) *CMS {
	return &CMS{
		broker:  broker,
		storage: storage,
		logger:  logger,
	}
}

func (c *CMS) Run(ctx context.Context) error {
	return c.broker.Consume(ctx, "create", c.handleMessage)
}

func (c *CMS) handleMessage(ctx context.Context, body []byte) error {
	callback, err := convertToRepositoryAndValidate(body)
	if err != nil {
		c.logger.Error("Error during validation or conversion: %s", err.Error())
		return c.broker.Produce(ctx, "error", []byte(err.Error()))
	}

	err = c.storage.CreateCallback(ctx, callback)
	if err != nil {
		c.logger.Error("Error inserting callback: %s", err.Error())
		return c.broker.Produce(ctx, "error", []byte(err.Error()))
	}

	successMsg := fmt.Sprintf("Callback created successfully")
	return c.broker.Produce(ctx, "success", []byte(successMsg))
}

func convertToRepositoryAndValidate(callbackSrc []byte) (database.Callback, error) {
	var callback database.Callback
	err := json.Unmarshal(callbackSrc, &callback)
	if err != nil {
		return database.Callback{}, err
	}
	return callback, nil
}
