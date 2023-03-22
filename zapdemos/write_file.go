package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	writetofile()
}

func writetofile() {
	// 设置一些配置参数
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	defaultLogLevel := zapcore.DebugLevel // 设置 loglevel

	logFile, _ := os.OpenFile("./log-test-zap.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 06666)
	// or os.Create()
	writer := zapcore.AddSync(logFile)

	logger := zap.New(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	defer logger.Sync()

	url := "http://www.test.com"
	logger.Info("write log to file",
		zap.String("url", url),
		zap.Int("attemp", 3),
	)
}
