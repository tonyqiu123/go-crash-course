package utils

import (
	"time"
)

// ParseUTCDateTime parses a datetime string into UTC time
// Supports various formats including ISO 8601
func ParseUTCDateTime(dateStr string) (time.Time, error) {
	// TODO: Implement datetime parsing
	// 1. Try multiple datetime formats
	// 2. Parse string into time.Time
	// 3. Convert to UTC
	// 4. Return parsed time or error

	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, dateStr); err == nil {
			return t.UTC(), nil
		}
	}

	return time.Time{}, nil
}

// FormatUTCDateTime formats a time.Time into ISO 8601 string
func FormatUTCDateTime(t time.Time) string {
	// TODO: Implement datetime formatting
	// Convert time to UTC and format as ISO 8601

	return t.UTC().Format(time.RFC3339)
}

// IsLiveEvent checks if an event is currently live
// Parameters:
//   - startTime: Event start time
//   - endTime: Event end time (can be nil)
//
// Returns: true if event is currently happening
func IsLiveEvent(startTime time.Time, endTime *time.Time) bool {
	// TODO: Implement live event check
	// 1. Get current time
	// 2. Check if now is between start and end
	// 3. If no end time, consider live if started within last 90 minutes

	now := time.Now()
	ninetyMinutesAgo := now.Add(-90 * time.Minute)

	if endTime != nil {
		// Event has end time - check if currently between start and end
		return now.After(startTime) && now.Before(*endTime)
	}

	// No end time - consider live if started within last 90 minutes
	return startTime.After(ninetyMinutesAgo) && startTime.Before(now)
}

// GetUpcomingEventDates filters and returns upcoming event dates
// Parameters:
//   - dates: Array of event dates
//
// Returns: Array of upcoming event dates (future or currently live)
func GetUpcomingEventDates(dates []time.Time) []time.Time {
	// TODO: Implement upcoming date filtering
	// Filter dates to only future or live events

	now := time.Now()
	upcoming := []time.Time{}

	for _, d := range dates {
		if d.After(now) {
			upcoming = append(upcoming, d)
		}
	}

	return upcoming
}
