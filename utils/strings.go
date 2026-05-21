package utils

import "strings"

func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

func Truncate(s string, max int) string {
	if max <= 0 || len(s) <= max {
		return s
	}
	return s[:max] + "..."
}

func DefaultIfBlank(s, fallback string) string {
	if IsBlank(s) {
		return fallback
	}
	return s
}
