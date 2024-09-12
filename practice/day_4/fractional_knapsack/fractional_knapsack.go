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
	n := len(items)
	fractions := make([]float64, n)

	ratios := make([]ratio, n)

	for i, item := range items {
		if item.Weight > 0 {
			ratios[i] = ratio{
				index: i,
				ratio: item.Value / item.Weight,
			}
		} else {
			ratios[i] = ratio{
				index: i,
				ratio: 0,
			}
		}
	}

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
			totalValue += item.Value * fraction
			fractions[r.index] = fraction
			remainingCapacity = 0
		}
	}

	return totalValue, fractions
}
