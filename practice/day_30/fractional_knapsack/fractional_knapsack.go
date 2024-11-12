package fractionalknapsack

import "sort"

// Item represents an item with its weight and value
type Item struct {
	Weight float64
	Value  float64
	ratio  float64
	index  int
}

// FractionalKnapsack solves the fractional knapsack problem
func FractionalKnapsack(items []Item, capacity float64) (float64, []float64) {
	// Sort by value/weight ratio
	for i := range items {
		items[i].ratio = items[i].Value / items[i].Weight
		items[i].index = i
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ratio > items[j].ratio
	})

	// Fill the knapsack
	remainingCapacity := capacity
	totalValue := 0.0
	fractions := make([]float64, len(items))

	for i := range items {
		if remainingCapacity == 0 {
			break
		}

		item := items[i]

		if item.Weight <= remainingCapacity {
			fractions[item.index] = 1.0
			totalValue += item.Value
			remainingCapacity -= item.Weight
		} else {
			ratio := remainingCapacity / item.Weight
			fractions[item.index] = ratio
			totalValue += item.Value * ratio
			remainingCapacity = 0
		}
	}

	return totalValue, fractions
}
