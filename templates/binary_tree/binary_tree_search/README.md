# Binary Tree Search Operations

## Problem Description

Implement the basic operations of a binary search tree (BST), including insertion, search, deletion, and inorder traversal. A binary search tree is a binary tree where for each node, all elements in the left subtree are less than the node, and all elements in the right subtree are greater than the node.

## Operations

1. `Insert(root *TreeNode, val int) *TreeNode`: Inserts a new value into the BST.
2. `Search(root *TreeNode, val int) bool`: Searches for a value in the BST.
3. `Delete(root *TreeNode, val int) *TreeNode`: Removes a value from the BST.
4. `InorderTraversal(root *TreeNode) []int`: Performs an inorder traversal of the BST.

## Approach

### Insertion
1. If the tree is empty, create a new node with the given value.
2. If the value is less than the current node's value, recursively insert into the left subtree.
3. If the value is greater than the current node's value, recursively insert into the right subtree.

### Search
1. If the tree is empty or the current node's value matches the search value, return the result.
2. If the search value is less than the current node's value, recursively search the left subtree.
3. If the search value is greater than the current node's value, recursively search the right subtree.

### Deletion
1. If the tree is empty, return nil.
2. If the value to delete is less than the current node's value, recursively delete from the left subtree.
3. If the value to delete is greater than the current node's value, recursively delete from the right subtree.
4. If the value to delete matches the current node's value:
   - If the node has no children or only one child, return the non-nil child (or nil if no children).
   - If the node has two children, find the inorder successor (minimum value in the right subtree), replace the current node's value with the successor's value, and recursively delete the successor from the right subtree.

### Inorder Traversal
1. Recursively traverse the left subtree.
2. Visit the current node (add its value to the result).
3. Recursively traverse the right subtree.

## Complexity Analysis

- Time Complexity:
  - Insert: O(h) average case, O(n) worst case
  - Search: O(h) average case, O(n) worst case
  - Delete: O(h) average case, O(n) worst case
  - Inorder Traversal: O(n)
- Space Complexity:
  - Insert: O(h) due to recursion
  - Search: O(h) due to recursion
  - Delete: O(h) due to recursion
  - Inorder Traversal: O(n) to store the result, O(h) for recursion stack

Where h is the height of the tree and n is the number of nodes. In a balanced BST, h = log(n), but in the worst case (skewed tree), h can be equal to n.

## Usage

```go
root := &TreeNode{Val: 5}
root = Insert(root, 3)
root = Insert(root, 7)
root = Insert(root, 1)

found := Search(root, 3) // true
notFound := Search(root, 6) // false

root = Delete(root, 3)

inorder := InorderTraversal(root) // [1, 5, 7]
```

## Testing

The implementation includes a comprehensive test suite that covers various scenarios for each operation. To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of all BST operations.
