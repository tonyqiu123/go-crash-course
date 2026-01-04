package utils

import (
	"errors"
	"regexp"
	"strings"
)

// ValidateEventData validates event submission data
// Returns cleaned data or validation error
func ValidateEventData(data map[string]interface{}) (map[string]interface{}, error) {
	// TODO: Implement event data validation
	// 1. Check required fields (title, location, occurrences)
	// 2. Validate field types and formats
	// 3. Sanitize text fields
	// 4. Validate occurrences array
	// 5. Return cleaned data or error

	cleaned := make(map[string]interface{})

	// Validate title
	title, ok := data["title"].(string)
	if !ok || strings.TrimSpace(title) == "" {
		return nil, errors.New("title is required")
	}
	cleaned["title"] = strings.TrimSpace(title)

	// Validate location
	location, ok := data["location"].(string)
	if !ok || strings.TrimSpace(location) == "" {
		return nil, errors.New("location is required")
	}
	cleaned["location"] = strings.TrimSpace(location)

	// TODO: Add more validation for other fields

	return cleaned, nil
}

// ValidateEmail validates email format
func ValidateEmail(email string) bool {
	// TODO: Implement email validation
	// Use regex to validate email format

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// SanitizeString removes potentially harmful characters from input
func SanitizeString(input string) string {
	// TODO: Implement string sanitization
	// Remove or escape potentially harmful characters

	// Basic sanitization - trim whitespace
	return strings.TrimSpace(input)
}

// ValidateURL validates URL format
func ValidateURL(url string) bool {
	// TODO: Implement URL validation
	// Check if string is a valid HTTP/HTTPS URL

	return strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
}
