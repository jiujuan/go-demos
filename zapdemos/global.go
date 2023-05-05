package main

import (
	"go.uber.org/zap"
)

func main() {
	// 直接调用是不会记录日志信息的，所以下面日志信息不会输出
	zap.L().Info("no log info")
	zap.S().Info("no log info [sugared]")

	logger := zap.NewExample()
	defer logger.Sync()

	zap.ReplaceGlobals(logger) // 全局logger，zap.L() 和 zap.S() 需要调用 ReplaceGlobals 函数
	zap.L().Info("log info")
	zap.S().Info("log info [sugared]")
}
