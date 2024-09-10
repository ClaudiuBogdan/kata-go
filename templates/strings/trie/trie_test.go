package trie

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie()

	// Test Insert and Search
	words := []string{"apple", "app", "apricot", "banana", "bat"}
	for _, word := range words {
		trie.Insert(word)
	}

	testCases := []struct {
		word     string
		expected bool
	}{
		{"apple", true},
		{"app", true},
		{"apricot", true},
		{"banana", true},
		{"bat", true},
		{"ap", false},
		{"banan", false},
		{"bananas", false},
		{"cat", false},
	}

	for _, tc := range testCases {
		t.Run("Search_"+tc.word, func(t *testing.T) {
			if result := trie.Search(tc.word); result != tc.expected {
				t.Errorf("Search(%q) = %v, want %v", tc.word, result, tc.expected)
			}
		})
	}

	// Test StartsWith
	prefixCases := []struct {
		prefix   string
		expected bool
	}{
		{"app", true},
		{"apr", true},
		{"ba", true},
		{"b", true},
		{"banan", true},
		{"cat", false},
		{"", true}, // empty prefix should return true
	}

	for _, tc := range prefixCases {
		t.Run("StartsWith_"+tc.prefix, func(t *testing.T) {
			if result := trie.StartsWith(tc.prefix); result != tc.expected {
				t.Errorf("StartsWith(%q) = %v, want %v", tc.prefix, result, tc.expected)
			}
		})
	}
}
