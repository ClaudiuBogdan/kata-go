package kadane

import (
	"math"
	"testing"
)

func TestMaxSubarraySum(t *testing.T) {
    testCases := []struct {
        name     string
        input    []int
        expected int
    }{
        {
            name:     "Basic case",
            input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
            expected: 6,
        },
        {
            name:     "All negative numbers",
            input:    []int{-1, -2, -3, -4},
            expected: -1,
        },
        {
            name:     "All positive numbers",
            input:    []int{1, 2, 3, 4},
            expected: 10,
        },
        {
            name:     "Mixed positive and negative",
            input:    []int{-2, -3, 4, -1, -2, 1, 5, -3},
            expected: 7,
        },
        {
            name:     "Single element array",
            input:    []int{5},
            expected: 5,
        },
        {
            name:     "Empty array",
            input:    []int{},
            expected: math.MinInt32,
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

func TestMaxSubarraySumWithIndices(t *testing.T) {
    testCases := []struct {
        name           string
        input          []int
        expectedSum    int
        expectedStart  int
        expectedEnd    int
    }{
        {
            name:           "Basic case",
            input:          []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
            expectedSum:    6,
            expectedStart:  3,
            expectedEnd:    6,
        },
        {
            name:           "All negative numbers",
            input:          []int{-1, -2, -3, -4},
            expectedSum:    -1,
            expectedStart:  0,
            expectedEnd:    0,
        },
        {
            name:           "All positive numbers",
            input:          []int{1, 2, 3, 4},
            expectedSum:    10,
            expectedStart:  0,
            expectedEnd:    3,
        },
        {
            name:           "Mixed positive and negative",
            input:          []int{-2, -3, 4, -1, -2, 1, 5, -3},
            expectedSum:    7,
            expectedStart:  2,
            expectedEnd:    6,
        },
        {
            name:           "Single element array",
            input:          []int{5},
            expectedSum:    5,
            expectedStart:  0,
            expectedEnd:    0,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            sum, start, end := MaxSubarraySumWithIndices(tc.input)
            if sum != tc.expectedSum || start != tc.expectedStart || end != tc.expectedEnd {
                t.Errorf("Expected sum %d, start %d, end %d, but got sum %d, start %d, end %d",
                    tc.expectedSum, tc.expectedStart, tc.expectedEnd, sum, start, end)
            }
        })
    }
}
