package routes

import (
	"github.com/ix-pay/ixpay/controllers"
	"github.com/ix-pay/ixpay/middleware"

	"github.com/gin-gonic/gin"
)

func SetupPaymentsRoutes(r *gin.RouterGroup) {
	payments := r.Group("/payments")
	{
		payments.POST("/create", middleware.JWTAuth(), controllers.CreatePayment)
		payments.GET("/:id", middleware.JWTAuth(), controllers.GetPayment)
		payments.GET("/list", middleware.JWTAuth(), controllers.ListPayments)
		payments.PUT("/:id/status", middleware.JWTAuth(), controllers.UpdatePaymentStatus)
	}
}
