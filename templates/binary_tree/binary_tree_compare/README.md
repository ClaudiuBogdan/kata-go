# Binary Tree Compare

## Problem Description

Implement an algorithm to compare two binary trees and determine if they are structurally identical and have the same node values. Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

## Approach

This implementation uses a recursive approach to compare two binary trees. The algorithm works as follows:

1. If both nodes are nil, return true (trees are identical up to this point).
2. If one node is nil and the other isn't, return false (trees are not identical).
3. If the values of the current nodes are different, return false.
4. Recursively compare the left subtrees and right subtrees.
5. Return true if all comparisons are true.

This approach ensures that we traverse both trees simultaneously, comparing each node and its structure.

## Complexity Analysis

- Time Complexity: O(min(n, m)), where n and m are the number of nodes in the two trees. We traverse each node of the smaller tree.
- Space Complexity: O(min(h1, h2)), where h1 and h2 are the heights of the trees. This space is used by the recursion stack.

## Usage

```go
// Create two binary trees
tree1 := NewNode(1)
tree1.Left = NewNode(2)
tree1.Right = NewNode(3)

tree2 := NewNode(1)
tree2.Left = NewNode(2)
tree2.Right = NewNode(3)

// Compare the trees
areEqual := CompareTrees(tree1, tree2)
fmt.Printf("Trees are equal: %v\n", areEqual)
```

## Implementation Details

The package provides the following main components:

1. `TreeNode` struct: Represents a node in a binary tree with an integer value and left and right child pointers.
2. `CompareTrees(p *TreeNode, q *TreeNode) bool`: Compares two binary trees and returns true if they are identical, false otherwise.
3. `NewNode(val int) *TreeNode`: Helper function to create a new tree node with a given value.

The implementation uses recursion to traverse and compare the trees.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Identical trees
2. Trees with different values
3. Trees with different structures
4. Comparison with nil trees
5. Single node trees

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the binary tree comparison implementation under different scenarios.

## Advantages and Limitations

Advantages:
- Simple and intuitive implementation
- Works for any binary tree structure
- Efficiently compares trees by short-circuiting on the first difference found

Limitations:
- Recursive approach may lead to stack overflow for very deep trees
- Not suitable for comparing very large trees due to potential stack overflow

## Applications

- Verifying the correctness of tree operations (e.g., after cloning or serialization/deserialization)
- Checking for changes in hierarchical data structures
- Implementing equality checks for tree-based data structures
- Debugging tree manipulation algorithms

Remember that while this implementation works for binary trees with integer values, it can be easily adapted for trees with different types of values by modifying the `TreeNode` struct and comparison logic.

