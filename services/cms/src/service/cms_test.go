package service_test

import (
	"callback_service/src/database"
	"callback_service/src/service"
	mocks "callback_service/src/service/mocks"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
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

func TestHandleMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBroker := mocks.NewMockIBroker(ctrl)
	mockStorage := mocks.NewMockICallback(ctrl)
	mockLogger := mocks.NewMockILogger(ctrl)

	cms := service.NewCMS(mockBroker, mockStorage, mockLogger)

	tests := []struct {
		name       string
		body       []byte
		setupMocks func()
		wantErr    bool
	}{
		{
			name: "Successful case",
			body: []byte(`{"Name":"testName","Phone":"123456","Type":"testType","Idea":"testIdea"}`),
			setupMocks: func() {
				mockStorage.EXPECT().CreateCallback(gomock.Any(), gomock.Any()).Return(nil)
				mockBroker.EXPECT().Produce(gomock.Any(), "success", []byte("Callback created successfully")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Validation error",
			body: []byte(`{"Name":testName,"Phone":"123456","Type":"testType","Idea":"testIdea"}`),
			setupMocks: func() {
				mockLogger.EXPECT().Error("error during validation or conversion: %s", gomock.Any())
				mockBroker.EXPECT().Produce(gomock.Any(), "error", gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Storage error",
			body: []byte(`{"Name":"testName","Phone":"123456","Type":"testType","Idea":"testIdea"}`),
			setupMocks: func() {
				mockStorage.EXPECT().CreateCallback(gomock.Any(), gomock.Any()).Return(errors.New("insert error"))
				mockLogger.EXPECT().Error("error inserting callback: %s", gomock.Any())
				mockBroker.EXPECT().Produce(gomock.Any(), "error", gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Broker produce error on success",
			body: []byte(`{"Name":"testName","Phone":"123456","Type":"testType","Idea":"testIdea"}`),
			setupMocks: func() {
				mockStorage.EXPECT().CreateCallback(gomock.Any(), gomock.Any()).Return(nil)
				mockBroker.EXPECT().Produce(gomock.Any(), "success", []byte("Callback created successfully")).Return(errors.New("produce error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			err := cms.HandleMessage(context.Background(), tt.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
