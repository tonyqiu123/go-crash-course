package payments

import (
	"time"

	"gorm.io/gorm"
)

// Payment represents a payment transaction
type Payment struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	UserID          string         `gorm:"size:255;index;not null" json:"user_id"`
	Amount          float64        `gorm:"not null" json:"amount"`
	Currency        string         `gorm:"size:3;default:'CAD'" json:"currency"`
	Status          string         `gorm:"size:32;not null" json:"status"` // pending, completed, failed, refunded
	PaymentMethod   string         `gorm:"size:64" json:"payment_method"`
	TransactionID   *string        `gorm:"size:255;uniqueIndex" json:"transaction_id"`
	StripeSessionID *string        `gorm:"size:255" json:"stripe_session_id"`
	Metadata        map[string]any `gorm:"type:jsonb" json:"metadata"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for GORM
func (Payment) TableName() string {
	return "payments"
}
