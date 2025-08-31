package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ix-pay/ixpay/container"
	"github.com/ix-pay/ixpay/controllers"
	"github.com/ix-pay/ixpay/middleware"
)

func SetupAuthRoutes(r *gin.RouterGroup, ctr container.IContainer) {
	con := controllers.NewAuthController(ctr)

	auth := r.Group("/auth")
	{
		auth.POST("/login", con.Login)
		auth.POST("/register", con.Register)
		auth.GET("/profile", middleware.JWTAuth(ctr), con.GetProfile)
	}
}
