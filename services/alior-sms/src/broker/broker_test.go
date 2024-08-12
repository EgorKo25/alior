package broker_test

import (
	"alior-sms/src/broker"
	"context"
	"testing"
	"time"
)

func TestDialSessionChan(t *testing.T) {
	tests := []struct {
		name        string
		ctx         context.Context
		Dconfig     broker.DialConfig
		expectError bool
	}{
		{
			name: "successful connection with queue",
			ctx:  context.Background(),
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
			ctx:  context.Background(),
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
			ctx:  context.Background(),
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
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			}(),
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
			ctx:  context.Background(),
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
			ctx:  context.Background(),
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
			name: "empty routing key",
			ctx:  context.Background(),
			Dconfig: *broker.NewDialConfig(
				"amqp://guest:guest@localhost:5672/",
				"test_exchange",
				"direct",
				"",
				"test_queue",
				true,
			),
			expectError: true,
		},
		{
			name: "empty queue name",
			ctx:  context.Background(),
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
			sessionChan, err := broker.DialSessionChan(tt.ctx, tt.Dconfig)

			// Ждем, пока установится соединение
			time.Sleep(1000 * time.Millisecond) // Ожидание 100 миллисекунд

			// Проверка, возникла ли ошибка
			if (err != nil) != tt.expectError {
				t.Fatalf("expected error: %v, but got: %v", tt.expectError, err)
			}

			// Если ожидается ошибка, то дальнейшие проверки не нужны
			if tt.expectError {
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
			case <-tt.ctx.Done():
				t.Fatal("context cancelled before session could be received")
			}

			// Закрываем сессию, чтобы не оставлять открытые соединения
			// Закрываем канал после завершения теста
			/*close(sessionChan) // Закрытие канала для завершения цикла

			for session := range sessionChan {
				err := session.Close()
				if err != nil {
					t.Errorf("failed to close session: %v", err)
				}
			}*/
		})
	}
}
