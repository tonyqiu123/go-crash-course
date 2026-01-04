package waitlist

import (
	"time"

	"gorm.io/gorm"
)

// WaitlistEntry represents a user on the waitlist
type WaitlistEntry struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Name      *string        `gorm:"size:255" json:"name"`
	School    *string        `gorm:"size:255" json:"school"`
	Metadata  map[string]any `gorm:"type:jsonb" json:"metadata"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for GORM
func (WaitlistEntry) TableName() string {
	return "waitlist_entries"
}
