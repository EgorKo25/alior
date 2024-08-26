package service_test

import (
	mocksbroker "callback_service/src/broker/mocks"
	"callback_service/src/database"
	mocksdb "callback_service/src/database/mocks"
	"callback_service/src/service"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	amqp "github.com/rabbitmq/amqp091-go"
	"testing"
)

func TestCreateCallbackHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocksdb.NewMockICallback(ctrl)
	mockLogger := mocksbroker.NewMockILogger(ctrl)

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

			if (err != nil) != tt.expectError {
				t.Errorf("CreateCallbackHandler() error = %v, wantErr %v", err, tt.expectError)
			}
		})
	}
}

func TestInitialCallbackHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocksdb.NewMockICallback(ctrl)
	mockLogger := mocksbroker.NewMockILogger(ctrl)
	mockBroker := mocksbroker.NewMockIBroker(ctrl)

	tests := []struct {
		name          string
		setupMocks    func(*service.CMS)
		expectedError bool
	}{
		{
			name: "Success",
			setupMocks: func(c *service.CMS) {
				callback := &database.Callback{ID: 1, Name: "test", Phone: "test", Idea: "test", Type: "test"}

				mockStorage.EXPECT().
					GetCallback(gomock.Any(), database.Limit, 0).
					Return(callback, nil)

				mockLogger.EXPECT().
					Info("got initial callback: %s", callback).AnyTimes()

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(nil).AnyTimes()
			},
			expectedError: false,
		},
		{
			name: "Error Getting Callback",
			setupMocks: func(c *service.CMS) {
				mockStorage.EXPECT().
					GetCallback(gomock.Any(), database.Limit, 0).
					Return(nil, errors.New("get callback error"))
				mockLogger.EXPECT().
					Error(gomock.Any(), gomock.Any())
			},
			expectedError: true,
		},
		{
			name: "Error Creating Response",
			setupMocks: func(c *service.CMS) {
				callback := &database.Callback{ID: 1, Name: "test", Phone: "test", Idea: "test", Type: "test"}

				mockStorage.EXPECT().
					GetCallback(gomock.Any(), database.Limit, 0).
					Return(callback, nil)

				mockLogger.EXPECT().
					Info("got initial callback: %s", callback)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(errors.New("publish error"))

				mockLogger.EXPECT().
					Error("error marshalling callback: %s", "publish error")
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &service.CMS{
				Storage: mockStorage,
				Logger:  mockLogger,
			}

			tt.setupMocks(c)

			err := c.InitialCallbackHandler(context.Background())
			if (err != nil) != tt.expectedError {
				t.Errorf("InitialCallbackHandler() error = %v, wantErr %v", err, tt.expectedError)
			}
		})
	}
}

func TestNextCallbackHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocksdb.NewMockICallback(ctrl)
	mockLogger := mocksbroker.NewMockILogger(ctrl)
	mockBroker := mocksbroker.NewMockIBroker(ctrl)

	tests := []struct {
		name          string
		total         int
		offset        int
		setupMocks    func(*service.CMS)
		expectedError bool
	}{
		{
			name:   "Success",
			total:  10,
			offset: 1,
			setupMocks: func(c *service.CMS) {
				callback := &database.Callback{ID: 1, Name: "test", Phone: "test", Idea: "test", Type: "test"}

				mockStorage.EXPECT().
					GetTotalCallbacks(gomock.Any()).
					Return(10, nil)

				mockStorage.EXPECT().
					GetCallback(gomock.Any(), database.Limit, 1).
					Return(callback, nil)

				mockLogger.EXPECT().
					Info("got next callback: %s", callback)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(nil)
			},
			expectedError: false,
		},
		{
			name:   "Error Getting Total Callbacks",
			total:  0,
			offset: 0,
			setupMocks: func(c *service.CMS) {
				mockStorage.EXPECT().
					GetTotalCallbacks(gomock.Any()).
					Return(0, errors.New("get total callbacks error"))

				mockLogger.EXPECT().
					Error("error getting total callbacks: %s", "get total callbacks error")

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(nil).AnyTimes()
			},
			expectedError: true,
		},
		{
			name:   "Error Getting Callback",
			total:  10,
			offset: 1,
			setupMocks: func(c *service.CMS) {
				mockStorage.EXPECT().
					GetTotalCallbacks(gomock.Any()).
					Return(10, nil)

				mockStorage.EXPECT().
					GetCallback(gomock.Any(), database.Limit, 1).
					Return(nil, errors.New("get callback error"))

				mockLogger.EXPECT().
					Error("error fetching callback: %s", "get callback error")

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(nil).AnyTimes()
			},
			expectedError: true,
		},
		{
			name:   "Error Creating Response",
			total:  10,
			offset: 1,
			setupMocks: func(c *service.CMS) {
				callback := &database.Callback{ID: 1, Name: "test", Phone: "test", Idea: "test", Type: "test"}

				mockStorage.EXPECT().
					GetTotalCallbacks(gomock.Any()).
					Return(10, nil)

				mockStorage.EXPECT().
					GetCallback(gomock.Any(), database.Limit, 1).
					Return(callback, nil)

				mockLogger.EXPECT().
					Info("got next callback: %s", callback)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(errors.New("publish error"))

				mockLogger.EXPECT().
					Error("error processing callback: %s", "publish error")
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &service.CMS{
				Storage: mockStorage,
				Logger:  mockLogger,
				Broker:  mockBroker,
			}

			tt.setupMocks(c)

			err := c.NextCallbackHandler(context.Background())
			if (err != nil) != tt.expectedError {
				t.Errorf("NextCallbackHandler() error = %v, wantErr %v", err, tt.expectedError)
			}
		})
	}
}

func TestPreviousCallbackHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocksdb.NewMockICallback(ctrl)
	mockLogger := mocksbroker.NewMockILogger(ctrl)
	mockBroker := mocksbroker.NewMockIBroker(ctrl)

	tests := []struct {
		name          string
		offset        int
		setupMocks    func(*service.CMS)
		expectedError bool
	}{
		{
			name:   "Success",
			offset: 1,
			setupMocks: func(c *service.CMS) {
				callback := &database.Callback{ID: 1, Name: "test", Phone: "test", Idea: "test", Type: "test"}

				mockStorage.EXPECT().
					GetCallback(gomock.Any(), database.Limit, 0).
					Return(callback, nil)

				mockLogger.EXPECT().
					Info("got previous callback: %s", callback)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(nil)
			},
			expectedError: false,
		},
		{
			name:   "Error Getting Callback",
			offset: 1,
			setupMocks: func(c *service.CMS) {
				mockStorage.EXPECT().
					GetCallback(gomock.Any(), database.Limit, 0).
					Return(nil, errors.New("get callback error"))

				mockLogger.EXPECT().
					Error("error fetching callback: %s", "get callback error")

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(nil)
			},
			expectedError: true,
		},
		{
			name:   "Error Creating Response",
			offset: 1,
			setupMocks: func(c *service.CMS) {
				callback := &database.Callback{ID: 1, Name: "test", Phone: "test", Idea: "test", Type: "test"}

				mockStorage.EXPECT().
					GetCallback(gomock.Any(), database.Limit, 0).
					Return(callback, nil)

				mockLogger.EXPECT().
					Info("got previous callback: %s", callback)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(errors.New("publish error"))

				mockLogger.EXPECT().
					Error("error marshalling callback: %s", "publish error")
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &service.CMS{
				Storage: mockStorage,
				Logger:  mockLogger,
				Broker:  mockBroker,
			}

			tt.setupMocks(c)

			err := c.PreviousCallbackHandler(context.Background())
			if (err != nil) != tt.expectedError {
				t.Errorf("PreviousCallbackHandler() error = %v, wantErr %v", err, tt.expectedError)
			}
		})
	}
}

func TestDeleteCallbackHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := mocksdb.NewMockICallback(ctrl)
	mockLogger := mocksbroker.NewMockILogger(ctrl)
	mockBroker := mocksbroker.NewMockIBroker(ctrl)

	tests := []struct {
		name          string
		setupMocks    func(*service.CMS)
		expectedError bool
	}{
		{
			name: "Success",
			setupMocks: func(c *service.CMS) {
				callback := &database.Callback{ID: 1, Name: "test", Phone: "test", Idea: "test", Type: "test"}

				mockStorage.EXPECT().
					DeleteCallbackByID(gomock.Any(), 1).
					Return(nil)

				mockLogger.EXPECT().
					Info("deleted callback: %s", callback).Times(1)
			},
			expectedError: false,
		},
		{
			name: "Error Deleting Callback",
			setupMocks: func(c *service.CMS) {
				mockStorage.EXPECT().
					DeleteCallbackByID(gomock.Any(), 1).
					Return(errors.New("delete callback error"))

				mockLogger.EXPECT().
					Error("error deleting callback: %s", "delete callback error")

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(nil)
			},
			expectedError: true,
		},
		{
			name: "Error Publishing Deletion",
			setupMocks: func(c *service.CMS) {
				callback := &database.Callback{ID: 1, Name: "test", Phone: "test", Idea: "test", Type: "test"}

				mockStorage.EXPECT().
					DeleteCallbackByID(gomock.Any(), 1).
					Return(nil)

				mockLogger.EXPECT().
					Info("callback deleted: %s", callback)

				mockBroker.EXPECT().
					Publish(gomock.Any()).
					Return(errors.New("publish error"))

				mockLogger.EXPECT().
					Error("error marshalling callback: %s", "publish error")
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &service.CMS{
				Storage: mockStorage,
				Logger:  mockLogger,
				Broker:  mockBroker,
			}

			tt.setupMocks(c)

			err := c.DeleteCallbackHandler(context.Background(), amqp.Delivery{})
			if (err != nil) != tt.expectedError {
				t.Errorf("DeleteCallbackHandler() error = %v, wantErr %v", err, tt.expectedError)
			}
		})
	}
}
