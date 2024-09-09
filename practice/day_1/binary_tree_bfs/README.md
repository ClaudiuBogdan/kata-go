
# Binary Tree BFS (Breadth-First Search)

## Init

```bash
./practice_kata.sh copy-kata binary_tree/binary_tree_bfs
```

## Problem Description

Implement a function that performs a Breadth-First Search (BFS) traversal on a binary tree. The function should return an array of node values in the order they are visited during the BFS traversal.

## Approach

1. Create a queue and enqueue the root node.
2. While the queue is not empty:
   a. Dequeue a node and add its value to the result array.
   b. Enqueue the left child if it exists.
   c. Enqueue the right child if it exists.
3. Return the result array.

## Complexity Analysis

- Time Complexity: O(n), where n is the number of nodes in the binary tree. We visit each node exactly once.
- Space Complexity: O(w), where w is the maximum width of the binary tree. In the worst case (a complete binary tree), this can be up to n/2 for the last level.

## Usage

```go
root := &TreeNode{
    Val: 1,
    Left: &TreeNode{Val: 2},
    Right: &TreeNode{Val: 3},
}
result := BFS(root)
fmt.Println(result) // Output: [1, 2, 3]
