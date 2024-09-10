package coinchange

// CoinChange solves the coin change problem using dynamic programming
func CoinChange(coins []int, amount int) int {
	// Initialize dp array with amount + 1 as the maximum value
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	// Base case: 0 coins needed to make amount 0
	dp[0] = 0

	// Fill the dp array
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	// If dp[amount] is still amount + 1, it means no solution was found
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
