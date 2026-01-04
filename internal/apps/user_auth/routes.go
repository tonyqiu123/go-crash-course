package user_auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers user_auth-related routes
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	auth := rg.Group("/auth")
	{
		auth.GET("/me", handler.GetCurrentUser)
		auth.PUT("/profile", handler.UpdateUserProfile)
		auth.POST("/sync", handler.SyncClerkUser)
		// TODO: Add JWT authentication middleware for protected routes
	}
}
