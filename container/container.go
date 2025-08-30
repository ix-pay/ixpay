package container

import (
	"log"
	"sync"

	"github.com/ix-pay/ixpay/config"
	"github.com/ix-pay/ixpay/service"
	"github.com/ix-pay/ixpay/utils"
)

type Container struct {
	mu             sync.Mutex
	initialized    bool
	cfg            *config.Config
	j              *utils.JwtUtil
	userService    service.UserService
	paymentService service.PaymentService
}

func (c *Container) Init() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.initialized {
		return
	}
	// 初始化配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("无法加载配置:", err)
		return
	}
	c.cfg = cfg

	// 初始化JWT
	c.j = utils.SetupJwt(cfg)

	// 初始化服务层
	c.userService = service.NewUserService()
	c.paymentService = service.NewPaymentService()
	c.initialized = true
}

// 获取配置的单例方法
func (c *Container) GetConfig() *config.Config {
	return c.cfg
}

// 获取JWT的单例方法
func (c *Container) GetJwt() *utils.JwtUtil {
	return c.j
}

// 获取服务方法（线程安全）
func (c *Container) GetUserService() service.UserService {
	return c.userService
}

func (c *Container) GetPaymentService() service.PaymentService {
	return c.paymentService
}
