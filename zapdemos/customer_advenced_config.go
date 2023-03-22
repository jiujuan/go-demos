package main

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 首先，定义不同级别日志处理逻辑
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	// 假设有2个kafka 的 topic，一个 debugging，一个 errors

	// zapcore.AddSync 添加一个文件句柄。
	topicDebugging := zapcore.AddSync(io.Discard)
	topicErrors := zapcore.AddSync(io.Discard)

	// 如果他们对并发使用不安全，我们可以用 zapcore.Lock 添加一个 mutex 互斥锁。
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// 设置 kafka 和 console 输出配置
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// 把上面的设置加入到 zapcore.NewCore() 函数里，然后再把他们加入到 zapcore.NewTee() 函数里
	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	// 最后调用 zap.New() 函数
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("constructed a logger")
}
