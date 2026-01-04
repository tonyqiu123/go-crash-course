package user_auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler holds dependencies for user_auth handlers
type Handler struct {
	DB *gorm.DB
}

// NewHandler creates a new user_auth handler
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// GetCurrentUser handles GET /api/auth/me - get current user info
// Requires: JWT authentication
func (h *Handler) GetCurrentUser(c *gin.Context) {
	// TODO: Implement current user retrieval
	// 1. Get user ID from JWT token (set by auth middleware)
	// 2. Query user record by Clerk ID
	// 3. Return user data

	c.JSON(http.StatusOK, gin.H{
		"id":       "",
		"email":    "",
		"name":     "",
		"role":     "user",
		"clerk_id": "",
	})
}

// UpdateUserProfile handles PUT /api/auth/profile - update user profile
// Requires: JWT authentication
func (h *Handler) UpdateUserProfile(c *gin.Context) {
	// TODO: Implement user profile update
	// 1. Get user ID from JWT token
	// 2. Parse update data from request body
	// 3. Update user record
	// 4. Return updated user data

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
	})
}

// SyncClerkUser handles POST /api/auth/sync - sync user from Clerk webhook
// This endpoint is called by Clerk webhooks to sync user data
func (h *Handler) SyncClerkUser(c *gin.Context) {
	// TODO: Implement Clerk user sync
	// 1. Verify Clerk webhook signature
	// 2. Parse webhook event (user.created, user.updated, user.deleted)
	// 3. Create/update/delete user record in database
	// 4. Return success response

	c.JSON(http.StatusOK, gin.H{
		"received": true,
	})
}
