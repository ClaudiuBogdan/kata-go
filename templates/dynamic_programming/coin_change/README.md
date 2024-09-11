# Coin Change (Dynamic Programming)

## Problem Description

Given an amount of money and an array of coin denominations, find the minimum number of coins needed to make up that amount. If it's not possible to make the amount using the given coin denominations, return -1.

## Approach: Dynamic Programming

Dynamic Programming is an algorithmic paradigm that solves a given complex problem by breaking it into subproblems and stores the results of subproblems to avoid computing the same results again. For the coin change problem:

1. Create a DP table `dp` of size `amount + 1` to store the minimum number of coins needed for each amount from 0 to the target amount.
2. Initialize `dp[0] = 0` (it takes 0 coins to make an amount of 0) and all other values to `amount + 1` (a value larger than any possible solution).
3. For each coin denomination:
   - For each amount from the coin value to the target amount:
     - Update `dp[amount]` to be the minimum of its current value and `1 + dp[amount - coin]`.
4. If `dp[amount]` is still `amount + 1`, return -1 (impossible to make the amount).
5. Otherwise, return `dp[amount]`.

## Complexity Analysis

- Time Complexity: O(amount * n), where n is the number of coin denominations. We have two nested loops: one over all amounts up to the target amount, and for each amount, we consider all coin denominations.
- Space Complexity: O(amount) to store the DP table.

## Implementation Details

The package provides the main function:

`CoinChange(coins []int, amount int) int`: Finds the minimum number of coins needed to make up the given amount.

The function takes two parameters:

- `coins`: A slice of integers representing the available coin denominations.
- `amount`: An integer representing the target amount.

It returns an integer representing the minimum number of coins needed, or -1 if it's impossible.

## Usage

```go
coins := []int{1, 2, 5}
amount := 11
result := CoinChange(coins, amount)
fmt.Println(result) // Output: 3 (11 = 5 + 5 + 1)
```

## Testing

The implementation should include a comprehensive test suite covering various scenarios:

1. Standard case with a solution
2. Case where it's impossible to make the amount
3. Case with a large amount and many coin denominations
4. Edge cases (amount = 0, single coin denomination, etc.)

To run the tests, use the following command:

```bash
go test
```

The test file should use Go's built-in testing package and include multiple test cases to ensure the correctness of the `CoinChange` function under different scenarios.

## Advantages of Dynamic Programming for Coin Change

1. Optimal solution: Guarantees the minimum number of coins
2. Efficiency: Solves the problem in polynomial time
3. Versatility: Can handle any set of coin denominations
4. Reusability: Subproblem solutions are stored and reused

## Applications

- Currency exchange systems
- Vending machine software
- Cashier systems in retail
- Financial software for cash management

## Visual Representation

Here's a visual representation of the DP table for coins [1, 2, 5] and amount 11:

```
Amount:   0  1  2  3  4  5  6  7  8  9 10 11
DP value: 0  1  1  2  2  1  2  2  3  3  2  3
```

## Variations and Extensions

1. Coin Change 2: Count the number of ways to make the amount
2. Minimum number of coins with limited supply of each denomination
3. Maximum number of coins (greedy approach with largest denominations first)
4. Coin change with additional constraints (e.g., must use at least one of each coin)

## Implementation Notes

- The implementation assumes that the input is valid (non-negative amount, positive coin denominations).
- This solution provides the minimum number of coins. It can be extended to also return the actual coins used.
- For very large amounts or many coin denominations, consider using big integer arithmetic to prevent overflow.

## Limitations

- May be slow for very large amounts or many coin denominations due to O(amount * n) time complexity
- Does not handle negative or floating-point amounts
- Assumes an infinite supply of each coin denomination

## Related Problems

- Knapsack Problem
- Minimum Cost Path
- Rod Cutting
- Subset Sum Problem
- Partition Equal Subset Sum

Remember that while dynamic programming is very efficient for this problem, it's important to understand the problem constraints and requirements. In some variations of this problem or in different contexts, other techniques like greedy algorithms or backtracking might be more appropriate.
