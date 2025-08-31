package routes

import (
	"github.com/ix-pay/ixpay/container"
	"github.com/ix-pay/ixpay/controllers"
	"github.com/ix-pay/ixpay/middleware"

	"github.com/gin-gonic/gin"
)

func SetupPaymentsRoutes(r *gin.RouterGroup, ctr container.IContainer) {
	con := controllers.NewPaymentController(ctr)
	payments := r.Group("/payments")
	{
		payments.POST("/create", middleware.JWTAuth(ctr), con.CreatePayment)
		payments.GET("/:id", middleware.JWTAuth(ctr), con.GetPayment)
		payments.GET("/list", middleware.JWTAuth(ctr), con.ListPayments)
		payments.PUT("/:id/status", middleware.JWTAuth(ctr), con.UpdatePaymentStatus)
	}
}
