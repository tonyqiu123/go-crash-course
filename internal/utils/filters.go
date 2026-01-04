package utils

import (
	"gorm.io/gorm"
)

// EventFilter provides filtering capabilities for event queries
type EventFilter struct {
	Search       string
	Categories   []string
	ClubType     string
	HasFood      *bool
	IsFree       *bool
	Registration *bool
	School       string
}

// ApplyEventFilters applies filters to a GORM query
func (f *EventFilter) ApplyEventFilters(query *gorm.DB) *gorm.DB {
	// TODO: Implement filter application
	// 1. Apply search filter if present
	// 2. Apply category filter
	// 3. Apply club_type filter
	// 4. Apply food filter
	// 5. Apply price filter
	// 6. Apply registration filter
	// 7. Apply school filter
	// Return modified query

	if f.Search != "" {
		// TODO: Add search across multiple fields
		query = query.Where("title ILIKE ? OR description ILIKE ? OR location ILIKE ?",
			"%"+f.Search+"%", "%"+f.Search+"%", "%"+f.Search+"%")
	}

	if len(f.Categories) > 0 {
		// TODO: Filter by categories (JSONB array contains)
	}

	if f.ClubType != "" {
		query = query.Where("club_type = ?", f.ClubType)
	}

	if f.HasFood != nil && *f.HasFood {
		query = query.Where("food IS NOT NULL AND food != ''")
	}

	if f.IsFree != nil && *f.IsFree {
		query = query.Where("price IS NULL OR price = 0")
	}

	if f.Registration != nil {
		query = query.Where("registration = ?", *f.Registration)
	}

	if f.School != "" {
		query = query.Where("school = ?", f.School)
	}

	return query
}

// ClubFilter provides filtering capabilities for club queries
type ClubFilter struct {
	Search   string
	Category string
	ClubType string
}

// ApplyClubFilters applies filters to a club query
func (f *ClubFilter) ApplyClubFilters(query *gorm.DB) *gorm.DB {
	// TODO: Implement club filter application

	if f.Search != "" {
		query = query.Where("club_name ILIKE ?", "%"+f.Search+"%")
	}

	if f.Category != "" && f.Category != "all" {
		// TODO: Filter by category (JSONB array contains)
	}

	if f.ClubType != "" {
		query = query.Where("club_type = ?", f.ClubType)
	}

	return query
}
