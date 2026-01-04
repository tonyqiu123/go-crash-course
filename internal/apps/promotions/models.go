package promotions

import (
	"time"

	"gorm.io/gorm"
)

// Promotion represents a promotional campaign or featured content
type Promotion struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"size:255;not null" json:"title"`
	Description *string        `gorm:"type:text" json:"description"`
	ImageURL    *string        `gorm:"type:text" json:"image_url"`
	LinkURL     *string        `gorm:"type:text" json:"link_url"`
	Active      bool           `gorm:"default:true" json:"active"`
	StartDate   *time.Time     `json:"start_date"`
	EndDate     *time.Time     `json:"end_date"`
	Priority    int            `gorm:"default:0" json:"priority"` // Higher priority shown first
	Metadata    map[string]any `gorm:"type:jsonb" json:"metadata"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for GORM
func (Promotion) TableName() string {
	return "promotions"
}
