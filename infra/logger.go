package infra

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.Logger
}

func NewLogger() *zapLogger {

	cfg := zap.NewDevelopmentConfig()

	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	log, err := cfg.Build()

	if err != nil {
		fmt.Printf("log: %v\n", err.Error())
	}

	return &zapLogger{
		logger: log,
	}
}

func (lg zapLogger) GetLogger() *zap.Logger {

	return lg.logger

}
