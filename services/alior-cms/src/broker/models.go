package broker

import (
	"time"
)

type Message struct {
	Properties MessageProperties `json:"properties"`
	Body       string            `json:"body"`
	Headers    MessageHeaders    `json:"headers"`
}

type MessageProperties struct {
	ContentType   string `json:"content-type"`
	DeliveryMode  int    `json:"delivery-mode"`
	CorrelationID string `json:"correlation-id"`
	ReplyTo       string `json:"reply-to"`
	Timestamp     int64  `json:"timestamp"`
	Type          string `json:"type"`
	AppID         string `json:"app-id"`
}

type MessageHeaders struct {
	Exchange    string `json:"exchange"`
	RoutingKey  string `json:"routing-key"`
	Mandatory   bool   `json:"mandatory"`
	Immediate   bool   `json:"immediate"`
	DeliveryTag uint64 `json:"delivery-tag"`
}

func NewMessage(body string, msgType string) *Message {
	return &Message{
		Properties: MessageProperties{
			ContentType:   "callback",
			DeliveryMode:  1,
			CorrelationID: "123",
			ReplyTo:       "ans",
			Timestamp:     time.Now().Unix(),
			Type:          msgType,
			AppID:         "cms",
		},
		Body: body,
		Headers: MessageHeaders{
			Exchange:    "ansask",
			RoutingKey:  "ans",
			Mandatory:   true,
			Immediate:   false,
			DeliveryTag: 123,
		},
	}
}
