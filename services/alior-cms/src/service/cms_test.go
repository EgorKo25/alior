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
	"testing"
)

func TestConvertToRepositoryAndValidate(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		wantErr  bool
		expected *database.Callback
	}{
		{
			name:    "Valid JSON",
			input:   []byte(`{"Name":"Name","Phone":"Phone", "Type":"Type","Idea":"Idea"}`),
			wantErr: false,
			expected: &database.Callback{
				Name:  "Name",
				Phone: "Phone",
				Type:  "Type",
				Idea:  "Idea",
			},
		},
		{
			name:    "Invalid JSON",
			input:   []byte(`{"field1":value1,"field2":"value2"}`),
			wantErr: true,
		},
		{
			name:     "Empty JSON",
			input:    []byte(`{}`),
			wantErr:  false,
			expected: &database.Callback{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ConvertToRepositoryAndValidate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToRepositoryAndValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !compareCallbacks(result, tt.expected) {
				t.Errorf("convertToRepositoryAndValidate() = %v, expected %v", result, tt.expected)
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
