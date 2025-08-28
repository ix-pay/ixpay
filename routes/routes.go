package routes

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/middleware"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupRoutes() *gin.Engine {
	// 初始化Gin引擎
	r := gin.Default()

	// 创建日志文件
	// file, _ := os.Create("logs/ixpay.log")
	// file, _ := os.OpenFile("logs/ixpay.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// 配置lumberjack日志轮转
	file := &lumberjack.Logger{
		Filename:   "logs/ixpay.log",
		MaxSize:    100,  // MB
		MaxBackups: 300,  // 保留旧文件数量
		MaxAge:     30,   // 保留天数
		Compress:   true, // 压缩旧日志
	}

	// 设置日志输出目标（文件+控制台）
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	// 设置日志输出文件位置
	log.SetOutput(file)

	// 注册中间件
	r.Use(middleware.CORS())
	// 注册中间件
	r.Use(middleware.Logs())

	api := r.Group("/api/v1")
	{

		SetupAuthRoutes(api)
		SetupPaymentsRoutes(api)
	}

	return r
}
