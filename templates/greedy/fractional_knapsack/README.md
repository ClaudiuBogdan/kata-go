# Greedy Approach to the Fractional Knapsack Problem

## Problem Description

The Fractional Knapsack Problem is a optimization problem where we need to fill a knapsack with fractions of items to maximize the total value, given that each item has a weight and a value. Unlike the 0/1 Knapsack Problem, we can take fractions of items, making it suitable for a greedy approach.

## Approach

This implementation uses a greedy algorithm to solve the Fractional Knapsack Problem. The algorithm works as follows:

1. Calculate the value-to-weight ratio for each item.
2. Sort the items in descending order of their value-to-weight ratios.
3. Iterate through the sorted items:
   a. If the entire item can fit, add it to the knapsack.
   b. If only a fraction can fit, add that fraction and fill the knapsack.
   c. Stop when the knapsack is full or all items have been considered.

This greedy approach ensures that we always consider the most valuable items (relative to their weight) first.

## Complexity Analysis

- Time Complexity: O(n log n), where n is the number of items. This is due to the sorting step.
- Space Complexity: O(n) to store the ratios and fractions.

## Usage

```go
items := []Item{
    {Weight: 10, Value: 60},
    {Weight: 20, Value: 100},
    {Weight: 30, Value: 120},
}
capacity := 50.0

totalValue, fractions := FractionalKnapsack(items, capacity)

fmt.Printf("Total value: %.2f\n", totalValue)
fmt.Printf("Fractions of each item: %v\n", fractions)
```

## Implementation Details

The package provides the following main components:

1. `Item` struct: Represents an item with its weight and value.
2. `FractionalKnapsack(items []Item, capacity float64) (float64, []float64)`: Solves the fractional knapsack problem and returns the total value and fractions of each item to include.

The implementation uses Go's built-in `sort` package to sort items by their value-to-weight ratio in descending order.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Simple case with mixed fractions
2. All items fitting in the knapsack
3. Only fractions of items being used
4. Empty input
5. Zero capacity knapsack

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the fractional knapsack implementation under different scenarios.

## Advantages and Limitations

Advantages:
- Efficiently solves the fractional knapsack problem
- Always produces the optimal solution
- Simple to understand and implement
- Linear time complexity (excluding sorting)

Limitations:
- Only applicable to the fractional knapsack problem, not the 0/1 knapsack problem
- Requires the ability to take fractions of items, which may not be realistic in some scenarios

## Applications

- Resource allocation in computing systems
- Cargo loading and logistics
- Financial portfolio optimization
- Cutting stock problems in manufacturing
- Bandwidth allocation in networks

Remember that while this greedy approach works optimally for the Fractional Knapsack Problem, it cannot be directly applied to the 0/1 Knapsack Problem or other variants where items cannot be fractionally selected.

