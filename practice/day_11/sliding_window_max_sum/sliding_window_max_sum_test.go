package slidingwindow

import (
	"testing"
)

func TestMaxSubarraySum(t *testing.T) {
    testCases := []struct {
        name     string
        input    []int
        expected int
    }{
        {
            name:     "Mixed positive and negative numbers",
            input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
            expected: 6,
        },
        {
            name:     "All positive numbers",
            input:    []int{1, 2, 3, 4, 5},
            expected: 15,
        },
        {
            name:     "All negative numbers",
            input:    []int{-1, -2, -3, -4, -5},
            expected: -1,
        },
        {
            name:     "Single element array",
            input:    []int{42},
            expected: 42,
        },
        {
            name:     "Empty array",
            input:    []int{},
            expected: 0,
        },
        {
            name:     "Array with zero",
            input:    []int{-2, 0, 3, -1, 2},
            expected: 4,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := MaxSubarraySum(tc.input)
            if result != tc.expected {
                t.Errorf("Expected %d, but got %d", tc.expected, result)
            }
        })
    }
}