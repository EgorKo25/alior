package service

import (
	"callback_service/src/database"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *CMS) HandleMessage(ctx context.Context, delivery amqp.Delivery) error {
	if delivery.ContentType == "callback" {
		switch delivery.Type {
		case "create":
			err := c.CreateCallbackHandler(ctx, delivery)
			if err != nil {
				c.logger.Error("failed to create callback: %s", err)
				return c.broker.Produce(ctx, delivery, err.Error(), "error")
			} else {
				c.logger.Info("callback created successfully")
				return c.broker.Produce(ctx, delivery, "callback created successfully", "success")
			}
		case "initial", "next", "previous":
			err := c.HandleAskMessage(ctx, delivery)
			delivery.Body = []byte("")
			if err != nil {
				c.logger.Error("error: %s", err.Error())
				return c.broker.Produce(ctx, delivery, err.Error(), "error")
			} else {
				c.logger.Info("successful query build")
				return c.broker.Produce(ctx, delivery, "successful query build", "success")
			}
		case "delete":
			err := c.HandleDeleteMessage(ctx, delivery)
			if err != nil {
				c.logger.Error("error: %s", err.Error())
				return c.broker.Produce(ctx, delivery, err.Error(), "error")
			} else {
				c.logger.Info("successful delete")
				return c.broker.Produce(ctx, delivery, "successful delete", "success")
			}
		default:
			c.logger.Error("unknown action type: %s", delivery.Type)
			return c.broker.Produce(ctx, delivery, "unknown action type", "error")
		}
	}
	return nil
}

func (c *CMS) CreateCallbackHandler(ctx context.Context, delivery amqp.Delivery) error {
	callback, err := convertToRepositoryAndValidate(delivery.Body)
	if err != nil {
		c.logger.Error("error during validation or conversion: %s", err.Error())
		return err
	}

	err = c.storage.CreateCallback(ctx, callback)
	if err != nil {
		c.logger.Error("error inserting callback: %s", err.Error())
		return err
	}
	return nil
}

func (c *CMS) HandleAskMessage(ctx context.Context, delivery amqp.Delivery) error {
	c.logger.Info("Offset: %d", database.Offset)
	switch delivery.Type {
	case "initial":
		database.Offset = 0
		callback, err := c.storage.GetCallback(ctx, database.Limit, 0)
		if err != nil {
			c.logger.Error("error fetching callback: %s", err.Error())
			return err
		}
		err = c.processRabbitResponse(ctx, callback, delivery)
		if err != nil {
			c.logger.Error("error processing rabbit response: %s", err.Error())
			return err
		}
	case "next":
		total, err := c.storage.GetTotalCallbacks(ctx)
		if err != nil {
			c.logger.Error("error getting total callbacks: %s", err.Error())
		}
		if database.Offset+1 < total {
			database.Offset += 1
		}
		callback, err := c.storage.GetCallback(ctx, database.Limit, database.Offset)
		if err != nil {
			c.logger.Error("error fetching callback: %s", err.Error())
			return err
		}
		err = c.processRabbitResponse(ctx, callback, delivery)
		if err != nil {
			c.logger.Error("error processing rabbit response: %s", err.Error())
			return err
		}
	case "previous":
		if database.Offset > 0 {
			database.Offset -= 1
		}
		callback, err := c.storage.GetCallback(ctx, database.Limit, database.Offset)
		if err != nil {
			c.logger.Error("error fetching callback: %s", err.Error())
			return err
		}
		err = c.processRabbitResponse(ctx, callback, delivery)
		if err != nil {
			c.logger.Error("error processing rabbit response: %s", err.Error())
			return err
		}
	}
	c.logger.Info("Offset new: %d", database.Offset)
	return nil
}

func (c *CMS) HandleDeleteMessage(ctx context.Context, delivery amqp.Delivery) error {
	var body database.Callback
	err := json.Unmarshal(delivery.Body, &body)
	if err != nil {
		c.logger.Error("error unmarshaling message body: %s", err.Error())
		return err
	}
	err = c.storage.DeleteCallbackByID(ctx, body.ID)
	if err != nil {
		c.logger.Error("error deleting callback: %s", err.Error())
		return err
	}
	return nil
}

func (c *CMS) processRabbitResponse(ctx context.Context, callback *database.Callback, delivery amqp.Delivery) error {
	callbackJSON, err := json.Marshal(callback)
	if err != nil {
		c.logger.Error("error marshalling callback: %s", err.Error())
		return err
	}
	if err := c.broker.Produce(ctx, delivery, string(callbackJSON), "success"); err != nil {
		return err
	}
	return nil
}
