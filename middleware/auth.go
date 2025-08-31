package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/container"
	"github.com/ix-pay/ixpay/models"
	"github.com/ix-pay/ixpay/service"
	"github.com/ix-pay/ixpay/utils"
)

func JWTAuth(ctr container.IContainer) gin.HandlerFunc {
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
		token, err := ctr.GetJwt().ParseJWT(realToken)

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

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.AbortError(c, http.StatusUnauthorized, "无效claims格式")
			return
		}

		// 安全获取user_id
		userID, exists := claims["user_id"]
		if !exists {
			utils.AbortError(c, http.StatusUnauthorized, "token缺少user_id字段")
			return
		}
		user_id := utils.InterfaceToInt64(userID)
		// 设置用户ID
		c.Set("userId", user_id)
		log.Printf("userId=%d\n", user_id)

		rc := ctr.GetRedis()
		key := fmt.Sprintf("%s:%d", utils.JwtCurrentUser, user_id)
		userJson, err := rc.Get(c, key).Result()
		if err != nil || userJson == "" {
			us := ctr.MustGet(container.AuthServiceName).(service.AuthService)

			user, err := us.GetCurrentUser(user_id)
			if err != nil {
				utils.AbortError(c, http.StatusUnauthorized, "1、用户不存在")
				return
			}
			currentUser := models.CurrentUser{
				Id:      user.Id,
				Name:    user.Name,
				Account: user.Account,
			}
			// 写redis
			userJson, err := json.Marshal(&currentUser)
			if err != nil {
				utils.AbortError(c, http.StatusInternalServerError, "服务异常")
				return
			}
			// 写入Redis
			if err := rc.Set(c, key, userJson, 0).Err(); err != nil {
				utils.AbortError(c, http.StatusInternalServerError, "服务异常")
				return
			}
			// 将用户信息存入上下文
			c.Set(utils.JwtCurrentUser, &currentUser)
		} else {
			// 反序列化为CurrentUser结构体
			var currentUser models.CurrentUser
			if err := json.Unmarshal([]byte(userJson), &currentUser); err != nil {
				utils.AbortError(c, http.StatusUnauthorized, "2、用户不存在")
				return
			}
			// 将用户信息存入上下文
			c.Set(utils.JwtCurrentUser, &currentUser)
		}
		c.Next()
	}
}
