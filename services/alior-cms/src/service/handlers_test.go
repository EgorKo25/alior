package service_test

import (
	"callback_service/src/broker"
	brokerMocks "callback_service/src/broker/mocks"
	"callback_service/src/database"
	dbMocks "callback_service/src/database/mocks"
	"callback_service/src/service"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateCallbackHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := dbMocks.NewMockICallback(ctrl)
	mockLogger := brokerMocks.NewMockILogger(ctrl)

	tests := []struct {
		name        string
		inputBody   []byte
		setupMocks  func()
		expectError bool
	}{
		{
			name:      "Success",
			inputBody: []byte(`{"id":1,"name":"test","phone":"test","idea":"test","type":"test"}`),
			setupMocks: func() {
				mockStorage.EXPECT().
					CreateCallback(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expectError: false,
		},
		{
			name:      "Validation Error",
			inputBody: []byte(`{"id":1,"name":"test"}`),
			setupMocks: func() {
				mockLogger.EXPECT().
					Error(gomock.Any(), gomock.Any())
			},
			expectError: true,
		},
		{
			name:      "Conversion Error",
			inputBody: []byte(`invalid-json`),
			setupMocks: func() {
				mockLogger.EXPECT().
					Error(gomock.Any(), gomock.Any())
			},
			expectError: true,
		},
		{
			name:      "Storage Error",
			inputBody: []byte(`{"id":1,"name":"test","phone":"test","idea":"test","type":"test"}`),
			setupMocks: func() {
				mockStorage.EXPECT().
					CreateCallback(gomock.Any(), gomock.Any()).
					Return(errors.New("storage error"))

				mockLogger.EXPECT().
					Error(gomock.Any(), gomock.Any()).
					Return()
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &service.CMS{
				Storage: mockStorage,
				Logger:  mockLogger,
			}

			tt.setupMocks()

			delivery := amqp.Delivery{Body: tt.inputBody}
			err := c.CreateCallbackHandler(context.Background(), delivery)

			require.Equal(t, tt.expectError, err != nil)
		})
	}
}

func TestInitialCallbackHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := dbMocks.NewMockICallback(ctrl)
	mockLogger := brokerMocks.NewMockILogger(ctrl)
	mockBroker := brokerMocks.NewMockIBroker(ctrl)

	ctx := context.Background()

	cms := &service.CMS{
		Storage: mockStorage,
		Logger:  mockLogger,
		Broker:  mockBroker,
	}

	tests := []struct {
		name          string
		setupMocks    func()
		expectedError error
	}{
		{
			name: "success",
			setupMocks: func() {
				callback := &database.Callback{ID: 1}
				callbackJSON, _ := json.Marshal(callback)

				mockStorage.EXPECT().
					GetCallback(ctx, gomock.Any(), 0).
					Return(callback, nil)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage(string(callbackJSON), "success")
						assert.Equal(t, expectedMsg, msg)
						return nil
					})

				mockLogger.EXPECT().
					Info(gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedError: nil,
		},
		{
			name: "storage get error",
			setupMocks: func() {
				mockStorage.EXPECT().
					GetCallback(ctx, gomock.Any(), 0).
					Return(nil, errors.New("get callback error"))

				mockLogger.EXPECT().
					Error(gomock.Any(), gomock.Any()).
					Times(1)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage("error getting initial callback", "error")
						assert.Equal(t, expectedMsg, msg)
						return nil
					})
			},
			expectedError: nil,
		},
		{
			name: "broker publish error",
			setupMocks: func() {
				callback := &database.Callback{ID: 1}
				callbackJSON, _ := json.Marshal(callback)

				mockStorage.EXPECT().
					GetCallback(ctx, gomock.Any(), 0).
					Return(callback, nil)

				mockLogger.EXPECT().
					Info(gomock.Any(), gomock.Any()).
					Times(1)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage(string(callbackJSON), "success")
						assert.Equal(t, expectedMsg, msg)
						return errors.New("publish error")
					})
			},
			expectedError: errors.New("publish error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := cms.InitialCallbackHandler(ctx)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNextCallbackHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := dbMocks.NewMockICallback(ctrl)
	mockLogger := brokerMocks.NewMockILogger(ctrl)
	mockBroker := brokerMocks.NewMockIBroker(ctrl)

	ctx := context.Background()

	cms := &service.CMS{
		Storage: mockStorage,
		Logger:  mockLogger,
		Broker:  mockBroker,
	}

	tests := []struct {
		name          string
		setupMocks    func()
		expectedError error
	}{
		{
			name: "success",
			setupMocks: func() {
				callback := &database.Callback{ID: 1}
				callbackJSON, _ := json.Marshal(callback)

				mockStorage.EXPECT().
					GetTotalCallbacks(ctx).
					Return(5, nil)

				mockStorage.EXPECT().
					GetCallback(ctx, database.Limit, database.Offset+1).
					Return(callback, nil)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage(string(callbackJSON), "success")
						assert.Equal(t, expectedMsg, msg)
						return nil
					})

				mockLogger.EXPECT().
					Info(gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedError: nil,
		},
		{
			name: "storage get callback error",
			setupMocks: func() {
				mockStorage.EXPECT().
					GetTotalCallbacks(ctx).
					Return(5, nil)

				mockStorage.EXPECT().
					GetCallback(ctx, database.Limit, database.Offset+1).
					Return(nil, errors.New("get callback error"))

				mockLogger.EXPECT().
					Error(gomock.Any(), gomock.Any()).
					Times(1)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage("error fetching next callback", "error")
						assert.Equal(t, expectedMsg, msg)
						return nil
					})
			},
			expectedError: nil,
		},
		{
			name: "broker publish error",
			setupMocks: func() {
				callback := &database.Callback{ID: 1}
				callbackJSON, _ := json.Marshal(callback)

				mockStorage.EXPECT().
					GetTotalCallbacks(ctx).
					Return(5, nil)

				mockStorage.EXPECT().
					GetCallback(ctx, database.Limit, database.Offset+1).
					Return(callback, nil)

				mockLogger.EXPECT().
					Info(gomock.Any(), gomock.Any()).
					Times(1)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage(string(callbackJSON), "success")
						assert.Equal(t, expectedMsg, msg)
						return errors.New("publish error")
					})
			},
			expectedError: errors.New("publish error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := cms.NextCallbackHandler(ctx)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPreviousCallbackHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := dbMocks.NewMockICallback(ctrl)
	mockLogger := brokerMocks.NewMockILogger(ctrl)
	mockBroker := brokerMocks.NewMockIBroker(ctrl)

	ctx := context.Background()

	cms := &service.CMS{
		Storage: mockStorage,
		Logger:  mockLogger,
		Broker:  mockBroker,
	}

	tests := []struct {
		name          string
		setupMocks    func()
		expectedError error
	}{
		{
			name: "success",
			setupMocks: func() {
				callback := &database.Callback{ID: 1}
				callbackJSON, _ := json.Marshal(callback)

				database.Offset++

				mockStorage.EXPECT().
					GetCallback(ctx, database.Limit, database.Offset-1).
					Return(callback, nil)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage(string(callbackJSON), "success")
						assert.Equal(t, expectedMsg, msg)
						return nil
					})

				mockLogger.EXPECT().
					Info(gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedError: nil,
		},
		{
			name: "storage get callback error",
			setupMocks: func() {
				database.Offset++

				mockStorage.EXPECT().
					GetCallback(ctx, database.Limit, database.Offset-1).
					Return(nil, errors.New("get callback error"))

				mockLogger.EXPECT().
					Error("error fetching previous callback: %s", "get callback error").
					Times(1)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage("error fetching previous callback", "error")
						assert.Equal(t, expectedMsg, msg)
						return nil
					})
			},
			expectedError: nil,
		},
		{
			name: "broker publish error",
			setupMocks: func() {
				callback := &database.Callback{ID: 1}
				callbackJSON, _ := json.Marshal(callback)

				database.Offset++

				mockStorage.EXPECT().
					GetCallback(ctx, database.Limit, database.Offset-1).
					Return(callback, nil)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage(string(callbackJSON), "success")
						assert.Equal(t, expectedMsg, msg)
						return errors.New("publish error")
					})

				mockLogger.EXPECT().
					Info(gomock.Any(), gomock.Any()).
					Times(1)
			},
			expectedError: errors.New("publish error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := cms.PreviousCallbackHandler(ctx)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteCallbackHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := dbMocks.NewMockICallback(ctrl)
	mockLogger := brokerMocks.NewMockILogger(ctrl)
	mockBroker := brokerMocks.NewMockIBroker(ctrl)

	ctx := context.Background()

	cms := &service.CMS{
		Storage: mockStorage,
		Logger:  mockLogger,
		Broker:  mockBroker,
	}

	tests := []struct {
		name          string
		inputBody     []byte
		setupMocks    func()
		expectedError error
	}{
		{
			name:      "success",
			inputBody: []byte(`{"id":1}`),
			setupMocks: func() {
				mockStorage.EXPECT().
					DeleteCallbackByID(ctx, int32(1)).
					Return(nil)

				mockLogger.EXPECT().
					Info(gomock.Any(), gomock.Any()).
					Times(1)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage("deleted callback", "success")
						assert.Equal(t, expectedMsg, msg)
						return nil
					})
			},
			expectedError: nil,
		},
		{
			name:      "unmarshal error",
			inputBody: []byte(`invalid-json`),
			setupMocks: func() {
				mockLogger.EXPECT().
					Error(gomock.Any(), gomock.Any()).
					Times(1)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage("error unmarshalling delete message", "error")
						assert.Equal(t, expectedMsg, msg)
						return nil
					})
			},
			expectedError: nil,
		},
		{
			name:      "storage delete error",
			inputBody: []byte(`{"id":1}`),
			setupMocks: func() {
				mockStorage.EXPECT().
					DeleteCallbackByID(ctx, int32(1)).
					Return(errors.New("delete error"))

				mockLogger.EXPECT().
					Error(gomock.Any(), gomock.Any()).
					Times(1)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					DoAndReturn(func(msg *broker.Message) error {
						expectedMsg := broker.NewMessage("error deleting callback", "error")
						assert.Equal(t, expectedMsg, msg)
						return nil
					})
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			delivery := amqp.Delivery{Body: tt.inputBody}
			err := cms.DeleteCallbackHandler(ctx, delivery)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
