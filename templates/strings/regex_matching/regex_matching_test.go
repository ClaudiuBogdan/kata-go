package regexmatching

import (
	"reflect"
	"testing"
)

func TestIsMatch(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		pattern  string
		expected bool
	}{
		{"Simple match", "aa", "a", false},
		{"Star match", "aa", "a*", true},
		{"Dot and star", "ab", ".*", true},
		{"Complex pattern", "aab", "c*a*b", true},
		{"No match", "mississippi", "mis*is*p*.", false},
		{"Empty pattern", "abc", "", false},
		{"Empty text", "", "a*", true},
		{"Dot match", "abc", "a.c", true},
		{"Star and dot", "aaa", "a*.", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsMatch(tc.text, tc.pattern)
			if result != tc.expected {
				t.Errorf("IsMatch(%q, %q) = %v, want %v", tc.text, tc.pattern, result, tc.expected)
			}
		})
	}
}

func TestFindAll(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		pattern  string
		expected []string
	}{
		{"Simple pattern", "ababab", "ab", []string{"ab", "ab", "ab"}},
		{"Star pattern", "aaabbbaaabbb", "a*b*", []string{"aaabbb", "aaabbb"}},
		{"Dot pattern", "abcabcabc", "a.c", []string{"abc", "abc", "abc"}},
		{"No match", "abcdef", "xyz", []string{}},
		{"Complex pattern", "aabbccaabbcc", "a*b*c*", []string{"aabbcc", "aabbcc"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FindAll(tc.text, tc.pattern)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("FindAll(%q, %q) = %v, want %v", tc.text, tc.pattern, result, tc.expected)
			}
		})
	}
}
