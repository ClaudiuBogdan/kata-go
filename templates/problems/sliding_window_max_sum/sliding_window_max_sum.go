package slidingwindow

import "errors"

// MaxSubarraySum finds the maximum subarray sum of size k using the sliding window technique
func MaxSubarraySum(arr []int, k int) (int, error) {
    n := len(arr)
    
    if n < k {
        return 0, errors.New("array length is less than the window size")
    }

    // Compute sum of first window of size k
    windowSum := 0
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }

    maxSum := windowSum

    // Slide the window and update the maximum sum
    for i := k; i < n; i++ {
        windowSum = windowSum - arr[i-k] + arr[i]
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }

    return maxSum, nil
}

// MaxSubarraySumWithIndices finds the maximum subarray sum of size k and its starting index
func MaxSubarraySumWithIndices(arr []int, k int) (int, int, error) {
    n := len(arr)
    
    if n < k {
        return 0, -1, errors.New("array length is less than the window size")
    }

    // Compute sum of first window of size k
    windowSum := 0
    for i := 0; i < k; i++ {
        windowSum += arr[i]
    }

    maxSum := windowSum
    maxStartIndex := 0

    // Slide the window and update the maximum sum
    for i := k; i < n; i++ {
        windowSum = windowSum - arr[i-k] + arr[i]
        if windowSum > maxSum {
            maxSum = windowSum
            maxStartIndex = i - k + 1
        }
    }

    return maxSum, maxStartIndex, nil
}
