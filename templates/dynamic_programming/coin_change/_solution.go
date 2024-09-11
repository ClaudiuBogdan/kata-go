package coinchange

import "math"

func CoinChange(coins []int, amount int) int {
	// Create a DP table to store the minimum number of coins for each amount
	dp := make([]int, amount+1)

	// Initialize the DP table with a value larger than any possible solution
	for i := range dp {
		dp[i] = amount + 1
	}

	// Base case: 0 coins needed to make amount 0
	dp[0] = 0

	// Fill the DP table
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
		}
	}

	// If dp[amount] is still amount+1, it means it's impossible to make the amount
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}