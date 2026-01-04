package clubs

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers club-related routes
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	clubs := rg.Group("/clubs")
	{
		clubs.GET("/", handler.GetClubs)
		// TODO: Add rate limiting middleware
		// TODO: Add authentication middleware where needed
	}
}
