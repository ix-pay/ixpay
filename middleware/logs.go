package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/utils"
	"github.com/sirupsen/logrus"
)

func Logs() gin.HandlerFunc {
	logger := utils.SetupLogrus()
	return func(c *gin.Context) {
		// 集成Logrus日志库，支持结构化JSON格式日志和更灵活的日志配置
		logger.WithFields(logrus.Fields{
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		}).Info("Request received")
		c.Next()
	}
}
