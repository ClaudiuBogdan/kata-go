package fractionalknapsack

import (
	"sort"
)

// Item represents an item with its weight and value
type Item struct {
	Weight float64
	Value  float64
}

// FractionalKnapsack solves the fractional knapsack problem
func FractionalKnapsack(items []Item, capacity float64) (float64, []float64) {
	n := len(items)
	if n == 0 || capacity <= 0 {
		return 0, []float64{}
	}

	// Calculate value-to-weight ratio for each item
	ratios := make([]struct {
		index int
		ratio float64
	}, n)

	for i, item := range items {
		ratios[i] = struct {
			index int
			ratio float64
		}{i, item.Value / item.Weight}
	}

	// Sort items by value-to-weight ratio in descending order
	sort.Slice(ratios, func(i, j int) bool {
		return ratios[i].ratio > ratios[j].ratio
	})

	totalValue := 0.0
	fractions := make([]float64, n)

	for _, r := range ratios {
		if capacity == 0 {
			break
		}

		item := items[r.index]
		if item.Weight <= capacity {
			fractions[r.index] = 1.0
			totalValue += item.Value
			capacity -= item.Weight
		} else {
			fraction := capacity / item.Weight
			fractions[r.index] = fraction
			totalValue += item.Value * fraction
			capacity = 0
		}
	}

	return totalValue, fractions
}
