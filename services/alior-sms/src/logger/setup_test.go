package logger_test

import (
	"os"
	"testing"

	"alior-sms/src/logger"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetupLogger(t *testing.T) {
	t.Run("successful logger setup", func(t *testing.T) {
		cleanupLogs()
		defer cleanupLogs()

		zapLogger, err := logger.SetupLogger()
		require.NoError(t, err)
		assert.NotNil(t, zapLogger)
	})

	t.Run("error creating log directory", func(t *testing.T) {
		file, err := os.Create("logs")
		require.NoError(t, err)
		defer file.Close()
		defer os.Remove("logs")

		zapLogger, err := logger.SetupLogger()

		require.Error(t, err)
		assert.Nil(t, zapLogger)
		assert.Contains(t, err.Error(), "not a directory")
	})

	t.Run("error creating log files", func(t *testing.T) {
		os.MkdirAll("logs", 0000)
		defer os.Chmod("logs", os.ModePerm)
		defer cleanupLogs()

		zapLogger, err := logger.SetupLogger()
		require.Error(t, err)
		assert.Nil(t, zapLogger)
	})
}

func cleanupLogs() {
	os.RemoveAll("logs")
}
