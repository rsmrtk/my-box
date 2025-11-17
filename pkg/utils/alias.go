package utils

import "strings"

// ToSnakeCase converts a string to snake case format by converting to lowercase
// and replacing spaces with underscores
func ToSnakeCase(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)
	// Replace spaces with underscores
	s = strings.ReplaceAll(s, " ", "_")
	return s
}
