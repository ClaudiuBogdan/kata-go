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
	fractions := make([]float64, n)

	if capacity <= 0 {
		return 0, fractions
	}

	// Calculate value-to-weight ratio for each item
	ratios := make([]struct {
		index int
		ratio float64
	}, n)

	for i, item := range items {
		if item.Weight > 0 {
			ratios[i] = struct {
				index int
				ratio float64
			}{i, item.Value / item.Weight}
		} else {
			ratios[i] = struct {
				index int
				ratio float64
			}{i, 0}
		}
	}

	// Sort items by value-to-weight ratio in descending order
	sort.Slice(ratios, func(i, j int) bool {
		return ratios[i].ratio > ratios[j].ratio
	})

	totalValue := 0.0
	remainingCapacity := capacity

	for _, r := range ratios {
		if remainingCapacity <= 0 {
			break
		}

		item := items[r.index]
		if item.Weight <= remainingCapacity {
			fractions[r.index] = 1.0
			totalValue += item.Value
			remainingCapacity -= item.Weight
		} else {
			fraction := remainingCapacity / item.Weight
			fractions[r.index] = fraction
			totalValue += item.Value * fraction
			remainingCapacity = 0
		}
	}

	// Round the totalValue to 2 decimal places to avoid precision issues
	totalValue = float64(int(totalValue*100)) / 100

	return totalValue, fractions
}