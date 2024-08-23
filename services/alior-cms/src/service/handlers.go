package service

import (
	"callback_service/src/broker"
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
				c.logger.Error("error creating callback: %s", err)
				response := broker.NewMessage("error creating callback", "error")
				return c.broker.Publish(response)
			}
			c.logger.Info("callback created")

			response := broker.NewMessage("created successfully", "success")
			return c.broker.Publish(response)

		case "initial":
			err := c.InitialCallbackHandler(ctx)
			if err != nil {
				c.logger.Error("error getting init callback: %s", err)
				response := broker.NewMessage("error getting init callback", "error")
				return c.broker.Publish(response)
			}
			c.logger.Info("got initial callback")

			response := broker.NewMessage("got initial callback", "success")
			return c.broker.Publish(response)

		case "next":
			err := c.NextCallbackHandler(ctx)
			if err != nil {
				c.logger.Error("error getting next callback: %s", err)
				response := broker.NewMessage("error getting next callback", "error")
				return c.broker.Publish(response)
			}
			c.logger.Info("got next callback")

			response := broker.NewMessage("got next callback", "success")
			return c.broker.Publish(response)

		case "previous":
			err := c.PreviousCallbackHandler(ctx)
			if err != nil {
				c.logger.Error("error getting previous callback: %s", err)
				response := broker.NewMessage("error getting previous callback", "error")
				return c.broker.Publish(response)
			}
			c.logger.Info("got previous callback")

			response := broker.NewMessage("got previous callback", "success")
			return c.broker.Publish(response)

		case "delete":
			err := c.DeleteCallbackHandler(ctx, delivery)
			if err != nil {
				c.logger.Error("error deleting callback: %s", err)
				response := broker.NewMessage("error deleting callback", "error")
				return c.broker.Publish(response)
			}
			c.logger.Info("got delete callback")

			response := broker.NewMessage("deleted successfully", "success")
			return c.broker.Publish(response)

		default:
			c.logger.Error("unknown message type: %s", delivery.Type)

			response := broker.NewMessage("unknown message type", "error")
			return c.broker.Publish(response)
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

func (c *CMS) InitialCallbackHandler(ctx context.Context) error {
	database.Offset = 0
	callback, err := c.storage.GetCallback(ctx, database.Limit, 0)
	if err != nil {
		c.logger.Error("error fetching callback: %s", err.Error())
		return err
	}

	err = c.createResponse(callback)
	if err != nil {
		c.logger.Error("error processing rabbit response: %s", err.Error())
		return err
	}

	c.logger.Info("got initial callback: %s", callback)
	return nil
}

func (c *CMS) NextCallbackHandler(ctx context.Context) error {
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

	err = c.createResponse(callback)
	if err != nil {
		c.logger.Error("error processing rabbit response: %s", err.Error())
		return err
	}

	c.logger.Info("got next callback: %s", callback)
	return nil
}

func (c *CMS) PreviousCallbackHandler(ctx context.Context) error {
	if database.Offset > 0 {
		database.Offset -= 1
	}
	callback, err := c.storage.GetCallback(ctx, database.Limit, database.Offset)
	if err != nil {
		c.logger.Error("error fetching callback: %s", err.Error())
		return err
	}
	err = c.createResponse(callback)
	if err != nil {
		c.logger.Error("error processing rabbit response: %s", err.Error())
		return err
	}

	c.logger.Info("got previous callback: %s", callback)
	return nil
}

func (c *CMS) DeleteCallbackHandler(ctx context.Context, delivery amqp.Delivery) error {
	var body database.Callback
	err := json.Unmarshal(delivery.Body, &body)
	c.logger.Info("got delete request: %s", delivery.Body)
	if err != nil {
		c.logger.Error("error unmarshalling message body: %s", err.Error())
		return err
	}

	err = c.storage.DeleteCallbackByID(ctx, body.ID)
	if err != nil {
		c.logger.Error("error deleting callback: %s", err.Error())
		return err
	}

	if database.Offset-1 >= 0 {
		database.Offset -= 1
	}

	c.logger.Info("deleted callback: %s", body.ID)
	return nil
}
