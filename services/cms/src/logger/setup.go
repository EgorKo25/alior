package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func SetupLogger() (*zap.Logger, error) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		return nil, err
	}

	debugFile, err := os.Create("logs/debug.log")
	if err != nil {
		return nil, err
	}
	errorFile, err := os.Create("logs/error.log")
	if err != nil {
		return nil, err
	}
	infoFile, err := os.Create("logs/info.log")
	if err != nil {
		return nil, err
	}

	debugWS := zapcore.AddSync(debugFile)
	errorWS := zapcore.AddSync(errorFile)
	infoWS := zapcore.AddSync(infoFile)

	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zap.DebugLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zap.InfoLevel
	})

	consoleDebugging := zapcore.Lock(os.Stderr)

	encoder := zapcore.NewJSONEncoder(encoderCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, consoleDebugging, zapcore.DebugLevel),
		zapcore.NewCore(encoder, debugWS, debugLevel),
		zapcore.NewCore(encoder, errorWS, errorLevel),
		zapcore.NewCore(encoder, infoWS, infoLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger, nil
}
