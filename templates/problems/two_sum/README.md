# Two Sum Problem

## Problem Description

Given an array of integers `nums` and an integer `target`, return indices of the two numbers in the array such that they add up to the `target`. You may assume that each input would have exactly one solution, and you may not use the same element twice.

## Solution Approach

We use a hash map to solve this problem efficiently:

1. Create an empty hash map to store numbers and their indices.
2. Iterate through the array:
   - Calculate the complement (target - current number).
   - If the complement exists in the hash map, return its index and the current index.
   - Otherwise, add the current number and its index to the hash map.
3. If no solution is found, return an empty slice.

## Implementation Details

The `TwoSum` function takes two parameters:
- `nums`: A slice of integers representing the input array.
- `target`: An integer representing the target sum.

It returns a slice of two integers representing the indices of the two numbers that add up to the target.

### Time Complexity
- O(n), where n is the number of elements in the input array.

### Space Complexity
- O(n), as in the worst case, we might need to store n-1 elements in the hash map.

## Usage

```go
nums := []int{2, 7, 11, 15}
target := 9
result := TwoSum(nums, target)
fmt.Println(result) // Output: [0 1]
```

## Testing

The implementation includes a test suite covering various scenarios. To run the tests:

```bash
go test
```

## Why This Solution?

This solution is optimal because:
1. It has a linear time complexity of O(n), which is the best possible for this problem.
2. It uses only one pass through the array, making it efficient for large inputs.
3. It handles various edge cases, such as no solution, multiple solutions, and negative numbers.

The use of a hash map allows for constant-time lookups, which is crucial for achieving the O(n) time complexity. This approach is much more efficient than the naive O(n^2) solution that involves checking every pair of numbers.
