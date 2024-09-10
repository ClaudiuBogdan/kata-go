package kadane

import "math"

// MaxSubarraySum finds the maximum subarray sum using Kadane's algorithm
func MaxSubarraySum(arr []int) int {
    maxSoFar := math.MinInt32
    maxEndingHere := 0

    for _, num := range arr {
        maxEndingHere = max(num, maxEndingHere+num)
        maxSoFar = max(maxSoFar, maxEndingHere)
    }

    return maxSoFar
}

// MaxSubarraySumWithIndices finds the maximum subarray sum and its indices using Kadane's algorithm
func MaxSubarraySumWithIndices(arr []int) (int, int, int) {
    maxSoFar := math.MinInt32
    maxEndingHere := 0
    start, end, tempStart := 0, 0, 0

    for i, num := range arr {
        if num > maxEndingHere+num {
            maxEndingHere = num
            tempStart = i
        } else {
            maxEndingHere += num
        }

        if maxEndingHere > maxSoFar {
            maxSoFar = maxEndingHere
            start = tempStart
            end = i
        }
    }

    return maxSoFar, start, end
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
