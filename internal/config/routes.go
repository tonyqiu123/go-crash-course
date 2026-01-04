package config

import (
	"github.com/ericahan22/bug-free-octo-spork/backend-go/internal/apps/clubs"
	"github.com/ericahan22/bug-free-octo-spork/backend-go/internal/apps/core"
	"github.com/ericahan22/bug-free-octo-spork/backend-go/internal/apps/events"
	"github.com/ericahan22/bug-free-octo-spork/backend-go/internal/apps/newsletter"
	"github.com/ericahan22/bug-free-octo-spork/backend-go/internal/apps/promotions"
	"github.com/ericahan22/bug-free-octo-spork/backend-go/internal/apps/waitlist"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Core routes
	core.RegisterRoutes(router, db)

	// API routes
	api := router.Group("/api")
	{
		// Events routes
		events.RegisterRoutes(api, db)

		// Clubs routes
		clubs.RegisterRoutes(api, db)

		// Newsletter routes
		newsletter.RegisterRoutes(api, db)

		// Promotions routes
		promotions.RegisterRoutes(api, db)

		// Waitlist routes
		waitlist.RegisterRoutes(api, db)

		// TODO: Add other app routes (payments, realtime, user_auth)
	}
}
