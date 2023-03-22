package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "http://zap.uber.io"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"time", time.Second,
	)

	sugar.Infof("Failed to fetch URL: %s", url)

	// 或更简洁 Sugar() 使用
	// sugar := zap.NewProduction().Sugar()
	// defer sugar.Sync()
}
