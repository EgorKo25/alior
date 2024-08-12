package broker_test

import (
	"alior-sms/src/broker"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRedial(t *testing.T) {
	tests := []struct {
		name         string
		ctx          context.Context
		uri          string
		exchangeName string
		exchangeKind string
		routingKey   string
		queueName    string
		expectError  bool
	}{
		{
			name:         "successful connection",
			ctx:          context.Background(),
			url:          "amqp://guest:guest@localhost:5672/",
			exchangeName: "test_exchange",
			routingKey:   "test_key",
			queueName:    "test_queue",
			consumerTag:  "test_consumer",
			expectError:  false,
		},
		{
			name:         "invalid URL",
			ctx:          context.Background(),
			url:          "amqp://invalid:url",
			exchangeName: "test_exchange",
			routingKey:   "test_key",
			queueName:    "test_queue",
			consumerTag:  "test_consumer",
			expectError:  true,
		},
		{
			name: "context cancellation before connection",
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
			url:          "amqp://guest:guest@localhost:5672/",
			exchangeName: "test_exchange",
			routingKey:   "test_key",
			queueName:    "test_queue",
			consumerTag:  "test_consumer",
			expectError:  true,
		},
		{
			name:         "connection to unavailable server",
			ctx:          context.Background(),
			url:          "amqp://guest:guest@unavailable:5672/",
			exchangeName: "test_exchange",
			routingKey:   "test_key",
			queueName:    "test_queue",
			consumerTag:  "test_consumer",
			expectError:  true,
		},
		{
			name:         "error creating channel",
			ctx:          context.Background(),
			url:          "amqp://guest:guest@localhost:5672/",
			exchangeName: "test_exchange",
			routingKey:   "test_key",
			queueName:    "test_queue",
			consumerTag:  "test_consumer",
			expectError:  true,
		},
		{
			name:         "context with timeout",
			ctx:          context.Background(),
			url:          "amqp://guest:guest@localhost:5672/",
			exchangeName: "test_exchange",
			routingKey:   "test_key",
			queueName:    "test_queue",
			consumerTag:  "test_consumer",
			expectError:  true,
		},
		{
			name:         "empty URL",
			ctx:          context.Background(),
			url:          "",
			exchangeName: "test_exchange",
			routingKey:   "test_key",
			queueName:    "test_queue",
			consumerTag:  "test_consumer",
			expectError:  true,
		},
		{
			name:         "reconnection after disconnection",
			ctx:          context.Background(),
			url:          "amqp://guest:guest@localhost:5672/",
			exchangeName: "test_exchange",
			routingKey:   "test_key",
			queueName:    "test_queue",
			consumerTag:  "test_consumer",
			expectError:  false,
		},
		{
			name:         "empty exchange name",
			ctx:          context.Background(),
			url:          "amqp://guest:guest@localhost:5672/",
			exchangeName: "",
			routingKey:   "test_key",
			queueName:    "test_queue",
			consumerTag:  "test_consumer",
			expectError:  true,
		},
		{
			name:         "empty routing key",
			ctx:          context.Background(),
			url:          "amqp://guest:guest@localhost:5672/",
			exchangeName: "test_exchange",
			routingKey:   "",
			queueName:    "test_queue",
			consumerTag:  "test_consumer",
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sessionChan := broker.Redial(tt.ctx, tt.url, tt.exchangeName, tt.routingKey, tt.queueName, tt.consumerTag)

			select {
			case session, ok := <-sessionChan:
				if tt.expectError {
					assert.False(t, ok, "expected no session to be received due to an error")
				} else {
					assert.True(t, ok, "expected session to be received")
					assert.NotNil(t, session.Channel, "expected a valid session channel")
				}
			case <-time.After(time.Second):
				if !tt.expectError {
					t.Fatal("expected a session to be sent, but none was received")
				}
			}
		})
	}
}
