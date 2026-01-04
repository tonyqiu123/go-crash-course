package waitlist

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler holds dependencies for waitlist handlers
type Handler struct {
	DB *gorm.DB
}

// NewHandler creates a new waitlist handler
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// Join handles POST /api/waitlist/join - join the waitlist
// Body: { "email": "user@example.com", "name": "John Doe", "school": "University" }
func (h *Handler) Join(c *gin.Context) {
	// TODO: Implement waitlist join
	// 1. Parse request body (email, name, school, metadata)
	// 2. Validate email format
	// 3. Check if email already exists
	// 4. Create new waitlist entry
	// 5. Send confirmation email
	// 6. Return success response

	c.JSON(http.StatusCreated, gin.H{
		"message": "Added to waitlist successfully",
	})
}

// GetStats handles GET /api/waitlist/stats - get waitlist statistics
// Requires: Admin authentication
func (h *Handler) GetStats(c *gin.Context) {
	// TODO: Implement waitlist stats retrieval
	// 1. Validate admin authentication
	// 2. Count total entries
	// 3. Group by school or other metadata
	// 4. Return statistics

	c.JSON(http.StatusOK, gin.H{
		"total": 0,
		"stats": map[string]int{},
	})
}
