package config_test

import (
	"errors"
	"testing"

	"callback_service/src/config"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name        string
		mockSetup   func()
		expectedCfg *config.Config
		expectedErr error
	}{
		{
			name: "Successful Load",
			mockSetup: func() {
				config.ReadConfig = func(path string, cfg interface{}) error {
					c := cfg.(*config.Config)
					c.Database.URL = "postgresql://user:password@localhost:5432/dbname"
					c.MsgBroker.URL = "amqp://guest:guest@localhost:5672/"
					return nil
				}
			},
			expectedCfg: &config.Config{
				Database: config.DatabaseConfig{
					URL: "postgresql://user:password@localhost:5432/dbname",
				},
				MsgBroker: config.MsgBrokerConfig{
					URL: "amqp://guest:guest@localhost:5672/",
				},
			},
			expectedErr: nil,
		},
		{
			name: "Failed Load",
			mockSetup: func() {
				config.ReadConfig = func(path string, cfg interface{}) error {
					return errors.New("file not found")
				}
			},
			expectedCfg: nil,
			expectedErr: errors.New("failed to load config: file not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			cfg, err := config.Load()
			assert.Equal(t, tt.expectedCfg, cfg)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
