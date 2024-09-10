# Two Crystal Balls Problem

## Problem Description

You are given two crystal balls that will break if dropped from high enough distance. You are in a building with N floors. Determine the exact floor from which the crystal balls will break, using the least number of drops.

## Solution Approach

We use a combination of jump search and linear search to solve this problem efficiently:

1. Calculate the optimal jump size as the square root of the number of floors.
2. Use the first crystal ball to perform a jump search, moving up by the jump size each time.
3. If the first ball breaks, use the second ball to perform a linear search in the last section.
4. If the first ball doesn't break even at the top floor, return -1 (indicating no breaking point).

## Implementation Details

The `TwoCrystalBalls` function takes a boolean slice `breaks` where `breaks[i]` is true if a crystal ball breaks when dropped from floor `i`.

### Time Complexity
- O(√n), where n is the number of floors.

### Space Complexity
- O(1), as we only use a constant amount of extra space.

## Usage

```go
breaks := []bool{false, false, false, true, true}
result := TwoCrystalBalls(breaks)
fmt.Println(result) // Output: 3
```

## Testing

The implementation includes a test suite covering various scenarios. To run the tests:

```bash
go test
```

## Why This Solution?

This solution is optimal because:
1. It minimizes the worst-case number of drops to O(√n).
2. It uses both crystal balls effectively.
3. It works well for both small and large numbers of floors.

The square root jump size balances the number of jumps with the first ball and the potential linear search with the second ball, leading to the optimal worst-case performance.
