package broker

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type ILogger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type Broker struct {
	Url    string
	logger ILogger
}

func NewBroker(Url string, logger ILogger) *Broker {
	return &Broker{
		Url:    Url,
		logger: logger,
	}
}

type RabbitMQMessage struct {
	Properties struct {
		ContentType   string    `json:"content-type"`
		DeliveryMode  int       `json:"delivery-mode"`
		CorrelationId string    `json:"correlation-id"`
		ReplyTo       string    `json:"reply-to"`
		Timestamp     time.Time `json:"timestamp"`
		Type          string    `json:"type"`
		AppId         string    `json:"app-id"`
	} `json:"properties"`
	Body    string `json:"body"`
	Headers struct {
		Exchange    string `json:"exchange"`
		RoutingKey  string `json:"routing-key"`
		Mandatory   bool   `json:"mandatory"`
		Immediate   bool   `json:"immediate"`
		DeliveryTag int    `json:"delivery-tag"`
	} `json:"headers"`
}

func NewMessage(body string, delivery amqp.Delivery) *RabbitMQMessage {
	return &RabbitMQMessage{
		Properties: struct {
			ContentType   string    `json:"content-type"`
			DeliveryMode  int       `json:"delivery-mode"`
			CorrelationId string    `json:"correlation-id"`
			ReplyTo       string    `json:"reply-to"`
			Timestamp     time.Time `json:"timestamp"`
			Type          string    `json:"type"`
			AppId         string    `json:"app-id"`
		}{
			ContentType:   "callback",
			DeliveryMode:  int(delivery.DeliveryMode),
			CorrelationId: delivery.CorrelationId,
			ReplyTo:       "ans",
			Timestamp:     time.Now(),
			AppId:         "callback-service",
		},
		Body: body,
		Headers: struct {
			Exchange    string `json:"exchange"`
			RoutingKey  string `json:"routing-key"`
			Mandatory   bool   `json:"mandatory"`
			Immediate   bool   `json:"immediate"`
			DeliveryTag int    `json:"delivery-tag"`
		}{
			Exchange:    "ansask",
			RoutingKey:  "ans",
			Mandatory:   true,
			Immediate:   false,
			DeliveryTag: int(delivery.DeliveryTag),
		},
	}
}

func (b *Broker) Consume(ctx context.Context, queueName string, handler func(context.Context, amqp.Delivery) error) error {
	conn, err := amqp.Dial(b.Url)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			b.logger.Error("failed to close connection: %v", err)
		}
	}()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			b.logger.Error("failed to close channel: %v", err)
		}
	}()

	q, err := DeclareQueue(ch, queueName)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			b.logger.Info("received a message: %s", d.Body)

			if err := handler(ctx, d); err != nil {
				b.logger.Error("failed to handle message: %v", err)
			}
		}
	}()

	b.logger.Info("waiting for messages. To exit press CTRL+C")
	<-ctx.Done()

	return nil
}

func (b *Broker) Produce(ctx context.Context, delivery amqp.Delivery, responseBody string, responseType string) error {
	msg := CreateResponseMessage(delivery, responseBody, responseType)

	conn, err := amqp.Dial(b.Url)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			b.logger.Error("Failed to close connection: %v", err)
		}
	}()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer func() {
		if err := ch.Close(); err != nil {
			b.logger.Error("failed to close channel: %v", err)
		}
	}()

	_, err = DeclareQueue(ch, msg.Headers.RoutingKey)
	if err != nil {
		return err
	}

	ctxWTO, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	msgBody, err := json.Marshal(msg)
	if err != nil {
		b.logger.Error("failed to marshal message: %v", err)
		return err
	}

	err = ch.PublishWithContext(ctxWTO,
		msg.Headers.Exchange,
		msg.Headers.RoutingKey,
		msg.Headers.Mandatory,
		msg.Headers.Immediate,
		amqp.Publishing{
			ContentType:   msg.Properties.ContentType,
			DeliveryMode:  uint8(msg.Properties.DeliveryMode),
			CorrelationId: msg.Properties.CorrelationId,
			ReplyTo:       msg.Properties.ReplyTo,
			Timestamp:     msg.Properties.Timestamp,
			Type:          msg.Properties.Type,
			AppId:         msg.Properties.AppId,
			Body:          msgBody,
		})
	if err != nil {
		b.logger.Error("failed to produce message: %v", err)
		return err
	}

	b.logger.Info("sent message: %s", msg.Body)

	return nil
}
