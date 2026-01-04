package core

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers core routes (health, home)
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	handler := NewHandler(db)

	// Root level routes
	router.GET("/", handler.Home)
	router.GET("/health", handler.Health)
}
