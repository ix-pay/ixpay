package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.AbortError(c, http.StatusUnauthorized, "未提供认证token")
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证token"})
			return
		}
		log.Printf("tokenString=%s\n", tokenString)
		// 3. 安全处理Bearer前缀
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.AbortError(c, http.StatusUnauthorized, "token格式错误")
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token格式错误"})
			return
		}
		realToken := tokenParts[1]

		// 4. 解析前二次验证
		log.Printf("待解析的realToken=%s\n", realToken)
		token, err := utils.ParseJWT(realToken)

		if err != nil {
			log.Printf("解析错误详情: %v\n", err) // 关键错误日志
			utils.AbortError(c, http.StatusUnauthorized, "token解析失败")
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token解析失败"})
			return
		}

		// 5. 验证有效性
		if !token.Valid {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					utils.AbortError(c, http.StatusUnauthorized, "token已过期")
					// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token已过期"})
					return
				}
			}
			utils.AbortError(c, http.StatusUnauthorized, "无效token")
			// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效token"})
			return
		}

		c.Next()
	}
}
