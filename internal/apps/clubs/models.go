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
