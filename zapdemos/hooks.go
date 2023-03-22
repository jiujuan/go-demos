package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger := zap.NewExample(zap.Hooks(func(entry zapcore.Entry) error {
		fmt.Println("[zap.Hooks]test Hooks")
		return nil
	}))
	defer logger.Sync()

	logger.Info("test output")

	logger.Warn("warn info")
}
