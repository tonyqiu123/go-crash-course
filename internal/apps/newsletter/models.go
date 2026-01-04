package newsletter

import (
	"time"

	"gorm.io/gorm"
)

// NewsletterSubscriber represents a newsletter subscriber
type NewsletterSubscriber struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Active    bool           `gorm:"default:true" json:"active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for GORM
func (NewsletterSubscriber) TableName() string {
	return "newsletter_subscribers"
}
