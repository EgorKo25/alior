package service

import (
	"callback_service/src/broker"
	"callback_service/src/database"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *CMS) handleError(errMsg string) error {
	c.Logger.Error(errMsg)
	response := broker.NewMessage(errMsg, "error")
	return c.Broker.Publish(response)
}

// HandleMessage is a CMS method to handle all messages,
// with delivery.Headers contains "action_type" and delivery.Type in (create, initial, next, previous, delete)
func (c *CMS) HandleMessage(ctx context.Context, delivery amqp.Delivery) error {
	actionType, ok := delivery.Headers["action_type"].(string)
	if !ok {
		return c.handleError("invalid or missing action_type header")
	}

	if delivery.Type == "callback" {
		switch actionType {
		case "create":
			if err := c.CreateCallbackHandler(ctx, delivery); err != nil {
				return c.handleError("error creating callback: " + err.Error())
			}
			return c.Broker.Publish(broker.NewMessage("created callback", "success"))

		case "initial":
			if err := c.InitialCallbackHandler(ctx); err != nil {
				return c.handleError("error getting initial callback: " + err.Error())
			}
			return nil

		case "next":
			if err := c.NextCallbackHandler(ctx); err != nil {
				return c.handleError("error getting next callback: " + err.Error())
			}
			return nil

		case "previous":
			if err := c.PreviousCallbackHandler(ctx); err != nil {
				return c.handleError("error getting previous callback: " + err.Error())
			}
			return nil

		case "delete":
			if err := c.DeleteCallbackHandler(ctx, delivery); err != nil {
				return c.handleError("error deleting callback: " + err.Error())
			}
			return nil

		default:
			return c.handleError("unknown action type: " + actionType)
		}
	}
	return nil
}

// CreateCallbackHandler is a CMS method to handle message with delivery.Headers "action_type" = "create"
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

// InitialCallbackHandler is a CMS method to handle message with delivery.Headers "action_type" = "initial"
func (c *CMS) InitialCallbackHandler(ctx context.Context) error {
	database.Offset = 0
	callback, err := c.Storage.GetCallback(ctx, database.Limit, 0)
	if err != nil {
		c.Logger.Error("error getting initial callback: %s", err.Error())
		response := broker.NewMessage("error getting initial callback", "error")
		return c.Broker.Publish(response)
	}

	callbackJSON, err := json.Marshal(callback)
	if err != nil {
		c.Logger.Error("error marshalling callback: %s", err.Error())
		response := broker.NewMessage("error marshalling callback", "error")
		return c.Broker.Publish(response)
	}

	c.Logger.Info("got initial callback: %s", callback)
	response := broker.NewMessage(string(callbackJSON), "success")
	return c.Broker.Publish(response)
}

// NextCallbackHandler is a CMS method to handle message with delivery.Headers "action_type" = "next"
func (c *CMS) NextCallbackHandler(ctx context.Context) error {
	total, err := c.Storage.GetTotalCallbacks(ctx)
	if err != nil {
		c.Logger.Error("error getting total callbacks: %s", err.Error())
		response := broker.NewMessage("error getting total callbacks", "error")
		return c.Broker.Publish(response)
	}

	if database.Offset+1 < total {
		database.Offset++
	}

	callback, err := c.Storage.GetCallback(ctx, database.Limit, database.Offset)
	if err != nil {
		c.Logger.Error("error fetching next callback: %s", err.Error())
		response := broker.NewMessage("error fetching next callback", "error")
		return c.Broker.Publish(response)
	}

	callbackJSON, err := json.Marshal(callback)
	if err != nil {
		c.Logger.Error("error marshalling callback: %s", err.Error())
		response := broker.NewMessage("error marshalling callback", "error")
		return c.Broker.Publish(response)
	}

	c.Logger.Info("got next callback: %s", callback)
	response := broker.NewMessage(string(callbackJSON), "success")
	return c.Broker.Publish(response)
}

// PreviousCallbackHandler is a CMS method to handle message with delivery.Headers "action_type" = "previous"
func (c *CMS) PreviousCallbackHandler(ctx context.Context) error {
	if database.Offset > 0 {
		database.Offset--
	}

	callback, err := c.Storage.GetCallback(ctx, database.Limit, database.Offset)
	if err != nil {
		c.Logger.Error("error fetching previous callback: %s", err.Error())
		response := broker.NewMessage("error fetching previous callback", "error")
		return c.Broker.Publish(response)
	}

	callbackJSON, err := json.Marshal(callback)
	if err != nil {
		c.Logger.Error("error marshalling callback: %s", err.Error())
		response := broker.NewMessage("error marshalling callback", "error")
		return c.Broker.Publish(response)
	}

	c.Logger.Info("got previous callback: %s", callback)
	response := broker.NewMessage(string(callbackJSON), "success")
	return c.Broker.Publish(response)
}

// DeleteCallbackHandler is a CMS method to handle message with delivery.Headers "action_type" = "delete"
func (c *CMS) DeleteCallbackHandler(ctx context.Context, delivery amqp.Delivery) error {
	var body database.Callback
	err := json.Unmarshal(delivery.Body, &body)
	if err != nil {
		c.Logger.Error("error unmarshalling message body: %s", err.Error())
		response := broker.NewMessage("error unmarshalling delete message", "error")
		return c.Broker.Publish(response)
	}

	err = c.Storage.DeleteCallbackByID(ctx, body.ID)
	if err != nil {
		c.Logger.Error("error deleting callback: %s", err.Error())
		response := broker.NewMessage("error deleting callback", "error")
		return c.Broker.Publish(response)
	}

	if database.Offset > 0 {
		database.Offset--
	}

	c.Logger.Info("deleted callback: %s", body.ID)
	response := broker.NewMessage("deleted callback", "success")
	return c.Broker.Publish(response)
}
