package lcs

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected int
	}{
		{
			name:     "Basic test",
			s1:       "ABCDGH",
			s2:       "AEDFHR",
			expected: 3,
		},
		{
			name:     "No common subsequence",
			s1:       "ABC",
			s2:       "DEF",
			expected: 0,
		},
		{
			name:     "One string is empty",
			s1:       "ABCDGH",
			s2:       "",
			expected: 0,
		},
		{
			name:     "Identical strings",
			s1:       "ABCDGH",
			s2:       "ABCDGH",
			expected: 6,
		},
		{
			name:     "Longer strings",
			s1:       "AGGTAB",
			s2:       "GXTXAYB",
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestCommonSubsequence(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("LongestCommonSubsequence() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLongestCommonSubsequenceString(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected string
	}{
		{
			name:     "Basic test",
			s1:       "ABCDGH",
			s2:       "AEDFHR",
			expected: "ADH",
		},
		{
			name:     "No common subsequence",
			s1:       "ABC",
			s2:       "DEF",
			expected: "",
		},
		{
			name:     "One string is empty",
			s1:       "ABCDGH",
			s2:       "",
			expected: "",
		},
		{
			name:     "Identical strings",
			s1:       "ABCDGH",
			s2:       "ABCDGH",
			expected: "ABCDGH",
		},
		{
			name:     "Longer strings",
			s1:       "AGGTAB",
			s2:       "GXTXAYB",
			expected: "GTAB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestCommonSubsequenceString(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("LongestCommonSubsequenceString() = %v, want %v", result, tt.expected)
			}
		})
	}
}
