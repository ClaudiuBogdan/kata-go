package palindromechecker

import (
	"unicode"
)

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	input := []rune(cleanString(s))
	n := len(input)
	for i := 0; i < n/2; i++ {
		if input[i] != input[n-i-1] {
			return false
		}
	}

	return true
}

// cleanString converts a string to lowercase and removes non-alphanumeric characters
func cleanString(s string) string {
	output := []rune(nil)

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			output = append(output, unicode.ToLower(r))
		}
	}

	return string(output)
}
