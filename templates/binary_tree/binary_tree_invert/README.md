# Binary Tree Invert

## Problem Description

Implement a function that inverts a binary tree. Inverting a binary tree means swapping the left and right children of every node in the tree.

## Approach

1. If the root is nil, return nil (base case).
2. Swap the left and right children of the current node.
3. Recursively invert the left subtree.
4. Recursively invert the right subtree.
5. Return the root of the inverted tree.

## Complexity Analysis

- Time Complexity: O(n), where n is the number of nodes in the binary tree. We visit each node exactly once.
- Space Complexity: O(h), where h is the height of the tree. This space is used by the call stack during recursion. In the worst case (skewed tree), h can be equal to n, so the space complexity could be O(n).

## Usage

```go
root := &TreeNode{
    Val: 4,
    Left: &TreeNode{Val: 2},
    Right: &TreeNode{Val: 7},
}
invertedRoot := InvertTree(root)
```

## Implementation Details

The `InvertTree` function takes the root of a binary tree as input and returns the root of the inverted tree. It uses a recursive approach to swap the left and right children of each node in the tree.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. Normal binary tree
2. Empty tree
3. Tree with only root node
4. Unbalanced tree

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the `InvertTree` function.
