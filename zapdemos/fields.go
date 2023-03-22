package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction(zap.Fields(
		zap.String("log_name", "testlog"),
		zap.String("log_author", "prometheus"),
	))
	defer logger.Sync()

	logger.Info("test fields output")

	logger.Warn("warn info")
}
