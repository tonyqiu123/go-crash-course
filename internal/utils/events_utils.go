package utils

// DetermineDisplayHandle returns the most relevant social media handle for an event
// Priority: IG > Discord > X > TikTok > FB > Other
func DetermineDisplayHandle(igHandle, discordHandle, xHandle, tiktokHandle, fbHandle, otherHandle *string) string {
	// TODO: Implement handle priority logic
	// Return first non-empty handle in priority order

	if igHandle != nil && *igHandle != "" {
		return "@" + *igHandle
	}
	if discordHandle != nil && *discordHandle != "" {
		return *discordHandle
	}
	if xHandle != nil && *xHandle != "" {
		return "@" + *xHandle
	}
	if tiktokHandle != nil && *tiktokHandle != "" {
		return "@" + *tiktokHandle
	}
	if fbHandle != nil && *fbHandle != "" {
		return *fbHandle
	}
	if otherHandle != nil && *otherHandle != "" {
		return *otherHandle
	}

	return ""
}

// BuildEventURL constructs a full event URL
func BuildEventURL(eventID uint, baseURL string) string {
	// TODO: Implement URL construction
	// Build full URL for event detail page

	return ""
}

// TruncateText truncates text to a maximum length with ellipsis
func TruncateText(text string, maxLength int) string {
	// TODO: Implement text truncation
	// Truncate text and add "..." if longer than maxLength

	if len(text) <= maxLength {
		return text
	}

	return text[:maxLength-3] + "..."
}
