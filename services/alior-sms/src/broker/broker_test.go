package broker_test

import (
	"alior-sms/src/broker"
	"context"
	"testing"
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
				if err == nil {
					t.Fatalf("expected error: %v, but got: %v", tt.expectError, err)
				}

				return
			}

			// Проверка, что канал сессий был создан
			if sessionChan == nil {
				t.Fatal("expected session channel to be created, but got nil")
			}

			// Ожидание первой сессии из канала, если это возможно
			select {
			case session := <-sessionChan:
				if session.Connection == nil || session.Channel == nil {
					t.Fatal("expected valid connection and channel, but got nil")
				}
			case <-ctx.Done():
				t.Fatal("context cancelled before session could be received")
			}
		})
	}
}
