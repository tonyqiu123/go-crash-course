package main

import (
	"log"
	"os"

	"github.com/ericahan22/bug-free-octo-spork/backend-go/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	db := config.InitDatabase(cfg)

	// Create Gin router
	router := gin.Default()

	// Setup middleware
	// TODO: Add CORS, rate limiting, authentication middleware

	// Register routes
	config.RegisterRoutes(router, db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
