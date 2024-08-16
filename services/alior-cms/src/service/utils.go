package service

import (
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

func convertToRepositoryAndValidate(callbackSrc []byte) (*database.Callback, error) {
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
