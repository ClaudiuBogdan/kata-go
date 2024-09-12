package coinchange

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{
			name:     "Example 1",
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			name:     "Example 2",
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			name:     "Example 3",
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
		{
			name:     "Large amount",
			coins:    []int{1, 2, 5, 10, 20, 50, 100},
			amount:   9876,
			expected: 102,
		},
		{
			name:     "No solution",
			coins:    []int{2, 5, 10},
			amount:   3,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CoinChange(tt.coins, tt.amount)
			if result != tt.expected {
				t.Errorf("CoinChange() = %v, want %v", result, tt.expected)
			}
		})
	}
}
