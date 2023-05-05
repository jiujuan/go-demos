package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	Hello()
}

func Hello() {
	Warn("hello", zap.String("h", "world"), zap.Int("c", 1))
}

func Warn(msg string, fields ...zap.Field) {
	zap.L().Warn(msg, fields...)
}
