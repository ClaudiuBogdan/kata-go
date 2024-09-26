package fractionalknapsack

import (
	"sort"
)

// Item represents an item with its weight and value
type Item struct {
	Weight float64
	Value  float64
}

type ratio struct {
	index int
	value float64
}

// FractionalKnapsack solves the fractional knapsack problem
func FractionalKnapsack(items []Item, capacity float64) (float64, []float64) {
	// Create a ratio
	// Sort
	// Get items until the capacity is full
	ratios := make([]ratio, len(items))
	fractions := make([]float64, len(items))

	for i, item := range items {
		// Maybe use insertion sort for improved performance?
		ratios = append(ratios, ratio{
			index: i,
			value: item.Value/item.Weight,
		})
	}

	sort.Slice(ratios, func(i, j int) bool {
		return ratios[i].value > ratios[j].value
	})

	remainingCapacity := capacity
	totalValue := 0.0

	for _, ratio := range ratios {
		if remainingCapacity == 0 {
			break
		}

		item := items[ratio.index]

		if item.Weight < remainingCapacity {
			fractions[ratio.index] = 1.0
			totalValue += item.Value
			remainingCapacity -= item.Weight
		} else {
			fraction := remainingCapacity / item.Weight
			fractions[ratio.index] = fraction
			totalValue += item.Value * fraction
			remainingCapacity = 0
		}
		
	}

	return totalValue, fractions
}