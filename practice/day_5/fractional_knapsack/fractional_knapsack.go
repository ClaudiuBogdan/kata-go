package fractionalknapsack

import "sort"

// Item represents an item with its weight and value
type Item struct {
	Weight float64
	Value  float64
}

type ratio struct {
	index int
	ratio float64
}

// FractionalKnapsack solves the fractional knapsack problem
func FractionalKnapsack(items []Item, capacity float64) (float64, []float64) {
	// Return totalValue and the fractions used
	fractions := make([]float64, len(items))
	ratios := make([]ratio, len(items))

	// I need to calculate the ratios and sort them
	for index, item := range items {
		ratios[index] = ratio{
			index: index,
			ratio: item.Value / item.Weight,
		}
	}

	sort.Slice(ratios, func(i, j int) bool {
		return ratios[i].ratio > ratios[j].ratio
	})

	totalValue := 0.0
	remainingCapacity := capacity

	// Iterate over the items
	for _, r := range ratios {
		item := items[r.index]

		if remainingCapacity <= 0 {
			break
		}

		if item.Weight <= remainingCapacity {
			fractions[r.index] = 1.0
			totalValue += item.Value
			remainingCapacity -= item.Weight
		} else {
			fraction := remainingCapacity / item.Weight
			fractions[r.index] = fraction
			totalValue += fraction * item.Value
			remainingCapacity = 0
		}
	}

	return totalValue, fractions
}
