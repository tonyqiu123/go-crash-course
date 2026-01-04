package events

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers event-related routes
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	events := rg.Group("/events")
	{
		events.GET("/latest-update", handler.GetLatestUpdate)
		events.GET("/", handler.GetEvents)
		events.GET("/:id", handler.GetEvent)
		events.GET("/export/ics", handler.ExportEventsICS)
		events.GET("/google-calendar-urls", handler.GetGoogleCalendarURLs)
		
		// Protected routes (require JWT)
		events.POST("/extract", handler.ExtractEventFromScreenshot)
		events.POST("/submit", handler.SubmitEvent)
		
		// TODO: Add rate limiting middleware
		// TODO: Add authentication middleware for protected routes
	}

	// RSS feed at root level
	// TODO: Register at router level, not under /api/events
	// router.GET("/rss.xml", handler.RSSFeed)
}
