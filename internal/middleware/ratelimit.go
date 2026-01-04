package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter provides rate limiting functionality
type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.Mutex
	rate     int           // requests per window
	window   time.Duration // time window
}

// NewRateLimiter creates a new rate limiter
// Parameters:
//   - rate: maximum number of requests allowed
//   - window: time window for rate limiting (e.g., 1 hour)
func NewRateLimiter(rate int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		rate:     rate,
		window:   window,
	}
}

// Limit returns a middleware that enforces rate limiting
func (rl *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement rate limiting logic
		// 1. Get client identifier (IP address)
		// 2. Check request history for this client
		// 3. Count requests in current window
		// 4. If over limit, return 429 Too Many Requests
		// 5. Otherwise, record request and continue

		clientIP := c.ClientIP()

		rl.mu.Lock()
		defer rl.mu.Unlock()

		now := time.Now()
		cutoff := now.Add(-rl.window)

		// Clean old requests
		if requests, exists := rl.requests[clientIP]; exists {
			filtered := []time.Time{}
			for _, reqTime := range requests {
				if reqTime.After(cutoff) {
					filtered = append(filtered, reqTime)
				}
			}
			rl.requests[clientIP] = filtered
		}

		// Check rate limit
		if len(rl.requests[clientIP]) >= rl.rate {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		// Record request
		rl.requests[clientIP] = append(rl.requests[clientIP], now)

		c.Next()
	}
}

// RateLimit creates a simple rate limiting middleware
// Parameters:
//   - requestsPerHour: maximum requests per hour
func RateLimit(requestsPerHour int) gin.HandlerFunc {
	limiter := NewRateLimiter(requestsPerHour, time.Hour)
	return limiter.Limit()
}
