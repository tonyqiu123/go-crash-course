package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler holds dependencies for core handlers
type Handler struct {
	DB *gorm.DB
}

// NewHandler creates a new core handler
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// Health handles GET /health - health check endpoint
func (h *Handler) Health(c *gin.Context) {
	// TODO: Implement health check
	// 1. Check database connection
	// 2. Check Redis connection
	// 3. Return health status

	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"time":   "2024-01-01T00:00:00Z",
	})
}

// Home handles GET / - root endpoint
func (h *Handler) Home(c *gin.Context) {
	// TODO: Implement home endpoint
	// Return API information and version

	c.JSON(http.StatusOK, gin.H{
		"message": "Wat2Do API",
		"version": "1.0.0",
	})
}
