package database_test

import (
	"callback_service/src/database"
	databasetest "callback_service/src/database/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := databasetest.NewMockICallback(ctrl)

	tests := []struct {
		name    string
		input   *database.Callback
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			input: &database.Callback{
				Name:  "John Weak",
				Phone: "+1234567890",
				Type:  "top rovniy",
				Idea:  "idea norm",
			},
			wantErr: false,
			mock: func() {
				mockDB.EXPECT().CreateCallback(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "database error",
			input: &database.Callback{
				Name:  "Vadim Kisliy",
				Phone: "+1234567890",
				Type:  "tip rovniy",
				Idea:  "idea tak sebe",
			},
			wantErr: true,
			mock: func() {
				mockDB.EXPECT().CreateCallback(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := mockDB.CreateCallback(context.Background(), tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
