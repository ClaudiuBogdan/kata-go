# Binary Tree DFS (Depth-First Search)

## Problem Description

Implement functions that perform Depth-First Search (DFS) traversals on a binary tree. The functions should return arrays of node values in the order they are visited during the DFS traversal. Implement the following DFS traversals:

1. Preorder DFS
2. Inorder DFS
3. Postorder DFS

## Approach

### Preorder DFS
1. Visit the current node.
2. Recursively traverse the left subtree.
3. Recursively traverse the right subtree.

### Inorder DFS
1. Recursively traverse the left subtree.
2. Visit the current node.
3. Recursively traverse the right subtree.

### Postorder DFS
1. Recursively traverse the left subtree.
2. Recursively traverse the right subtree.
3. Visit the current node.

## Complexity Analysis

For all DFS traversals:
- Time Complexity: O(n), where n is the number of nodes in the binary tree. We visit each node exactly once.
- Space Complexity: O(h), where h is the height of the tree. This space is used by the call stack during recursion. In the worst case (skewed tree), h can be equal to n, so the space complexity could be O(n).

## Usage

```go
root := &TreeNode{
    Val: 1,
    Left: &TreeNode{Val: 2},
    Right: &TreeNode{Val: 3},
}

preorder := PreorderDFS(root)
fmt.Println("Preorder:", preorder) // Output: [1, 2, 3]

inorder := InorderDFS(root)
fmt.Println("Inorder:", inorder) // Output: [2, 1, 3]

postorder := PostorderDFS(root)
fmt.Println("Postorder:", postorder) // Output: [2, 3, 1]
```

## Implementation Details

The package provides three main functions:

1. `PreorderDFS(root *TreeNode) []int`: Performs a preorder DFS traversal.
2. `InorderDFS(root *TreeNode) []int`: Performs an inorder DFS traversal.
3. `PostorderDFS(root *TreeNode) []int`: Performs a postorder DFS traversal.

Each function takes the root of a binary tree as input and returns a slice of integers representing the values of the nodes in the order they are visited during the respective DFS traversal.

The implementation uses recursive helper functions to perform the traversals, which allows for a clean and intuitive implementation of the DFS algorithms.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. Normal binary tree with all DFS traversals
2. Empty tree
3. Tree with only root node

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of all DFS functions under different tree structures.

## Advantages of DFS

- Uses less memory than BFS for very deep trees.
- Can be easily implemented using recursion.
- Useful for tasks such as finding the deepest node in the tree or checking if a path exists between two nodes.

## Applications

- Detecting cycles in graphs
- Topological sorting
- Solving puzzles with only one solution (e.g., maze problems)
- Finding strongly connected components in a graph

Remember that while these implementations return arrays of node values, DFS can be adapted for various purposes, such as searching for a specific node, calculating the maximum depth of the tree, or checking if the tree is balanced.
