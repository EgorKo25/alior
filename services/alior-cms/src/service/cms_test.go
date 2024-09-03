package service_test

import (
	"callback_service/src/broker"
	brokertest "callback_service/src/broker/mocks"
	"callback_service/src/database"
	"callback_service/src/service"
	loggertest "callback_service/src/service/mocks"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConvertToRepositoryAndValidate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name          string
		input         []byte
		expected      *database.Callback
		expectedError error
	}{
		{
			name: "success",
			input: []byte(`{
				"name": "John Doe",
				"phone": "1234567890",
				"type": "Inquiry",
				"idea": "New project idea"
			}`),
			expected: &database.Callback{
				Name:      "John Doe",
				Phone:     "1234567890",
				Type:      "Inquiry",
				Idea:      "New project idea",
				CreatedAt: time.Now(),
			},
			expectedError: nil,
		},
		{
			name:          "invalid JSON",
			input:         []byte(`invalid-json`),
			expected:      nil,
			expectedError: errors.New("invalid character 'i' looking for beginning of value"),
		},
		{
			name: "missing required fields",
			input: []byte(`{
				"name": "John Doe",
				"phone": "1234567890"
			}`),
			expected:      nil,
			expectedError: errors.New("one or more required fields are empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ConvertToRepositoryAndValidate(tt.input)

			if tt.expectedError != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.NotNil(t, result)
				require.Equal(t, tt.expected.Name, result.Name)
				require.Equal(t, tt.expected.Phone, result.Phone)
				require.Equal(t, tt.expected.Type, result.Type)
				require.Equal(t, tt.expected.Idea, result.Idea)
				require.WithinDuration(t, tt.expected.CreatedAt, result.CreatedAt, time.Second)
			}
		})
	}
}

func compareCallbacks(a, b *database.Callback) bool {
	aj, _ := json.Marshal(a)
	bj, _ := json.Marshal(b)
	return string(aj) == string(bj)
}

func TestCreateResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBroker := brokertest.NewMockIBroker(ctrl)
	mockLogger := loggertest.NewMockILogger(ctrl)

	cms := &service.CMS{
		Broker: mockBroker,
		Logger: mockLogger,
	}

	tests := []struct {
		name       string
		input      *database.Callback
		setupMocks func()
		wantErr    bool
	}{
		{
			name: "success",
			input: &database.Callback{
				Name:  "John Doe",
				Phone: "+1234567890",
				Type:  "Inquiry",
				Idea:  "New Feature",
			},
			setupMocks: func() {
				callbackJSON, _ := json.Marshal(&database.Callback{
					Name:  "John Doe",
					Phone: "+1234567890",
					Type:  "Inquiry",
					Idea:  "New Feature",
				})
				msg := broker.NewMessage(string(callbackJSON), "callback")
				mockBroker.EXPECT().Publish(msg).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "broker publish error",
			input: &database.Callback{
				Name:  "John Doe",
				Phone: "+1234567890",
				Type:  "Inquiry",
				Idea:  "New Feature",
			},
			setupMocks: func() {
				callbackJSON, _ := json.Marshal(&database.Callback{
					Name:  "John Doe",
					Phone: "+1234567890",
					Type:  "Inquiry",
					Idea:  "New Feature",
				})
				msg := broker.NewMessage(string(callbackJSON), "callback")
				mockBroker.EXPECT().Publish(msg).Return(errors.New("broker error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			err := cms.CreateResponse(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
