package newsletter

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers newsletter-related routes
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	newsletter := rg.Group("/newsletter")
	{
		newsletter.POST("/subscribe", handler.Subscribe)
		newsletter.POST("/unsubscribe", handler.Unsubscribe)
		// TODO: Add rate limiting middleware
	}
}
