package broker_test

import (
	"alior-sms/src/broker"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDialSessionChan(t *testing.T) {
	tests := []struct {
		name        string
		Dconfig     broker.DialConfig
		expectError bool
	}{
		{
			name: "successful connection with queue",
			Dconfig: *broker.NewDialConfig(
				"amqp://guest:guest@localhost:5672/",
				"test_exchange",
				"direct",
				"test_key",
				"test_queue",
				true,
			),
			expectError: false,
		},
		{
			name: "successful connection without queue",
			Dconfig: *broker.NewDialConfig(
				"amqp://guest:guest@localhost:5672/",
				"test_exchange",
				"direct",
				"test_key",
				"",
				false,
			),
			expectError: false,
		},
		{
			name: "invalid URL",
			Dconfig: *broker.NewDialConfig(
				"amqp://guest:guest@fkganoewn:asdasdsa/",
				"test_exchange",
				"direct",
				"test_key",
				"test_queue",
				true,
			),
			expectError: true,
		},
		{
			name: "context cancellation before connection",
			Dconfig: *broker.NewDialConfig(
				"amqp://guest:guest@localhost:5672/",
				"test_exchange",
				"direct",
				"test_key",
				"test_queue",
				true,
			),
			expectError: true,
		},
		{
			name: "empty URL",
			Dconfig: *broker.NewDialConfig(
				"",
				"test_exchange",
				"direct",
				"test_key",
				"test_queue",
				true,
			),
			expectError: true,
		},
		{
			name: "empty exchange name",
			Dconfig: *broker.NewDialConfig(
				"amqp://guest:guest@localhost:5672/",
				"",
				"direct",
				"test_key",
				"test_queue",
				true,
			),
			expectError: true,
		},
		{
			name: "empty queue name",
			Dconfig: *broker.NewDialConfig(
				"amqp://guest:guest@localhost:5672/",
				"test_exchange",
				"direct",
				"test_key",
				"",
				true,
			),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())

			if tt.name == "context cancellation before connection" {
				cancel()
			} else {
				defer cancel()
			}

			sessionChan, err := broker.DialSessionChan(ctx, tt.Dconfig)

			// Проверка, возникла ли ошибка
			if tt.expectError {
				require.Error(t, err, "expected an error but got none")
				return // Если ожидали ошибку, код ниже не нужен
			}

			// Проверка, что канал сессий был создан
			require.NotNil(t, sessionChan, "expected session channel to be created, but got nil")

			// Ожидание первой сессии из канала, если это возможно
			select {
			case session := <-sessionChan:
				require.NotNil(t, session.Connection, "expected valid connection, but got nil")
				require.NotNil(t, session.Channel, "expected valid channel, but got nil")
			case <-ctx.Done():
				require.FailNow(t, "Test failed", "context was cancelled before a session could be received in test: %s", tt.name)
			}
		})
	}
}
