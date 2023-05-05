package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Hello() {
	Warn("hello", zap.String("h", "world"), zap.Int("c", 1))
}

func Warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}

func main() {
	logger, _ := zap.NewProduction(zap.AddStacktrace(zapcore.WarnLevel))
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	Warn()
}
