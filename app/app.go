package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/config"
	"github.com/ix-pay/ixpay/container"
	"github.com/ix-pay/ixpay/models"
	"github.com/ix-pay/ixpay/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	engine *gin.Engine
	cfg    *config.Config
}

func SetupApp() (*App, error) {
	// 初始化容器
	ctr := &container.Container{}
	ctr.Init()

	cfg := ctr.GetConfig()

	// 初始化数据库
	err := models.InitDB(cfg.DB)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 初始化雪花算法节点（节点号0-1023）
	if err := models.InitSnowflake(cfg); err != nil {
		panic(err)
	}

	// 创建Gin引擎
	// 加载路由
	engine := routes.SetupRoutes(ctr)

	// 开发环境启用文档
	if gin.Mode() != gin.ReleaseMode {
		url := ginSwagger.URL("/swagger/doc.json")
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	// 注册路由

	return &App{engine: engine, cfg: cfg}, nil
}

func (a *App) Run() error {
	port := a.cfg.ServerPort
	return a.engine.Run(":" + port)
}
