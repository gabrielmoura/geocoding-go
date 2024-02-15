package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	Logger, _ = config.Build()
	defer Logger.Sync()
}
