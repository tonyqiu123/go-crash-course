# Follow SETUP_GUIDE.md for step by step instructions of building the backend service

# Backend-Go

Go implementation of the backend service, providing a RESTful API for the bug-free-octo-spork application.

## Project Structure

```
backend-go/
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── apps/            # Application modules
│   │   ├── clubs/
│   │   ├── events/
│   │   ├── newsletter/
│   │   ├── payments/
│   │   ├── promotions/
│   │   ├── realtime/
│   │   ├── user_auth/
│   │   ├── waitlist/
│   │   └── core/
│   ├── services/        # Shared services (OpenAI, Email, Storage)
│   ├── utils/           # Utility functions
│   └── middleware/      # HTTP middleware
├── migrations/          # Database migrations
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites
- Go 1.21 or higher
- PostgreSQL database
- Environment variables configured (see `.env.example`)

### Installation

```bash
# Install dependencies
go mod download

# Run migrations
# TODO: Add migration tool instructions

# Run the server
go run cmd/server/main.go
```

## Development

This is scaffolding code with function signatures and struct definitions. Implementation logic needs to be added for:
- Business logic in handlers
- Database operations using GORM
- External service integrations
- Validation and error handling

## Architecture

- **Gin** - HTTP web framework
- **GORM** - ORM for database operations
- **PostgreSQL** - Primary database
- **JWT** - Authentication
- **Redis** - Caching and rate limiting

## API Endpoints

See individual app route files for endpoint definitions:
- `/api/events/` - Events management
- `/api/clubs/` - Clubs directory
- `/api/newsletter/` - Newsletter subscriptions
- `/api/promotions/` - Promotional content
- `/api/waitlist/` - Waitlist management
- `/health/` - Health check endpoint
