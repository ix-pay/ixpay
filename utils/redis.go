package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/ix-pay/ixpay/config"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
)

func SetupRedis(cfg *config.RedisConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// 测试连接
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatal("Redis连接失败: ", err)
	}
	log.Println("Redis连接成功")
	return client
}
