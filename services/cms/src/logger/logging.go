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

func NewZapLogger() *Logger {
	logger, _ := SetupLogger()
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
	l.logger.Debug(formatMessage(msg, args...))
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.Info(formatMessage(msg, args...))
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.logger.Warn(formatMessage(msg, args...))
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.logger.Error(formatMessage(msg, args...))
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatal(formatMessage(msg, args...))
}
