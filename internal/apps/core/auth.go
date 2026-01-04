package core

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Claims represents the JWT claims structure
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// JWTRequired is a middleware that requires valid JWT authentication
func JWTRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement JWT validation
		// 1. Extract token from Authorization header
		// 2. Validate token signature with Clerk public key
		// 3. Parse claims
		// 4. Set user_id in context
		// 5. Call next handler or abort with 401

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, gin.H{"error": "Missing or invalid authorization header"})
			c.Abort()
			return
		}

		// token := strings.TrimPrefix(authHeader, "Bearer ")
		// TODO: Validate token and extract claims

		c.Next()
	}
}

// OptionalJWT is a middleware that optionally validates JWT if present
func OptionalJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement optional JWT validation
		// Similar to JWTRequired but doesn't abort if token is missing
		// Just sets user_id in context if valid token is present

		c.Next()
	}
}

// AdminRequired is a middleware that requires admin role
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement admin role check
		// 1. Check if user is authenticated (should be used after JWTRequired)
		// 2. Check if user role is "admin"
		// 3. Call next handler or abort with 403

		c.Next()
	}
}
