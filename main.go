package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/config"
	"github.com/ix-pay/ixpay/models"
	"github.com/ix-pay/ixpay/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ix-pay/ixpay/docs"
)

// @title API文档
// @version 1.0
// @description ixpay项目Swagger集成
// @host localhost:8989
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	// 初始化配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("无法加载配置:", err)
	}

	// 初始化数据库
	err = models.InitDB(cfg.DB)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 初始化雪花算法节点（节点号0-1023）
	if err := models.InitSnowflake(cfg); err != nil {
		panic(err)
	}

	// 加载路由
	r := routes.SetupRoutes()

	// 开发环境启用文档
	if gin.Mode() != gin.ReleaseMode {
		url := ginSwagger.URL("/swagger/doc.json")
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	// 启动服务器
	r.Run(":8989")
}
