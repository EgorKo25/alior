package service

import (
	"callback_service/src/broker"
	"callback_service/src/database"
	"encoding/json"
	"errors"
	"time"
)

func validateCallbackFields(callback *database.Callback) error {
	if callback.Name == "" || callback.Phone == "" || callback.Type == "" || callback.Idea == "" {
		return errors.New("one or more required fields are empty")
	}
	return nil
}

func ConvertToRepositoryAndValidate(callbackSrc []byte) (*database.Callback, error) {
	callback := new(database.Callback)
	err := json.Unmarshal(callbackSrc, callback)
	if err != nil {
		return nil, err
	}

	callback.CreatedAt = time.Now()
	err = validateCallbackFields(callback)
	if err != nil {
		return nil, err
	}
	return callback, nil
}

func (c *CMS) CreateResponse(callback *database.Callback) error {
	callbackJSON, err := json.Marshal(callback)
	if err != nil {
		c.Logger.Error("error marshalling callback: %s", err.Error())
		return err
	}

	msg := broker.NewMessage(string(callbackJSON), "callback")
	if err := c.Broker.Publish(msg); err != nil {
		return err
	}
	return nil
}
