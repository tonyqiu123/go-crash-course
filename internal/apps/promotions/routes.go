package promotions

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers promotion-related routes
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	promotions := rg.Group("/promotions")
	{
		promotions.GET("/", handler.GetPromotions)
		promotions.POST("/", handler.CreatePromotion)
		promotions.PUT("/:id", handler.UpdatePromotion)
		promotions.DELETE("/:id", handler.DeletePromotion)
		// TODO: Add admin authentication middleware for protected routes
	}
}
