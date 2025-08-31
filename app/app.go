package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/config"
	"github.com/ix-pay/ixpay/container"
	"github.com/ix-pay/ixpay/models"
	"github.com/ix-pay/ixpay/routes"
	"github.com/ix-pay/ixpay/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	engine *gin.Engine
	cfg    *config.Config
}

func SetupApp() (*App, error) {
	// 初始化容器
	ctr := container.SetupContainer()

	// 注册服务
	ctr.Register(container.AuthServiceName, func() interface{} {
		return service.NewAuthService()
	})
	ctr.Register(container.UserServiceName, func() interface{} {
		return service.NewUserService()
	})
	ctr.Register(container.PaymentServiceName, func() interface{} {
		return service.NewPaymentService()
	})

	cfg := ctr.GetConfig()

	// 初始化数据库
	err := models.SetupDB(cfg.DB)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 初始化雪花算法节点（节点号0-1023）
	if err := models.InitSnowflake(cfg.MachineId); err != nil {
		panic(err)
	}

	// 注册路由
	engine := routes.SetupRoutes(ctr)

	// 开发环境启用文档
	if gin.Mode() != gin.ReleaseMode {
		url := ginSwagger.URL("/swagger/doc.json")
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	return &App{engine: engine, cfg: cfg}, nil
}

func (a *App) Run() error {
	port := a.cfg.ServerPort
	a.engine.SetTrustedProxies(nil)
	return a.engine.Run(":" + port)
}
