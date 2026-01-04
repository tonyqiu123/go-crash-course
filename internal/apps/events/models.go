package events

import (
	"time"

	"gorm.io/gorm"
)

// Events represents an event in the database
type Events struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           *string        `gorm:"type:text" json:"title"`
	Description     *string        `gorm:"type:text" json:"description"`
	Location        *string        `gorm:"type:text" json:"location"`
	Categories      []string       `gorm:"type:jsonb;default:'[]'" json:"categories"`
	Status          *string        `gorm:"size:32" json:"status"`
	SourceURL       *string        `gorm:"type:text" json:"source_url"`
	SourceImageURL  *string        `gorm:"type:text" json:"source_image_url"`
	Reactions       map[string]int `gorm:"type:jsonb;default:'{}'" json:"reactions"`
	PostedAt        *time.Time     `json:"posted_at"`
	CommentsCount   int            `gorm:"default:0" json:"comments_count"`
	LikesCount      int            `gorm:"default:0" json:"likes_count"`
	Food            *string        `gorm:"size:255" json:"food"`
	Registration    bool           `gorm:"default:false" json:"registration"`
	AddedAt         *time.Time     `json:"added_at"`
	Price           *float64       `json:"price"`
	School          *string        `gorm:"size:255" json:"school"`
	ClubType        *string        `gorm:"size:50" json:"club_type"`
	IGHandle        *string        `gorm:"size:100;column:ig_handle" json:"ig_handle"`
	DiscordHandle   *string        `gorm:"size:100" json:"discord_handle"`
	XHandle         *string        `gorm:"size:100;column:x_handle" json:"x_handle"`
	TiktokHandle    *string        `gorm:"size:100" json:"tiktok_handle"`
	FBHandle        *string        `gorm:"size:100;column:fb_handle" json:"fb_handle"`
	OtherHandle     *string        `gorm:"size:100" json:"other_handle"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	EventDates []EventDates `gorm:"foreignKey:EventID" json:"event_dates,omitempty"`
}

// TableName specifies the table name for GORM
func (Events) TableName() string {
	return "events"
}

// EventDates represents individual occurrence dates for events
type EventDates struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	EventID     uint           `gorm:"index:idx_event_dtstart;not null" json:"event_id"`
	DtstartUTC  time.Time      `gorm:"index:idx_dtstart_utc;index:idx_event_dtstart;not null" json:"dtstart_utc"`
	DtendUTC    *time.Time     `gorm:"index:idx_dtend_utc" json:"dtend_utc"`
	Duration    *time.Duration `json:"duration"`
	TZ          *string        `gorm:"size:64" json:"tz"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	Event Events `gorm:"foreignKey:EventID" json:"-"`
}

// TableName specifies the table name for GORM
func (EventDates) TableName() string {
	return "event_dates"
}

// EventSubmission represents user-submitted events pending admin review
type EventSubmission struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	SubmittedBy    string         `gorm:"size:255;not null" json:"submitted_by"`
	SubmittedAt    time.Time      `gorm:"index;not null" json:"submitted_at"`
	ReviewedAt     *time.Time     `json:"reviewed_at"`
	ReviewedBy     *string        `gorm:"size:255" json:"reviewed_by"`
	CreatedEventID uint           `gorm:"not null" json:"created_event_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	CreatedEvent Events `gorm:"foreignKey:CreatedEventID" json:"created_event,omitempty"`
}

// TableName specifies the table name for GORM
func (EventSubmission) TableName() string {
	return "event_submissions"
}

// EventInterest tracks user interest in events
type EventInterest struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	EventID   uint           `gorm:"index;not null;uniqueIndex:idx_event_user" json:"event_id"`
	UserID    string         `gorm:"size:255;index;not null;uniqueIndex:idx_event_user" json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	Event Events `gorm:"foreignKey:EventID" json:"-"`
}

// TableName specifies the table name for GORM
func (EventInterest) TableName() string {
	return "event_interests"
}

// IgnoredPost represents Instagram posts that should be ignored
type IgnoredPost struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Shortcode string         `gorm:"size:32;uniqueIndex;not null" json:"shortcode"`
	AddedAt   time.Time      `gorm:"not null" json:"added_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for GORM
func (IgnoredPost) TableName() string {
	return "ignored_posts"
}
