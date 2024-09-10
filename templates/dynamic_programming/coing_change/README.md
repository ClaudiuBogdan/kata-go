# Coin Change Problem (Dynamic Programming Solution)

## Problem Description

Given an array of coin denominations and a target amount, find the minimum number of coins needed to make up that amount. If the amount cannot be made up by any combination of the coins, return -1.

## Approach

This solution uses dynamic programming to solve the coin change problem efficiently. The algorithm builds a table (dp array) where each entry dp[i] represents the minimum number of coins needed to make up the amount i.

### Algorithm Steps:

1. Initialize a dp array of size `amount + 1` with a value greater than the maximum possible number of coins (amount + 1).
2. Set dp[0] = 0 as the base case (0 coins needed to make up amount 0).
3. Iterate through amounts from 1 to the target amount:
   - For each coin denomination:
     - If the coin value is less than or equal to the current amount:
       - Update dp[i] with the minimum of its current value and dp[i - coin] + 1.
4. If dp[amount] is still greater than amount, return -1 (no solution).
5. Otherwise, return dp[amount] as the minimum number of coins needed.

## Complexity Analysis

- Time Complexity: O(amount * len(coins)), where amount is the target amount and len(coins) is the number of coin denominations.
- Space Complexity: O(amount) for the dp array.

## Usage

```go
coins := []int{1, 2, 5}
amount := 11
result := CoinChange(coins, amount)
fmt.Println(result) // Output: 3 (11 = 5 + 5 + 1)
```

## Implementation Details

The package provides one main function:

`CoinChange(coins []int, amount int) int`: Solves the coin change problem and returns the minimum number of coins needed or -1 if no solution exists.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Basic examples
2. Large amounts
3. Cases with no solution
4. Edge cases (e.g., amount = 0)

To run the tests, use the following command:

```bash
go test
```

## Advantages of Dynamic Programming

- Optimal substructure: The solution to the problem can be constructed from optimal solutions of its subproblems.
- Overlapping subproblems: The algorithm solves each subproblem only once and stores the results for future use, avoiding redundant computations.

## Applications

- Currency exchange systems
- Vending machine coin dispensers
- Financial planning and budgeting tools

This dynamic programming solution to the coin change problem provides an efficient way to calculate the minimum number of coins needed for any given amount, making it useful in various financial and computational applications.
