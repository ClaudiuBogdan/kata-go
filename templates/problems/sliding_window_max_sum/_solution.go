package slidingwindow

// MaxSubarraySum finds the maximum subarray sum using the sliding window technique
func MaxSubarraySum(arr []int) int {
    if len(arr) == 0 {
        return 0
    }

    maxSum := arr[0]
    currentSum := 0

    for _, num := range arr {
        currentSum = max(num, currentSum+num)
        maxSum = max(maxSum, currentSum)
    }

    return maxSum
}

// max returns the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}