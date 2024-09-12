package fractionalknapsack

import (
	"math"
	"testing"
)

func TestFractionalKnapsack(t *testing.T) {
	tests := []struct {
		name              string
		items             []Item
		capacity          float64
		expectedValue     float64
		expectedFractions []float64
	}{
		{
			name: "Simple case",
			items: []Item{
				{Weight: 10, Value: 60},
				{Weight: 20, Value: 100},
				{Weight: 30, Value: 120},
			},
			capacity:          50,
			expectedValue:     240,
			expectedFractions: []float64{1, 1, 0.6666666666666666},
		},
		{
			name: "All items fit",
			items: []Item{
				{Weight: 5, Value: 25},
				{Weight: 10, Value: 40},
				{Weight: 15, Value: 50},
			},
			capacity:          30,
			expectedValue:     115,
			expectedFractions: []float64{1, 1, 1},
		},
		{
			name: "Only fractions",
			items: []Item{
				{Weight: 10, Value: 50},
				{Weight: 20, Value: 80},
				{Weight: 30, Value: 100},
			},
			capacity:          15,
			expectedValue:     70,
			expectedFractions: []float64{1, 0.25, 0},
		},
		{
			name: "All fractional items",
			items: []Item{
				{Weight: 30, Value: 120},
				{Weight: 40, Value: 160},
				{Weight: 50, Value: 200},
			},
			capacity:          60,
			expectedValue:     240,
			expectedFractions: []float64{1, 0.75, 0},
		},
		{
			name:              "Empty input",
			items:             []Item{},
			capacity:          50,
			expectedValue:     0,
			expectedFractions: []float64{},
		},
		{
			name: "Zero capacity",
			items: []Item{
				{Weight: 10, Value: 60},
				{Weight: 20, Value: 100},
			},
			capacity:          0,
			expectedValue:     0,
			expectedFractions: []float64{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, fractions := FractionalKnapsack(tt.items, tt.capacity)

			if !almostEqual(value, tt.expectedValue, 1e-6) {
				t.Errorf("Expected value %.2f, but got %.2f", tt.expectedValue, value)
			}

			if len(fractions) != len(tt.expectedFractions) {
				t.Errorf("Expected %d fractions, but got %d", len(tt.expectedFractions), len(fractions))
			} else {
				for i := range fractions {
					if !almostEqual(fractions[i], tt.expectedFractions[i], 1e-6) {
						t.Errorf("For item %d, expected fraction %.6f, but got %.6f", i, tt.expectedFractions[i], fractions[i])
					}
				}
			}

			if !isValidSolution(tt.items, tt.capacity, fractions) {
				t.Errorf("Invalid solution: fractions %v exceed capacity %.2f", fractions, tt.capacity)
			}
		})
	}
}

func almostEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func isValidSolution(items []Item, capacity float64, fractions []float64) bool {
	if len(fractions) == 0 {
		return true // Consider an empty solution valid
	}
	totalWeight := 0.0
	for i, item := range items {
		totalWeight += item.Weight * fractions[i]
	}
	return totalWeight <= capacity+1e-6 // Allow for small floating-point errors
}
