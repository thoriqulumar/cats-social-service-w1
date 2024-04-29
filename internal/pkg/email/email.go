package email

import (
	"regexp"
)

// IsValidEmail checks if the given email address has a valid format
func IsValidEmail(email string) bool {
	// Regular expression for email validation
	// This regex pattern is a simplified version and may not cover all edge cases
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Check if the email matches the pattern
	return regex.MatchString(email)
}
