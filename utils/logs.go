package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogrus() *logrus.Logger {
	log := logrus.New()
	// file, err := os.OpenFile("logs/req.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }

	// 配置lumberjack日志轮转
	file := &lumberjack.Logger{
		Filename:   "logs/req.log",
		MaxSize:    100,  // MB
		MaxBackups: 300,  // 保留旧文件数量
		MaxAge:     30,   // 保留天数
		Compress:   true, // 压缩旧日志
	}
	log.Out = file

	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}
