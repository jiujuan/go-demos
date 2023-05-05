package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction(zap.AddCaller())
	defer logger.Sync()

	logger.Info("AddCaller：line No and filename")
}
