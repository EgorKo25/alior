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

func TestGetCallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := databasetest.NewMockICallback(ctrl)

	tests := []struct {
		name    string
		limit   int
		offset  int
		want    *database.Callback
		wantErr bool
		mock    func()
	}{
		{
			name:   "success with offset = 0",
			limit:  1,
			offset: 0,
			want: &database.Callback{
				ID:    1,
				Name:  "John Weak",
				Phone: "+1234567890",
				Type:  "top rovniy",
				Idea:  "idea norm",
			},
			wantErr: false,
			mock: func() {
				mockDB.EXPECT().GetCallback(gomock.Any(), 1, 0).Return(&database.Callback{
					ID:    1,
					Name:  "John Weak",
					Phone: "+1234567890",
					Type:  "top rovniy",
					Idea:  "idea norm",
				}, nil)
			},
		},
		{
			name:    "database error",
			limit:   1,
			offset:  0,
			want:    nil,
			wantErr: true,
			mock: func() {
				mockDB.EXPECT().GetCallback(gomock.Any(), 1, 0).Return(nil, errors.New("database error"))
			},
		},
		{
			name:    "offset exceeds total callbacks",
			limit:   1,
			offset:  5,
			want:    nil,
			wantErr: true,
			mock: func() {
				mockDB.EXPECT().GetCallback(gomock.Any(), 1, 5).Return(nil, errors.New("offset out of range"))
			},
		},
		{
			name:    "no callbacks available",
			limit:   1,
			offset:  0,
			want:    nil,
			wantErr: true,
			mock: func() {
				mockDB.EXPECT().GetCallback(gomock.Any(), 1, 0).Return(nil, errors.New("no records found"))
			},
		},
		{
			name:    "only one callback available, offset = 1",
			limit:   1,
			offset:  1,
			want:    nil,
			wantErr: true,
			mock: func() {
				mockDB.EXPECT().GetCallback(gomock.Any(), 1, 1).Return(nil, errors.New("offset out of range"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := mockDB.GetCallback(context.Background(), tt.limit, tt.offset)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestGetTotalCallbacks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := databasetest.NewMockICallback(ctrl)

	tests := []struct {
		name    string
		want    int
		wantErr bool
		mock    func()
	}{
		{
			name:    "success",
			want:    10,
			wantErr: false,
			mock: func() {
				mockDB.EXPECT().GetTotalCallbacks(gomock.Any()).Return(10, nil)
			},
		},
		{
			name:    "database error",
			want:    0,
			wantErr: true,
			mock: func() {
				mockDB.EXPECT().GetTotalCallbacks(gomock.Any()).Return(0, errors.New("database error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := mockDB.GetTotalCallbacks(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDeleteCallbackByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := databasetest.NewMockICallback(ctrl)

	tests := []struct {
		name    string
		id      int32
		wantErr bool
		mock    func()
	}{
		{
			name:    "success",
			id:      1,
			wantErr: false,
			mock: func() {
				mockDB.EXPECT().DeleteCallbackByID(gomock.Any(), int32(1)).Return(nil)
			},
		},
		{
			name:    "object not found",
			id:      2,
			wantErr: true,
			mock: func() {
				mockDB.EXPECT().DeleteCallbackByID(gomock.Any(), int32(2)).Return(errors.New("object not found"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := mockDB.DeleteCallbackByID(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
