package coinchange

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount + 1)

	for i := range dp {
		dp[i] = amount + 1 // We can't use MaxInt value. Watch out for overflow: dp[i - coin] + 1
	}

	// Base case
	dp[0] = 0

	// Fil the dp table
	for _, coin := range coins{
		for i := coin; i <= amount; i++{k
			dp[i] = min(dp[i], dp[i - coin] + 1)
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}