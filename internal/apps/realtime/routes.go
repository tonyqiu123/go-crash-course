package realtime

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers realtime-related routes
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	realtime := rg.Group("/realtime")
	{
		realtime.GET("/ws", handler.HandleWebSocket)
		realtime.POST("/broadcast", handler.BroadcastUpdate)
		// TODO: Add WebSocket upgrade middleware
		// TODO: Add admin authentication for broadcast endpoint
	}
}
