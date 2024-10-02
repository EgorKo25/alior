package database_test

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/stdlib"
	"testing"

	databasetest "callback_service/src/database/mocks"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func TestRunMigrations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name      string
		setupMock func(m *databasetest.MockIMigrator)
		wantErr   bool
	}{
		{
			name: "Migrations succeed",
			setupMock: func(m *databasetest.MockIMigrator) {
				m.EXPECT().Up(gomock.Any(), "migrations").Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Migrations fail",
			setupMock: func(m *databasetest.MockIMigrator) {
				m.EXPECT().Up(gomock.Any(), "migrations").Return(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pool, err := pgxpool.New(context.Background(), "postgres://user:password@localhost:5432/db?sslmode=disable")
			assert.NoError(t, err)
			defer pool.Close()

			mockMigrator := databasetest.NewMockIMigrator(ctrl)
			tt.setupMock(mockMigrator)

			err = RunMigrations(pool, mockMigrator)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func RunMigrations(pool *pgxpool.Pool, migrator *databasetest.MockIMigrator) error {
	db := stdlib.OpenDBFromPool(pool)
	return migrator.Up(db, "migrations")
}
