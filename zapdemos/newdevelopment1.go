package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Info("failed to fetch url",
		// 强类型字段
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("duration", time.Second),
	)

	logger.With(
		// 强类型字段
		zap.String("url", "http://development.com"),
		zap.Int("attempt", 4),
		zap.Duration("duration", time.Second*5),
	).Info("[With] failed to fetch url")
}
