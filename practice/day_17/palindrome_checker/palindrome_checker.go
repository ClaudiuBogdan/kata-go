package palindromechecker

import (
	"unicode"
)

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	alphanumericRunes := cleanString(s)
	return isPalindromeRunes(alphanumericRunes)
}

// cleanString create an array of runes with only the alphanumeric runes in lower case
func cleanString(s string) []rune{
	var result []rune

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r){
			result = append(result, unicode.ToLower(r))
		}
	}
	return result
}

// isPalindromeRunes check if an array of runes is palindrome
func isPalindromeRunes(runes []rune) bool {
	n := len(runes)
	for i := 0; i < n/2; i++{
		if runes[i] != runes[n - i - 1] {
			return false
		}
	}
	return true
}