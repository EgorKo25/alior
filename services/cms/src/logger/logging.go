package logger

import (
	"fmt"
	"go.uber.org/zap"
)

type ILogger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

type Logger struct {
	logger *zap.Logger
}

func NewZapLogger(logger *zap.Logger) ILogger {
	return &Logger{
		logger: logger,
	}
}

func formatMessage(msg string, args ...interface{}) string {
	if len(args) > 0 {
		return fmt.Sprintf(msg, args...)
	}
	return msg
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	formattedMsg := formatMessage(msg, args...)
	l.logger.Debug(formattedMsg)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	formattedMsg := formatMessage(msg, args...)
	l.logger.Info(formattedMsg)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	formattedMsg := formatMessage(msg, args...)
	l.logger.Warn(formattedMsg)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	formattedMsg := formatMessage(msg, args...)
	l.logger.Error(formattedMsg)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	formattedMsg := formatMessage(msg, args...)
	l.logger.Fatal(formattedMsg)
}
