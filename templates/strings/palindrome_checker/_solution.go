package palindromechecker

import (
	"strings"
	"unicode"
)

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	// Convert to lowercase and remove non-alphanumeric characters
	s = cleanString(s)

	// Use two pointers to check if the string is a palindrome
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// cleanString converts a string to lowercase and removes non-alphanumeric characters
func cleanString(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, s)
}
