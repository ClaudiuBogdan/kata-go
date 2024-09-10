package fibonacci

import (
	"reflect"
	"testing"
)

func TestGenerateFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "Zero terms",
			n:        0,
			expected: []int{},
		},
		{
			name:     "One term",
			n:        1,
			expected: []int{0},
		},
		{
			name:     "Two terms",
			n:        2,
			expected: []int{0, 1},
		},
		{
			name:     "Five terms",
			n:        5,
			expected: []int{0, 1, 1, 2, 3},
		},
		{
			name:     "Ten terms",
			n:        10,
			expected: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateFibonacci(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GenerateFibonacci() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFibonacciTerm(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "0th term",
			n:        0,
			expected: 0,
		},
		{
			name:     "1st term",
			n:        1,
			expected: 0,
		},
		{
			name:     "2nd term",
			n:        2,
			expected: 1,
		},
		{
			name:     "5th term",
			n:        5,
			expected: 3,
		},
		{
			name:     "10th term",
			n:        10,
			expected: 34,
		},
		{
			name:     "20th term",
			n:        20,
			expected: 4181,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FibonacciTerm(tt.n)
			if result != tt.expected {
				t.Errorf("FibonacciTerm() = %v, want %v", result, tt.expected)
			}
		})
	}
}
