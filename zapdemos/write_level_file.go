package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	writeToFileWithLogLevel()
}

func writeToFileWithLogLevel() {
	// 设置配置
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)

	logFile, _ := os.OpenFile("./log-debug-zap.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666) //日志记录debug信息

	errFile, _ := os.OpenFile("./log-err-zap.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666) //日志记录error信息

	teecore := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), zap.DebugLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(errFile), zap.ErrorLevel),
	)

	logger := zap.New(teecore, zap.AddCaller())
	defer logger.Sync()

	url := "http://www.diff-log-level.com"
	logger.Info("write log to file",
		zap.String("url", url),
		zap.Int("time", 3),
	)

	logger.With(
		zap.String("url", url),
		zap.String("name", "jimmmyr"),
	).Error("test error ")
}
