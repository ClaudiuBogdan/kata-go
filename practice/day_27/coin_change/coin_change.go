package coinchange

// CoinChange Given an amount of money and an array of coin denominations,
// find the minimum number of coins needed to make up that amount.
// If it's not possible to make the amount using the given coin denominations, return -1.
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0;

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
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
	}
	return b
}
