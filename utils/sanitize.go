package utils

import (
	"html"
	"regexp"
	"strings"
)

// SanitizeString trims whitespace and removes dangerous characters
func SanitizeString(input string) string {
	// Trim whitespace
	sanitized := strings.TrimSpace(input)

	// Escape HTML entities (prevents script injection)
	sanitized = html.EscapeString(sanitized)

	// Remove unwanted special characters but allow those commonly used in URLs
	re := regexp.MustCompile(`[^\w\s\-._~:/?#[\]@!$&'()*+,;=%]`)
	sanitized = re.ReplaceAllString(sanitized, "")

	return sanitized
}
