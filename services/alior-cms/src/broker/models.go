package broker

import (
	"time"
)

// Message structure of broker delivery
type Message struct {
	Properties MessageProperties `json:"properties"`
	Body       string            `json:"body"`
	Headers    MessageHeaders    `json:"headers"`
}

// MessageProperties structure of delivery properties
type MessageProperties struct {
	ContentType   string `json:"content-type"`
	DeliveryMode  int    `json:"delivery-mode"`
	CorrelationID string `json:"correlation-id"`
	ReplyTo       string `json:"reply-to"`
	Timestamp     int64  `json:"timestamp"`
	Type          string `json:"type"`
	AppID         string `json:"app-id"`
}

// MessageHeaders structure of delivery headers
type MessageHeaders struct {
	Exchange    string `json:"exchange"`
	RoutingKey  string `json:"routing-key"`
	Mandatory   bool   `json:"mandatory"`
	Immediate   bool   `json:"immediate"`
	DeliveryTag uint64 `json:"delivery-tag"`
}

// NewMessage is a broker message constructor
func NewMessage(body string, msgType string) *Message {
	return &Message{
		Properties: MessageProperties{
			ContentType:   "text/plain",
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
