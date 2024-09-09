# Binary Tree Traversal

## Problem Description

Implement functions that perform various traversals on a binary tree. The functions should return arrays of node values in the order they are visited during the traversal. Implement the following traversals:

1. Inorder Traversal
2. Preorder Traversal
3. Postorder Traversal
4. Level-order Traversal

## Approach

### Inorder Traversal
1. Recursively traverse the left subtree.
2. Visit the current node.
3. Recursively traverse the right subtree.

### Preorder Traversal
1. Visit the current node.
2. Recursively traverse the left subtree.
3. Recursively traverse the right subtree.

### Postorder Traversal
1. Recursively traverse the left subtree.
2. Recursively traverse the right subtree.
3. Visit the current node.

### Level-order Traversal
1. Use a queue to keep track of nodes at each level.
2. While the queue is not empty:
   a. Process all nodes at the current level.
   b. Enqueue their children for the next level.

## Complexity Analysis

For Inorder, Preorder, and Postorder traversals:
- Time Complexity: O(n), where n is the number of nodes in the binary tree.
- Space Complexity: O(h), where h is the height of the tree. In the worst case, h can be n for a skewed tree.

For Level-order traversal:
- Time Complexity: O(n)
- Space Complexity: O(w), where w is the maximum width of the tree. In the worst case, this can be n/2 for a complete binary tree.

## Usage

```go
root := &TreeNode{
    Val: 1,
    Left: &TreeNode{Val: 2},
    Right: &TreeNode{Val: 3},
}

inorder := InorderTraversal(root)
fmt.Println("Inorder:", inorder) // Output: [2, 1, 3]

preorder := PreorderTraversal(root)
fmt.Println("Preorder:", preorder) // Output: [1, 2, 3]

postorder := PostorderTraversal(root)
fmt.Println("Postorder:", postorder) // Output: [2, 3, 1]

levelOrder := LevelOrderTraversal(root)
fmt.Println("Level-order:", levelOrder) // Output: [[1], [2, 3]]
```

## Implementation Details

The package provides four main functions:

1. `InorderTraversal(root *TreeNode) []int`: Performs an inorder traversal.
2. `PreorderTraversal(root *TreeNode) []int`: Performs a preorder traversal.
3. `PostorderTraversal(root *TreeNode) []int`: Performs a postorder traversal.
4. `LevelOrderTraversal(root *TreeNode) [][]int`: Performs a level-order traversal.

The first three functions return a slice of integers, while the level-order traversal returns a slice of slices, where each inner slice represents a level in the tree.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. Normal binary tree with all traversal types
2. Empty tree
3. Tree with only root node

To run the tests, use the following command:

```bash
go test
```

## Applications of Tree Traversals

- Inorder: Used in BST to retrieve elements in sorted order.
- Preorder: Used to create a copy of the tree or to get prefix expression on of an expression tree.
- Postorder: Used to delete the tree or to get the postfix expression of an expression tree.
- Level-order: Used in serialization of binary trees, and when you need to process nodes level by level.

Remember that while these implementations return arrays of node values, the traversal algorithms can be adapted for various purposes, such as searching for a specific node, calculating tree properties, or performing operations on the tree structure.
