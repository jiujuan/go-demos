package main

import (
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	logger.Debug("this is debug message")
	logger.Info("this is info message")
	logger.Info("this is info message with fileds",
		zap.Int("age", 37),
		zap.String("agender", "man"),
	)
	logger.Warn("this is warn message")
	logger.Error("this is error message")
	// logger.Panic("this is panic message")
}
