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

func FormatMessage(msg string, args ...interface{}) string {
	if len(args) > 0 {
		return fmt.Sprintf(msg, args...)
	}
	return msg
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.logger.Debug(FormatMessage(msg, args...))
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.logger.Info(FormatMessage(msg, args...))
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.logger.Warn(FormatMessage(msg, args...))
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.logger.Error(FormatMessage(msg, args...))
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatal(FormatMessage(msg, args...))
}
