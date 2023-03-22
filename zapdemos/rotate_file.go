package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	lumberjacklogger := &lumberjack.Logger{
		Filename:   "./log-rotate-test.json",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
	defer lumberjacklogger.Close()

	config := zap.NewProductionEncoderConfig()

	config.EncodeTime = zapcore.ISO8601TimeEncoder // 设置时间格式
	fileEncoder := zapcore.NewJSONEncoder(config)

	core := zapcore.NewCore(
		fileEncoder,                       //编码设置
		zapcore.AddSync(lumberjacklogger), //输出到文件
		zap.InfoLevel,                     //日志等级
	)

	logger := zap.New(core)
	defer logger.Sync()

	// 测试分割日志
	for i := 0; i < 8000; i++ {
		logger.With(
			zap.String("url", fmt.Sprintf("www.test%d.com", i)),
			zap.String("name", "jimmmyr"),
			zap.Int("age", 23),
			zap.String("agradege", "no111-000222"),
		).Info("test info ")
	}

}
