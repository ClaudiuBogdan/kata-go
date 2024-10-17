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
			prevItem := items[i-1]
			if prevItem.Weight <= w {
				valueIfTaken := dp[i-1][w-prevItem.Weight] + prevItem.Value
				valueIfNotTaken := dp[i-1][w]
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
			prevItem := items[i-1]
			if prevItem.Weight <= w {
				valueIfTaken := dp[i-1][w-prevItem.Weight] + prevItem.Value
				valueIfNotTaken := dp[i-1][w]
				dp[i][w] = max(valueIfNotTaken, valueIfTaken)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	selectedItems := make([]Item, 0)
	w := capacity
	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
			// TODO: append the prev items, as i is out of bound from items i = n
			prevItem := items[i-1]
			selectedItems = append(selectedItems, prevItem)
			w -= prevItem.Weight
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
