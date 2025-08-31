package container

import (
	"errors"
	"log"
	"sync"

	"github.com/ix-pay/ixpay/config"
	"github.com/ix-pay/ixpay/utils"
	"github.com/redis/go-redis/v9"
)

type IContainer interface {
	GetConfig() *config.Config
	GetJwt() *utils.JwtUtil
	GetRedis() *redis.Client
	Register(name string, factory func() interface{})
	Get(name string) (interface{}, error)
	MustGet(name string) interface{}
}

type Container struct {
	mu       sync.RWMutex //读写锁
	cfg      *config.Config
	j        *utils.JwtUtil
	rc       *redis.Client
	services map[string]func() interface{}
}

func SetupContainer() IContainer {
	// 初始化配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("无法加载配置:", err)
		return nil
	}

	// 初始化JWT
	j := utils.SetupJwt(cfg.JWTSecret)

	// 初始化Redis
	rc := utils.SetupRedis(&cfg.Redis)
	return &Container{
		cfg:      cfg,
		j:        j,
		rc:       rc,
		services: make(map[string]func() interface{}),
	}
}

// 获取配置的单例方法
func (c *Container) GetConfig() *config.Config {
	return c.cfg
}

// 获取JWT的单例方法
func (c *Container) GetJwt() *utils.JwtUtil {
	return c.j
}

// 获取Redis的单例方法
func (c *Container) GetRedis() *redis.Client {
	return c.rc
}

// 注册服务
func (c *Container) Register(name string, factory func() interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.services[name] = factory
}

// 获取服务
func (c *Container) Get(name string) (interface{}, error) {
	c.mu.RLock()
	factory, exists := c.services[name]
	c.mu.RUnlock()

	if !exists {
		return nil, errors.New("服务不存在")
	}
	return factory(), nil
}

// 获取服务
func (c *Container) MustGet(name string) interface{} {
	service, err := c.Get(name)
	if err != nil {
		panic(err)
	}
	return service
}
