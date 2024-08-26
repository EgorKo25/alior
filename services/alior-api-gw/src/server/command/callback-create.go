package command

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CallbackCreate struct {
	Idea                string `json:"idea"  validate:"required,min=3"`
	Username            string `json:"name" validate:"required,min=2,max=100"`
	PhoneNumber         string `json:"phone_number" validate:"required,e164"`
	CommunicationMethod int32  `json:"communication_method" validate:"required,oneof=1 2 3"`
}

func (c *CallbackCreate) Name() string {
	return "callback/create"
}

func (c *CallbackCreate) Parse(ctx *gin.Context) error {
	var body []byte
	if _, err := ctx.Request.Body.Read(body); err != nil {
		return err
	}
	if err := json.Unmarshal(body, c); err != nil {
		return err
	}
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("Field: %s, Error: %s\n", err.Field(), err.Tag())
		}
		return err
	}
	return nil
}

func (c *CallbackCreate) Apply() (any, error) {
	/*
		body, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}

		id := sha256.Sum256(body)
		request := &amqp.Publishing{
			CorrelationId: fmt.Sprintf("%x", id),
			Timestamp:     time.Now(),
			Type:          "create",
			ReplyTo:       "ANS",
			DeliveryMode:  1,
			ContentType:   "callback",
			Body:          body,
		}
	*/
	// TODO:отправка в Rabbit
	//
	// TODO: и ожидание ответа
	//
	// responce :=
	return nil, nil
}
