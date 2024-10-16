package knapsack

// Item represents an item with a weight and a value
type Item struct {
	Weight int
	Value  int
}

// Knapsack solves the 0/1 Knapsack problem using dynamic programming
func Knapsack(items []Item, capacity int) int {
	// We create a table with items as row and capacity as columns
	n := len(items)
	dp := make([][]int, n+1) // Use use n + 1 for item 0 or empty item

	// We add the rows
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, capacity+1) // Same, we have capacity 0 or empty capacity, so we need to add 1 to account for it
	}

	// We need to construct the table
	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			// If the prev item capacity doesn't exceed the current capacity that we are trying to fill.
			prevItem := items[i-1]
			if prevItem.Weight <= w {
				valueIfNotTaken := dp[i-1][w] // the value from the previous row, in the upwards direction. As if we don't take the prev item
				valueIfTaken := dp[i-1][w-prevItem.Weight] + prevItem.Value
				dp[i][w] = max(valueIfNotTaken, valueIfTaken)
			} else {
				// The value remains the same, because we don't take the previous items as the weight exceeds the max weight capacity
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

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			prevItem := items[i-1]

			if prevItem.Weight <= w {
				valueIfNotTaken := dp[i-1][w]
				valueIfTaken := dp[i-1][w-prevItem.Weight] + prevItem.Value
				dp[i][w] = max(valueIfNotTaken, valueIfTaken)
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// We need to reconstruct the path taken to calculate the max value
	// Backtrack to find the selected items
	selectedItems := make([]Item, 0)
	w := capacity

	for i := n; i > 0 && w > 0; i-- {
		if dp[i][w] != dp[i-1][w] {
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
