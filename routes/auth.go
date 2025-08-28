package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/controllers"
	"github.com/ix-pay/ixpay/middleware"
)

func SetupAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
		auth.GET("/profile/:userId", middleware.JWTAuth(), controllers.GetProfile)
	}
}
