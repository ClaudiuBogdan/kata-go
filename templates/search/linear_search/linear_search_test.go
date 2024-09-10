package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
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
			name:     "Duplicate elements, return first occurrence",
			arr:      []int{1, 2, 3, 4, 2, 5},
			target:   2,
			expected: 1,
		},
		{
			name:     "Large unsorted array",
			arr:      []int{15, 3, 27, 8, 19, 11, 23, 5, 17, 9, 21, 13, 25, 7, 29},
			target:   23,
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LinearSearch(tt.arr, tt.target)
			if result != tt.expected {
				t.Errorf("LinearSearch() = %v, want %v", result, tt.expected)
			}
		})
	}
}
