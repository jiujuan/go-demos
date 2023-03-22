package main

import (
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("some message",
		zap.Namespace("shop"),
		zap.String("name", "LiLei"),
		zap.String("grade", "No2"),
	)

	logger.Error("some error message",
		zap.Namespace("shop"),
		zap.String("name", "LiLei"),
		zap.String("grade", "No3"),
	)

}
