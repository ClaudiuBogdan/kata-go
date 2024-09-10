package knapsack

// Item represents an item with a weight and a value
type Item struct {
	Weight int
	Value  int
}

// Knapsack solves the 0/1 Knapsack problem using dynamic programming
func Knapsack(items []Item, capacity int) int {
	n := len(items)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if items[i-1].Weight <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-items[i-1].Weight]+items[i-1].Value)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

// KnapsackWithItems solves the 0/1 Knapsack problem and returns the maximum value
// along with the selected items
func KnapsackWithItems(items []Item, capacity int) (int, []Item) {
	n := len(items)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if items[i-1].Weight <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-items[i-1].Weight]+items[i-1].Value)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// Backtrack to find the selected items
	selectedItems := []Item{}
	w := capacity
	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			selectedItems = append(selectedItems, items[i-1])
			w -= items[i-1].Weight
		}
	}

	return dp[n][capacity], selectedItems
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
