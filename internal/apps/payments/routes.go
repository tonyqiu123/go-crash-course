package payments

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers payment-related routes
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	payments := rg.Group("/payments")
	{
		payments.POST("/create-checkout-session", handler.CreateCheckoutSession)
		payments.POST("/webhook", handler.HandleWebhook)
		payments.GET("/:id/status", handler.GetPaymentStatus)
		// TODO: Add JWT authentication middleware for protected routes
	}
}
