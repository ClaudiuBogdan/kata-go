package anagramchecker

import "testing"

func TestAreAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "Simple anagrams",
			s1:       "listen",
			s2:       "silent",
			expected: true,
		},
		{
			name:     "Not anagrams",
			s1:       "hello",
			s2:       "world",
			expected: false,
		},
		{
			name:     "Anagrams with spaces",
			s1:       "debit card",
			s2:       "bad credit",
			expected: true,
		},
		{
			name:     "Anagrams with different cases",
			s1:       "Dormitory",
			s2:       "Dirty Room",
			expected: true,
		},
		{
			name:     "Anagrams with punctuation",
			s1:       "A decimal point",
			s2:       "I'm a dot in place.",
			expected: true,
		},
		{
			name:     "Not anagrams with same length",
			s1:       "software",
			s2:       "hardware",
			expected: false,
		},
		{
			name:     "Empty strings",
			s1:       "",
			s2:       "",
			expected: true,
		},
		{
			name:     "Single character strings",
			s1:       "a",
			s2:       "A",
			expected: true,
		},
		{
			name:     "Strings with numbers",
			s1:       "12345",
			s2:       "54321",
			expected: true,
		},
		{
			name:     "Strings with repeating characters",
			s1:       "aabbcc",
			s2:       "abcabc",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AreAnagrams(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("AreAnagrams(%q, %q) = %v, want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}
