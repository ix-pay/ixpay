package main

import (
	"log"

	"github.com/ix-pay/ixpay/app"

	_ "github.com/ix-pay/ixpay/docs"
)

// @title API文档
// @version 1.0
// @description ixpay项目Swagger集成
// @host localhost:8989
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header [Bearer ]
// @name Authorization
func main() {

	// 构建应用
	application, err := app.SetupApp()
	if err != nil {
		log.Fatal("应用初始化失败:", err)
	}

	// 启动服务
	if err := application.Run(); err != nil {
		log.Fatal("服务运行失败:", err)
	}
}
