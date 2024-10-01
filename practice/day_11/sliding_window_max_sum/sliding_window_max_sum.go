package slidingwindow

// MaxSubarraySum finds the maximum subarray sum using the sliding window technique
func MaxSubarraySum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	currMax := arr[0]
	globalMax := arr[0]

	// TODO: forgot to start i from 1, not 0.
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		currMax = max(currMax+val, val)
		globalMax = max(currMax, globalMax)
	}

	return globalMax
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
