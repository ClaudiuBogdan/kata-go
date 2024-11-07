package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Target in middle",
			arr:      []int{1, 3, 5, 7, 9},
			target:   5,
			expected: 2,
		},
		{
			name:     "Target at beginning",
			arr:      []int{1, 3, 5, 7, 9},
			target:   1,
			expected: 0,
		},
		{
			name:     "Target at end",
			arr:      []int{1, 3, 5, 7, 9},
			target:   9,
			expected: 4,
		},
		{
			name:     "Target not in array",
			arr:      []int{1, 3, 5, 7, 9},
			target:   6,
			expected: -1,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "Single element array, target found",
			arr:      []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "Single element array, target not found",
			arr:      []int{1},
			target:   2,
			expected: -1,
		},
		{
			name:     "Large sorted array",
			arr:      []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29},
			target:   23,
			expected: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BinarySearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("BinarySearch() = %v, want %v", result, tt.expected)
			}
		})
	}
}
