# Building a Go Web Backend from Scratch

> **Course Guide**: A step-by-step walkthrough for creating a professional, production-ready Go web server.

This guide walks you through setting up a Go web backend like the one in this repository. By the end, you'll have a well-organized project with:

- ✅ A proper Go project structure
- ✅ HTTP routing with the Gin framework
- ✅ Database integration with GORM (PostgreSQL)
- ✅ Environment configuration
- ✅ Middleware patterns (CORS, rate limiting)
- ✅ Feature-based app organization
- ✅ Database migrations

---

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Project Setup](#step-1-project-setup)
3. [Creating the Entry Point](#step-2-creating-the-entry-point)
4. [Configuration Management](#step-3-configuration-management)
5. [Database Connection](#step-4-database-connection)
6. [Route Registration](#step-5-route-registration)
7. [Creating Your First App (Feature Module)](#step-6-creating-your-first-app)
8. [Adding Middleware](#step-7-adding-middleware)
9. [Creating Services](#step-8-creating-services)
10. [Utility Functions](#step-9-utility-functions)
11. [Database Migrations](#step-10-database-migrations)
12. [Running and Testing](#step-11-running-and-testing)
13. [Project Structure Summary](#project-structure-summary)

---

## Prerequisites

Before starting, ensure you have:

- **Go 1.21+** installed ([Download Go](https://go.dev/dl/))
- **PostgreSQL** running locally or accessible remotely
- A code editor (VS Code with Go extension recommended)
- Basic familiarity with Go syntax

---

## Step 1: Project Setup

### 1.1 Create the Project Directory

```bash
mkdir backend-go
cd backend-go
```

### 1.2 Initialize the Go Module

```bash
go mod init github.com/yourusername/yourproject/backend-go
```

> **Tip**: Use a full module path that matches where your code will be hosted. This is a Go convention.

### 1.3 Create the Directory Structure

```bash
# Create the main directories
mkdir -p cmd/server
mkdir -p internal/{apps,config,middleware,services,utils}
mkdir -p migrations
```

Your folder structure should look like this:

```
backend-go/
├── cmd/
│   └── server/          # Application entry point
├── internal/            # Private application code
│   ├── apps/            # Feature modules (clubs, events, etc.)
│   ├── config/          # Configuration and setup
│   ├── middleware/      # HTTP middleware
│   ├── services/        # Shared services (email, storage, etc.)
│   └── utils/           # Utility/helper functions
├── migrations/          # SQL migration files
├── go.mod               # Go module file
└── .env.example         # Environment variables template
```

> **Why `internal/`?** The `internal` directory is special in Go. Packages inside it cannot be imported by code outside your module. This protects your private implementation details.

### 1.4 Install Dependencies

```bash
go get github.com/gin-gonic/gin           # Web framework
go get gorm.io/gorm                        # ORM
go get gorm.io/driver/postgres             # PostgreSQL driver for GORM
go get github.com/joho/godotenv            # .env file loading
go get github.com/lib/pq                   # PostgreSQL driver
go get github.com/go-redis/redis/v8        # Redis client (optional)
go get github.com/golang-jwt/jwt/v5        # JWT handling (optional)
```

---

## Step 2: Creating the Entry Point

The entry point is where your application starts. Create `cmd/server/main.go`:

```go
package main

import (
	"log"
	"os"

	"github.com/yourusername/yourproject/backend-go/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from environment variables
	cfg := config.LoadConfig()

	// Initialize database connection
	db := config.InitDatabase(cfg)

	// Create Gin router with default middleware (logger, recovery)
	router := gin.Default()

	// Setup custom middleware (CORS, rate limiting, etc.)
	// middleware.Setup(router, cfg)

	// Register all routes
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
```

### Key Concepts

| Concept | Description |
|---------|-------------|
| `package main` | Required for executable programs |
| `func main()` | The entry point of your application |
| `gin.Default()` | Creates a Gin router with logging and recovery middleware |
| `router.Run()` | Starts the HTTP server |

---

## Step 3: Configuration Management

Centralized configuration makes your app easier to deploy and test.

### 3.1 Create the Config Struct

Create `internal/config/config.go`:

```go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	DatabaseURL    string
	Port           string
	Environment    string
	OpenAIAPIKey   string
	AWSRegion      string
	AWSBucket      string
	JWTSecret      string
	RedisURL       string
	AllowedOrigins []string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if it exists (development only)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config := &Config{
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://localhost:5432/mydb?sslmode=disable"),
		Port:           getEnv("PORT", "8000"),
		Environment:    getEnv("ENVIRONMENT", "development"),
		OpenAIAPIKey:   getEnv("OPENAI_API_KEY", ""),
		AWSRegion:      getEnv("AWS_REGION", "us-east-1"),
		AWSBucket:      getEnv("AWS_BUCKET", ""),
		JWTSecret:      getEnv("JWT_SECRET", ""),
		RedisURL:       getEnv("REDIS_URL", "localhost:6379"),
		AllowedOrigins: []string{getEnv("ALLOWED_ORIGINS", "*")},
	}

	return config
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
```

### 3.2 Create the Environment Template

Create `.env.example` in your project root:

```bash
# Server Configuration
PORT=8000
ENVIRONMENT=development

# Database
DATABASE_URL=postgres://user:password@localhost:5432/mydb?sslmode=disable

# External APIs
OPENAI_API_KEY=your_api_key_here

# AWS (Optional)
AWS_REGION=us-east-1
AWS_BUCKET=your_bucket_name

# Authentication
JWT_SECRET=your_jwt_secret_here

# Redis (Optional)
REDIS_URL=localhost:6379

# CORS
ALLOWED_ORIGINS=http://localhost:3000
```

> **Important**: Copy `.env.example` to `.env` and fill in your actual values. Never commit `.env` to version control!

---

## Step 4: Database Connection

### Create the Database Module

Create `internal/config/database.go`:

```go
package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDatabase initializes the database connection
func InitDatabase(cfg *Config) *gorm.DB {
	// Configure GORM logger based on environment
	var logLevel logger.LogLevel
	if cfg.Environment == "production" {
		logLevel = logger.Error
	} else {
		logLevel = logger.Info
	}

	// Open database connection
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogLevel(logLevel),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate models (optional - use migrations in production)
	// err = db.AutoMigrate(&models.User{}, &models.Club{})
	// if err != nil {
	//     log.Fatalf("Failed to auto-migrate: %v", err)
	// }

	log.Println("Database connection established")
	return db
}
```

### Key GORM Concepts

| Concept | Description |
|---------|-------------|
| `gorm.Open()` | Opens a database connection |
| `postgres.Open()` | Creates a PostgreSQL-specific driver |
| `AutoMigrate()` | Automatically creates/updates tables based on model structs |
| `logger.LogLevel` | Controls SQL query logging |

---

## Step 5: Route Registration

Centralize route registration for cleaner organization.

Create `internal/config/routes.go`:

```go
package config

import (
	"github.com/yourusername/yourproject/backend-go/internal/apps/clubs"
	"github.com/yourusername/yourproject/backend-go/internal/apps/events"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Health check route (no /api prefix)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API routes group
	api := router.Group("/api")
	{
		// Register each app's routes
		clubs.RegisterRoutes(api, db)
		events.RegisterRoutes(api, db)
		// Add more apps here...
	}
}
```

---

## Step 6: Creating Your First App

Each "app" is a self-contained feature module with its own models, handlers, and routes.

### 6.1 Directory Structure

```
internal/apps/clubs/
├── handlers.go    # HTTP request handlers
├── models.go      # Database models
└── routes.go      # Route definitions
```

### 6.2 Create the Model

Create `internal/apps/clubs/models.go`:

```go
package clubs

import (
	"time"

	"gorm.io/gorm"
)

// Clubs represents a club in the database
type Clubs struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	ClubName   string         `gorm:"size:100;uniqueIndex;not null" json:"club_name"`
	Categories []string       `gorm:"type:jsonb;default:'[]'" json:"categories"`
	ClubPage   *string        `gorm:"type:text" json:"club_page"`
	IG         *string        `gorm:"type:text" json:"ig"`
	Discord    *string        `gorm:"type:text" json:"discord"`
	ClubType   *string        `gorm:"size:50" json:"club_type"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for GORM
func (Clubs) TableName() string {
	return "clubs"
}
```

### GORM Tag Reference

| Tag | Description |
|-----|-------------|
| `gorm:"primaryKey"` | Marks field as primary key |
| `gorm:"size:100"` | Sets VARCHAR length |
| `gorm:"uniqueIndex"` | Creates unique index |
| `gorm:"not null"` | Field cannot be null |
| `gorm:"type:jsonb"` | PostgreSQL JSONB type |
| `gorm:"default:'[]'"` | Default value |
| `json:"field_name"` | JSON serialization name |
| `json:"-"` | Exclude from JSON output |

### 6.3 Create the Handler

Create `internal/apps/clubs/handlers.go`:

```go
package clubs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler holds dependencies for club handlers
type Handler struct {
	DB *gorm.DB
}

// NewHandler creates a new clubs handler
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

// GetClubs handles GET /api/clubs/
// Query params: search, category, cursor, limit
func (h *Handler) GetClubs(c *gin.Context) {
	// Parse query parameters
	search := c.Query("search")
	limit := c.DefaultQuery("limit", "50")
	
	var clubs []Clubs
	query := h.DB.Model(&Clubs{})

	// Apply search filter if provided
	if search != "" {
		query = query.Where("club_name ILIKE ?", "%"+search+"%")
	}

	// Execute query
	if err := query.Limit(50).Find(&clubs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clubs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results":    clubs,
		"nextCursor": nil,
		"hasMore":    false,
		"totalCount": len(clubs),
	})
}

// CreateClub handles POST /api/clubs/
func (h *Handler) CreateClub(c *gin.Context) {
	var club Clubs
	
	// Bind JSON body to struct
	if err := c.ShouldBindJSON(&club); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create record in database
	if err := h.DB.Create(&club).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create club"})
		return
	}

	c.JSON(http.StatusCreated, club)
}

// GetClubByID handles GET /api/clubs/:id
func (h *Handler) GetClubByID(c *gin.Context) {
	id := c.Param("id")
	
	var club Clubs
	if err := h.DB.First(&club, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Club not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch club"})
		return
	}

	c.JSON(http.StatusOK, club)
}
```

### Gin Context Methods

| Method | Description |
|--------|-------------|
| `c.Query("key")` | Get query parameter |
| `c.DefaultQuery("key", "default")` | Get query param with default |
| `c.Param("id")` | Get URL path parameter |
| `c.ShouldBindJSON(&struct)` | Parse JSON body into struct |
| `c.JSON(status, data)` | Send JSON response |

### 6.4 Create the Routes

Create `internal/apps/clubs/routes.go`:

```go
package clubs

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers club-related routes
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	handler := NewHandler(db)

	clubs := rg.Group("/clubs")
	{
		clubs.GET("/", handler.GetClubs)
		clubs.POST("/", handler.CreateClub)
		clubs.GET("/:id", handler.GetClubByID)
	}
}
```

---

## Step 7: Adding Middleware

Middleware intercepts requests before they reach your handlers.

### 7.1 CORS Middleware

Create `internal/middleware/cors.go`:

```go
package middleware

import (
	"github.com/gin-gonic/gin"
)

// CORS returns a middleware that handles Cross-Origin Resource Sharing
func CORS(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if origin is allowed
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == "*" || allowedOrigin == origin {
				allowed = true
				break
			}
		}

		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "86400")
		}

		// Handle preflight request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
```

### 7.2 Rate Limiting Middleware

Create `internal/middleware/ratelimit.go`:

```go
package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter implements a simple rate limiter
type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.Mutex
	limit    int
	window   time.Duration
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

// Middleware returns a Gin middleware function
func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		rl.mu.Lock()
		defer rl.mu.Unlock()

		// Clean old requests
		cutoff := now.Add(-rl.window)
		var valid []time.Time
		for _, t := range rl.requests[ip] {
			if t.After(cutoff) {
				valid = append(valid, t)
			}
		}
		rl.requests[ip] = valid

		// Check if over limit
		if len(rl.requests[ip]) >= rl.limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		// Record this request
		rl.requests[ip] = append(rl.requests[ip], now)
		c.Next()
	}
}
```

### 7.3 Using Middleware

In `main.go`, add middleware before routes:

```go
// After creating the router
router := gin.Default()

// Apply global middleware
router.Use(middleware.CORS(cfg.AllowedOrigins))

// Apply rate limiting to specific routes
limiter := middleware.NewRateLimiter(100, time.Minute)
api := router.Group("/api")
api.Use(limiter.Middleware())
```

---

## Step 8: Creating Services

Services encapsulate external integrations and shared business logic.

### Example: Email Service

Create `internal/services/email_service.go`:

```go
package services

// EmailService provides email sending functionality
type EmailService struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	FromEmail    string
	FromName     string
}

// NewEmailService creates a new email service instance
func NewEmailService(host string, port int, username, password, fromEmail, fromName string) *EmailService {
	return &EmailService{
		SMTPHost:     host,
		SMTPPort:     port,
		SMTPUsername: username,
		SMTPPassword: password,
		FromEmail:    fromEmail,
		FromName:     fromName,
	}
}

// SendEmail sends a plain text email
func (s *EmailService) SendEmail(to, subject, body string) error {
	// Implementation: Connect to SMTP, authenticate, send email
	// Use net/smtp package or a library like gomail
	return nil
}

// SendHTMLEmail sends an HTML email
func (s *EmailService) SendHTMLEmail(to, subject, htmlBody string) error {
	// Similar to SendEmail but with HTML content type
	return nil
}
```

---

## Step 9: Utility Functions

Create reusable helper functions in `internal/utils/`.

### Example: Validation Utilities

Create `internal/utils/validation.go`:

```go
package utils

import (
	"errors"
	"regexp"
	"strings"
)

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// SanitizeString removes potentially harmful characters
func SanitizeString(input string) string {
	return strings.TrimSpace(input)
}

// ValidateURL validates URL format
func ValidateURL(url string) bool {
	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}

// ValidateRequired checks if required fields are present
func ValidateRequired(data map[string]interface{}, fields []string) error {
	for _, field := range fields {
		val, exists := data[field]
		if !exists {
			return errors.New(field + " is required")
		}
		if str, ok := val.(string); ok && strings.TrimSpace(str) == "" {
			return errors.New(field + " cannot be empty")
		}
	}
	return nil
}
```

---

## Step 10: Database Migrations

For production, use explicit SQL migrations instead of AutoMigrate.

### Migration File Naming Convention

```
migrations/
├── 000001_initial.up.sql     # Creates tables
├── 000001_initial.down.sql   # Drops tables (rollback)
├── 000002_add_users.up.sql
└── 000002_add_users.down.sql
```

### Example: Initial Migration

Create `migrations/000001_initial.up.sql`:

```sql
-- Create clubs table
CREATE TABLE IF NOT EXISTS clubs (
    id SERIAL PRIMARY KEY,
    club_name VARCHAR(100) NOT NULL UNIQUE,
    categories JSONB DEFAULT '[]',
    club_page TEXT,
    ig TEXT,
    discord TEXT,
    club_type VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_clubs_club_name ON clubs(club_name);
CREATE INDEX IF NOT EXISTS idx_clubs_deleted_at ON clubs(deleted_at);
```

Create `migrations/000001_initial.down.sql`:

```sql
DROP TABLE IF EXISTS clubs;
```

### Running Migrations

Install `golang-migrate`:

```bash
brew install golang-migrate  # macOS

# Run migrations
migrate -path ./migrations -database "postgres://user:pass@localhost:5432/mydb?sslmode=disable" up

# Rollback
migrate -path ./migrations -database "..." down 1
```

---

## Step 11: Running and Testing

### Run the Server

```bash
# Development
go run cmd/server/main.go

# Build and run
go build -o server cmd/server/main.go
./server
```

### Test Endpoints

```bash
# Health check
curl http://localhost:8000/health

# Get clubs
curl http://localhost:8000/api/clubs/

# Create a club
curl -X POST http://localhost:8000/api/clubs/ \
  -H "Content-Type: application/json" \
  -d '{"club_name": "Go Learners Club"}'

# Get specific club
curl http://localhost:8000/api/clubs/1
```

---

## Project Structure Summary

```
backend-go/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── apps/                     # Feature modules
│   │   ├── clubs/
│   │   │   ├── handlers.go      # HTTP handlers
│   │   │   ├── models.go        # Database models
│   │   │   └── routes.go        # Route definitions
│   │   ├── events/
│   │   └── ...
│   ├── config/
│   │   ├── config.go            # Configuration struct & loading
│   │   ├── database.go          # Database initialization
│   │   └── routes.go            # Central route registration
│   ├── middleware/
│   │   ├── cors.go              # CORS middleware
│   │   └── ratelimit.go         # Rate limiting
│   ├── services/
│   │   ├── email_service.go     # Email functionality
│   │   └── storage_service.go   # File storage
│   └── utils/
│       └── validation.go        # Helper functions
├── migrations/
│   ├── 000001_initial.up.sql
│   └── 000001_initial.down.sql
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

---

## Quick Reference Commands

| Command | Description |
|---------|-------------|
| `go mod init <path>` | Initialize Go module |
| `go mod tidy` | Clean up dependencies |
| `go get <package>` | Add a dependency |
| `go run ./...` | Run the application |
| `go build -o <name> ./...` | Build executable |
| `go test ./...` | Run all tests |
| `go fmt ./...` | Format all code |
| `go vet ./...` | Check for common mistakes |

---

## Next Steps

1. **Add Authentication** - Implement JWT or OAuth middleware
2. **Add Tests** - Write unit and integration tests
3. **Add Logging** - Use structured logging (zerolog, zap)
4. **Add API Documentation** - Use Swagger/OpenAPI
5. **Containerize** - Add Dockerfile for deployment

Happy coding! 🚀
