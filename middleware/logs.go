package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ix-pay/ixpay/utils"
	"github.com/sirupsen/logrus"
)

func Logs() gin.HandlerFunc {
	logger := utils.SetupLogrus()
	return func(c *gin.Context) {
		id := uuid.New()
		// 集成Logrus日志库，支持结构化JSON格式日志和更灵活的日志配置
		logger.WithFields(logrus.Fields{
			"id":     id,
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		}).Info("请求已接收")
		c.Next()
		logger.WithFields(logrus.Fields{
			"id":         id,
			"statusCode": c.Writer.Status(),
		}).Info("请求已处理")
	}
}
