package broker_test

import (
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"testing"

	"callback_service/src/broker"
	mocks "callback_service/src/broker/mocks"
	"github.com/golang/mock/gomock"
)

func TestBroker_Publish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		message       *broker.Message
		setupMocks    func(broker *mocks.MockIBroker)
		expectedError error
	}{
		{
			name: "successful publish",
			message: &broker.Message{
				Headers: broker.MessageHeaders{
					Exchange:   "test-exchange",
					RoutingKey: "test-routing-key",
				},
				Body: "test-message",
			},
			setupMocks: func(broker *mocks.MockIBroker) {
				broker.EXPECT().Publish(
					gomock.Any(),
				).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "publish error",
			message: &broker.Message{
				Headers: broker.MessageHeaders{
					Exchange:   "test-exchange",
					RoutingKey: "test-routing-key",
				},
				Body: "test-message",
			},
			setupMocks: func(broker *mocks.MockIBroker) {
				broker.EXPECT().Publish(
					gomock.Any(),
				).Return(errors.New("publish error"))
			},
			expectedError: errors.New("publish error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := mocks.NewMockIBroker(ctrl)

			tt.setupMocks(b)

			err := b.Publish(tt.message)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestBroker_Subscribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		queue         string
		setupMocks    func(broker *mocks.MockIBroker)
		handlerFunc   func(ctx context.Context, delivery amqp.Delivery) error
		expectedError error
	}{
		{
			name: "successful subscribe",
			setupMocks: func(broker *mocks.MockIBroker) {
				broker.EXPECT().Subscribe(
					gomock.Any(),
					gomock.Any(),
				).Return(nil)
			},
			handlerFunc: func(ctx context.Context, delivery amqp.Delivery) error {
				return nil
			},
			expectedError: nil,
		},
		{
			name: "subscribe error",
			setupMocks: func(broker *mocks.MockIBroker) {
				broker.EXPECT().Subscribe(
					gomock.Any(),
					gomock.Any(),
				).Return(errors.New("subscribe error"))
			},
			handlerFunc: func(ctx context.Context, delivery amqp.Delivery) error {
				return nil
			},
			expectedError: errors.New("subscribe error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockBroker := mocks.NewMockIBroker(ctrl)

			tt.setupMocks(mockBroker)

			err := mockBroker.Subscribe(context.Background(), tt.handlerFunc)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
