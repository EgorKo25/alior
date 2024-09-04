package service

import (
	"callback_service/src/broker"
	"callback_service/src/database"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

// HandleMessage is a CMS method to handle all messages,
// which delivery.ContentType == "callback" and delivery.Type in (create, initial, next, previous, delete)
func (c *CMS) HandleMessage(ctx context.Context, delivery amqp.Delivery) error {
	if delivery.ContentType == "callback" {
		switch delivery.Type {
		case "create":
			err := c.CreateCallbackHandler(ctx, delivery)
			if err != nil {
				c.Logger.Error("error creating callback: %s", err)
				response := broker.NewMessage("error creating callback", "error")
				return c.Broker.Publish(response)
			}
			c.Logger.Info("callback created")

			response := broker.NewMessage("created successfully", "success")
			return c.Broker.Publish(response)

		case "initial":
			err := c.InitialCallbackHandler(ctx)
			if err != nil {
				c.Logger.Error("error getting init callback: %s", err)
				response := broker.NewMessage("error getting init callback", "error")
				return c.Broker.Publish(response)
			}
			c.Logger.Info("got initial callback")

			response := broker.NewMessage("got initial callback", "success")
			return c.Broker.Publish(response)

		case "next":
			err := c.NextCallbackHandler(ctx)
			if err != nil {
				c.Logger.Error("error getting next callback: %s", err)
				response := broker.NewMessage("error getting next callback", "error")
				return c.Broker.Publish(response)
			}
			c.Logger.Info("got next callback")

			response := broker.NewMessage("got next callback", "success")
			return c.Broker.Publish(response)

		case "previous":
			err := c.PreviousCallbackHandler(ctx)
			if err != nil {
				c.Logger.Error("error getting previous callback: %s", err)
				response := broker.NewMessage("error getting previous callback", "error")
				return c.Broker.Publish(response)
			}
			c.Logger.Info("got previous callback")

			response := broker.NewMessage("got previous callback", "success")
			return c.Broker.Publish(response)

		case "delete":
			err := c.DeleteCallbackHandler(ctx, delivery)
			if err != nil {
				c.Logger.Error("error deleting callback: %s", err)
				response := broker.NewMessage("error deleting callback", "error")
				return c.Broker.Publish(response)
			}
			c.Logger.Info("got delete callback")

			response := broker.NewMessage("deleted successfully", "success")
			return c.Broker.Publish(response)

		default:
			c.Logger.Error("unknown message type: %s", delivery.Type)

			response := broker.NewMessage("unknown message type", "error")
			return c.Broker.Publish(response)
		}
	}
	return nil
}

// CreateCallbackHandler is a CMS method to handle message with delivery.type = create
func (c *CMS) CreateCallbackHandler(ctx context.Context, delivery amqp.Delivery) error {
	callback, err := ConvertToRepositoryAndValidate(delivery.Body)
	if err != nil {
		c.Logger.Error("error during validation or conversion: %s", err.Error())
		return err
	}

	err = c.Storage.CreateCallback(ctx, callback)
	if err != nil {
		c.Logger.Error("error creating callback: %s", err.Error())
		return err
	}

	return nil
}

// InitialCallbackHandler is a CMS method to handle message with delivery.type = initial
func (c *CMS) InitialCallbackHandler(ctx context.Context) error {
	database.Offset = 0
	callback, err := c.Storage.GetCallback(ctx, database.Limit, 0)
	if err != nil {
		c.Logger.Error("error getting init callback: %s", err.Error())
		return err
	}

	err = c.CreateResponse(callback)
	if err != nil {
		c.Logger.Error("error processing rabbit response: %s", err.Error())
		return err
	}

	c.Logger.Info("got initial callback: %s", callback)
	return nil
}

// NextCallbackHandler is a CMS method to handle message with delivery.type = next
func (c *CMS) NextCallbackHandler(ctx context.Context) error {
	total, err := c.Storage.GetTotalCallbacks(ctx)
	if err != nil {
		c.Logger.Error("error getting total callbacks: %s", err.Error())
		return err
	}

	if database.Offset+1 < total {
		database.Offset++
	}

	callback, err := c.Storage.GetCallback(ctx, database.Limit, database.Offset)
	if err != nil {
		c.Logger.Error("error fetching callback: %s", err.Error())
		return err
	}

	err = c.CreateResponse(callback)
	if err != nil {
		c.Logger.Error("error processing rabbit response: %s", err.Error())
		return err
	}

	c.Logger.Info("got next callback: %s", callback)
	return nil
}

// PreviousCallbackHandler is a CMS method to handle message with delivery.type = previous
func (c *CMS) PreviousCallbackHandler(ctx context.Context) error {
	if database.Offset > 0 {
		database.Offset--
	}
	callback, err := c.Storage.GetCallback(ctx, database.Limit, database.Offset)
	if err != nil {
		c.Logger.Error("error fetching callback: %s", err.Error())
		return err
	}
	err = c.CreateResponse(callback)
	if err != nil {
		c.Logger.Error("error processing rabbit response: %s", err.Error())
		return err
	}

	c.Logger.Info("got previous callback: %s", callback)
	return nil
}

// DeleteCallbackHandler is a CMS method to handle message with delivery.type = delete
func (c *CMS) DeleteCallbackHandler(ctx context.Context, delivery amqp.Delivery) error {
	var body database.Callback
	err := json.Unmarshal(delivery.Body, &body)
	if err != nil {
		c.Logger.Error("error unmarshalling message body: %s", err.Error())
		return err
	}

	err = c.Storage.DeleteCallbackByID(ctx, body.ID)
	if err != nil {
		c.Logger.Error("error deleting callback: %s", err.Error())
		return err
	}

	if database.Offset-1 >= 0 {
		database.Offset--
	}

	c.Logger.Info("deleted callback: %s", body.ID)
	return nil
}
