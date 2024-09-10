package stringmatching

import (
	"reflect"
	"testing"
)

func TestStringMatching(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		pattern  string
		expected []int
	}{
		{
			name:     "Simple match",
			text:     "ABABDABACDABABCABAB",
			pattern:  "ABABCABAB",
			expected: []int{10},
		},
		{
			name:     "Multiple matches",
			text:     "ABABABABAB",
			pattern:  "ABAB",
			expected: []int{0, 2, 4, 6},
		},
		{
			name:     "No match",
			text:     "ABCDEFGHIJKLMNOP",
			pattern:  "XYZ",
			expected: []int{},
		},
		{
			name:     "Pattern longer than text",
			text:     "ABC",
			pattern:  "ABCDEF",
			expected: []int{},
		},
		{
			name:     "Empty pattern",
			text:     "ABCDEF",
			pattern:  "",
			expected: []int{},
		},
		{
			name:     "Empty text",
			text:     "",
			pattern:  "ABC",
			expected: []int{},
		},
		{
			name:     "Single character match",
			text:     "AAAA",
			pattern:  "A",
			expected: []int{0, 1, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name+" (Naive)", func(t *testing.T) {
			result := NaiveStringMatch(tc.text, tc.pattern)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("NaiveStringMatch(%q, %q) = %v, want %v", tc.text, tc.pattern, result, tc.expected)
			}
		})

		t.Run(tc.name+" (KMP)", func(t *testing.T) {
			result := KMPStringMatch(tc.text, tc.pattern)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("KMPStringMatch(%q, %q) = %v, want %v", tc.text, tc.pattern, result, tc.expected)
			}
		})
	}
}
