package knapsack

import (
	"reflect"
	"testing"
)

func TestKnapsack(t *testing.T) {
	tests := []struct {
		name     string
		items    []Item
		capacity int
		expected int
	}{
		{
			name: "Basic test",
			items: []Item{
				{Weight: 10, Value: 60},
				{Weight: 20, Value: 100},
				{Weight: 30, Value: 120},
			},
			capacity: 50,
			expected: 220,
		},
		{
			name: "No items fit",
			items: []Item{
				{Weight: 10, Value: 60},
				{Weight: 20, Value: 100},
			},
			capacity: 5,
			expected: 0,
		},
		{
			name: "All items fit",
			items: []Item{
				{Weight: 5, Value: 50},
				{Weight: 10, Value: 60},
				{Weight: 15, Value: 90},
			},
			capacity: 30,
			expected: 200,
		},
		{
			name:     "Empty knapsack",
			items:    []Item{},
			capacity: 50,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Knapsack(tt.items, tt.capacity)
			if result != tt.expected {
				t.Errorf("Knapsack() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestKnapsackWithItems(t *testing.T) {
	tests := []struct {
		name            string
		items           []Item
		capacity        int
		expectedValue   int
		expectedItems   []Item
	}{
		{
			name: "Basic test",
			items: []Item{
				{Weight: 10, Value: 60},
				{Weight: 20, Value: 100},
				{Weight: 30, Value: 120},
			},
			capacity:      50,
			expectedValue: 220,
			expectedItems: []Item{
				{Weight: 30, Value: 120},
				{Weight: 20, Value: 100},
			},
		},
		{
			name: "No items fit",
			items: []Item{
				{Weight: 10, Value: 60},
				{Weight: 20, Value: 100},
			},
			capacity:      5,
			expectedValue: 0,
			expectedItems: []Item{},
		},
		{
			name: "All items fit",
			items: []Item{
				{Weight: 5, Value: 50},
				{Weight: 10, Value: 60},
				{Weight: 15, Value: 90},
			},
			capacity:      30,
			expectedValue: 200,
			expectedItems: []Item{
				{Weight: 15, Value: 90},
				{Weight: 10, Value: 60},
				{Weight: 5, Value: 50},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, items := KnapsackWithItems(tt.items, tt.capacity)
			if value != tt.expectedValue {
				t.Errorf("KnapsackWithItems() value = %v, want %v", value, tt.expectedValue)
			}
			if !reflect.DeepEqual(items, tt.expectedItems) {
				t.Errorf("KnapsackWithItems() items = %v, want %v", items, tt.expectedItems)
			}
		})
	}
}
