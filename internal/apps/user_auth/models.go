package user_auth

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user account
// Note: This app uses Clerk for authentication, so this model may be minimal
// or used for storing additional user metadata not in Clerk
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ClerkID   string         `gorm:"size:255;uniqueIndex;not null" json:"clerk_id"` // Clerk user ID
	Email     string         `gorm:"size:255;index" json:"email"`
	Name      *string        `gorm:"size:255" json:"name"`
	Role      string         `gorm:"size:32;default:'user'" json:"role"` // user, admin
	Metadata  map[string]any `gorm:"type:jsonb" json:"metadata"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "users"
}
