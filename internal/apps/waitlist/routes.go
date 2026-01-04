package waitlist

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers waitlist-related routes
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	waitlist := rg.Group("/waitlist")
	{
		waitlist.POST("/join", handler.Join)
		waitlist.GET("/stats", handler.GetStats)
		// TODO: Add rate limiting middleware
		// TODO: Add admin authentication for stats endpoint
	}
}
