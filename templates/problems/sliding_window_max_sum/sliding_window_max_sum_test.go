package slidingwindow

import (
    "testing"
)

func TestMaxSubarraySum(t *testing.T) {
    testCases := []struct {
        name        string
        input       []int
        k           int
        expected    int
        expectError bool
    }{
        {
            name:        "Basic case",
            input:       []int{1, 4, 2, 10, 23, 3, 1, 0, 20},
            k:           4,
            expected:    39,
            expectError: false,
        },
        {
            name:        "All positive numbers",
            input:       []int{1, 2, 3, 4, 5},
            k:           3,
            expected:    12,
            expectError: false,
        },
        {
            name:        "Mixed positive and negative",
            input:       []int{-1, 2, -3, 4, -5, 6},
            k:           3,
            expected:    5,
            expectError: false,
        },
        {
            name:        "Window size equal to array length",
            input:       []int{1, 2, 3, 4, 5},
            k:           5,
            expected:    15,
            expectError: false,
        },
        {
            name:        "Window size greater than array length",
            input:       []int{1, 2, 3},
            k:           4,
            expected:    0,
            expectError: true,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result, err := MaxSubarraySum(tc.input, tc.k)
            
            if tc.expectError {
                if err == nil {
                    t.Errorf("Expected an error, but got none")
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if result != tc.expected {
                    t.Errorf("Expected %d, but got %d", tc.expected, result)
                }
            }
        })
    }
}

func TestMaxSubarraySumWithIndices(t *testing.T) {
    testCases := []struct {
        name           string
        input          []int
        k              int
        expectedSum    int
        expectedIndex  int
        expectError    bool
    }{
        {
            name:           "Basic case",
            input:          []int{1, 4, 2, 10, 23, 3, 1, 0, 20},
            k:              4,
            expectedSum:    39,
            expectedIndex:  3,
            expectError:    false,
        },
        {
            name:           "All positive numbers",
            input:          []int{1, 2, 3, 4, 5},
            k:              3,
            expectedSum:    12,
            expectedIndex:  2,
            expectError:    false,
        },
        {
            name:           "Mixed positive and negative",
            input:          []int{-1, 2, -3, 4, -5, 6},
            k:              3,
            expectedSum:    5,
            expectedIndex:  3,
            expectError:    false,
        },
        {
            name:           "Window size equal to array length",
            input:          []int{1, 2, 3, 4, 5},
            k:              5,
            expectedSum:    15,
            expectedIndex:  0,
            expectError:    false,
        },
        {
            name:           "Window size greater than array length",
            input:          []int{1, 2, 3},
            k:              4,
            expectedSum:    0,
            expectedIndex:  -1,
            expectError:    true,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            sum, index, err := MaxSubarraySumWithIndices(tc.input, tc.k)
            
            if tc.expectError {
                if err == nil {
                    t.Errorf("Expected an error, but got none")
                }
            } else {
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
                if sum != tc.expectedSum {
                    t.Errorf("Expected sum %d, but got %d", tc.expectedSum, sum)
                }
                if index != tc.expectedIndex {
                    t.Errorf("Expected index %d, but got %d", tc.expectedIndex, index)
                }
            }
        })
    }
}
