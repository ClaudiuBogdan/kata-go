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
			item := items[i-1]
			if item.Weight <= w {
				valueIfNotTaken := dp[i-1][w]
				valueIfTaken := dp[i-1][w-item.Weight] + item.Value

				dp[i][w] = max(valueIfNotTaken, valueIfTaken)
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
			item := items[i-1]
			if item.Weight <= w {
				valueIfNotTaken := dp[i-1][w]
				valueIfTaken := dp[i-1][w-item.Weight] + item.Value

				dp[i][w] = max(valueIfNotTaken, valueIfTaken)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	selected := make([]Item, 0)
	w := capacity

	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			item := items[i-1]
			w -= item.Weight
			selected = append(selected, item)
		}
	}

	return dp[len(items)][capacity], selected
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
