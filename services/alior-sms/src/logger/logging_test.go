package logger_test

import (
	"testing"

	"alior-sms/src/logger"
	mocks "alior-sms/src/logger/mocks"

	"github.com/golang/mock/gomock"
)

func TestLoggerMethods(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := mocks.NewMockILogger(ctrl)

	tests := []struct {
		name      string
		logMethod func(l logger.ILogger, msg string, args ...interface{})
		expect    func()
		msg       string
		args      []interface{}
	}{
		{
			name: "Debug",
			logMethod: func(l logger.ILogger, msg string, args ...interface{}) {
				msg = logger.FormatMessage("debug message: %v, %v", "arg1", "arg2")
				l.Debug(msg, args...)
			},
			expect: func() {
				msg := logger.FormatMessage("debug message: %v, %v", "arg1", "arg2")
				mockLogger.EXPECT().Debug(msg, "arg1", "arg2")
			},
			msg:  "debug message: %v, %v",
			args: []interface{}{"arg1", "arg2"},
		},
		{
			name: "Info",
			logMethod: func(l logger.ILogger, msg string, args ...interface{}) {
				msg = logger.FormatMessage("info message: %v, %v", "arg1", "arg2")
				l.Info(msg, args...)
			},
			expect: func() {
				msg := logger.FormatMessage("info message: %v, %v", "arg1", "arg2")
				mockLogger.EXPECT().Info(msg, "arg1", "arg2")
			},
			msg:  "info message: %v, %v",
			args: []interface{}{"arg1", "arg2"},
		},
		{
			name: "Warn",
			logMethod: func(l logger.ILogger, msg string, args ...interface{}) {
				msg = logger.FormatMessage("warn message: %v, %v", "arg1", "arg2")
				l.Warn(msg, args...)
			},
			expect: func() {
				msg := logger.FormatMessage("warn message: %v, %v", "arg1", "arg2")
				mockLogger.EXPECT().Warn(msg, "arg1", "arg2")
			},
			msg:  "warn message: %v, %v",
			args: []interface{}{"arg1", "arg2"},
		},
		{
			name: "Error",
			logMethod: func(l logger.ILogger, msg string, args ...interface{}) {
				msg = logger.FormatMessage("error message: %v, %v", "arg1", "arg2")
				l.Error(msg, args...)
			},
			expect: func() {
				msg := logger.FormatMessage("error message: %v, %v", "arg1", "arg2")
				mockLogger.EXPECT().Error(msg, "arg1", "arg2")
			},
			msg:  "error message: %v, %v",
			args: []interface{}{"arg1", "arg2"},
		},
		{
			name: "Fatal",
			logMethod: func(l logger.ILogger, msg string, args ...interface{}) {
				msg = logger.FormatMessage("fatal message: %v, %v", "arg1", "arg2")
				l.Fatal(msg, args...)
			},
			expect: func() {
				msg := logger.FormatMessage("fatal message: %v, %v", "arg1", "arg2")
				mockLogger.EXPECT().Fatal(msg, "arg1", "arg2")
			},
			msg:  "fatal message: %v, %v",
			args: []interface{}{"arg1", "arg2"},
		},

		{
			name: "NilMessageDebug",
			logMethod: func(l logger.ILogger, msg string, args ...interface{}) {
				l.Debug(msg, args...)
			},
			expect: func() {
				mockLogger.EXPECT().Debug("", "arg1", "arg2").Times(1)
			},
			msg:  "",
			args: []interface{}{"arg1", "arg2"},
		},
		{
			name: "IncorrectArgsInfo",
			logMethod: func(l logger.ILogger, msg string, args ...interface{}) {
				l.Info(msg, args...)
			},
			expect: func() {
				mockLogger.EXPECT().Info("info message: %v, %v", "один аргумент у женщины").Times(1)
			},
			msg:  "info message: %v, %v",
			args: []interface{}{"один аргумент у женщины"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.expect()
			tt.logMethod(mockLogger, tt.msg, tt.args...)
		})
	}
}
