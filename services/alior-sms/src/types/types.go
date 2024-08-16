package types

type Service struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
}

type AMQPMessage struct {
	Properties Properties `json:"properties"`
	Body       string     `json:"body"`
	Headers    Headers    `json:"headers"`
}

type Properties struct {
	ContentType   string `json:"content-type"`
	DeliveryMode  int    `json:"delivery-mode"`
	CorrelationID string `json:"correlation-id"`
	ReplyTo       string `json:"reply-to"`
	Timestamp     int64  `json:"timestamp"`
	Type          string `json:"type"`
	AppID         string `json:"app-id"`
}

type Headers struct {
	Exchange    string `json:"exchange"`
	RoutingKey  string `json:"routing-key"`
	Mandatory   bool   `json:"mandatory"`
	Immediate   bool   `json:"immediate"`
	DeliveryTag uint64 `json:"delivery-tag"`
}
