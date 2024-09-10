package anagramchecker

import (
	"strings"
	"unicode"
)

// AreAnagrams checks if two strings are anagrams
func AreAnagrams(s1, s2 string) bool {
	// Convert strings to lowercase and remove non-alphanumeric characters
	s1 = cleanString(s1)
	s2 = cleanString(s2)

	// If the lengths are different, they can't be anagrams
	if len(s1) != len(s2) {
		return false
	}

	// Create a map to store character counts
	charCount := make(map[rune]int)

	// Count characters in s1
	for _, char := range s1 {
		charCount[char]++
	}

	// Decrement counts for characters in s2
	for _, char := range s2 {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}

	// If all counts are zero, the strings are anagrams
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
