# Binary Search Tree Operations

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

## Implementation Details

The package provides four main functions:

1. `Insert(root *TreeNode, val int) *TreeNode`: Inserts a new value into the BST and returns the root of the updated tree.
2. `Search(root *TreeNode, val int) bool`: Searches for a value in the BST and returns true if found, false otherwise.
3. `Delete(root *TreeNode, val int) *TreeNode`: Removes a value from the BST and returns the root of the updated tree.
4. `InorderTraversal(root *TreeNode) []int`: Performs an inorder traversal of the BST and returns a slice of node values in order.

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

The implementation includes a comprehensive test suite that covers various scenarios for each operation:

1. Insertion: inserting into an empty tree, inserting smaller and larger values, inserting duplicates
2. Search: searching in an empty tree, searching for existing and non-existing values
3. Deletion: deleting from an empty tree, deleting leaf nodes, nodes with one child, nodes with two children, and the root node
4. Inorder Traversal: traversing empty trees, trees with one node, and complex trees

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of all BST operations under different scenarios.

## Advantages of Binary Search Trees

1. Efficient searching, insertion, and deletion operations (O(log n) on average for balanced trees)
2. Maintains sorted order of elements
3. Allows for range queries
4. Can be easily augmented to support additional operations

## Applications

- Implementing associative arrays (symbol tables)
- Database indexing
- Priority queues
- Syntax trees in compilers
- Huffman coding trees in data compression

## Visual Representation

Here's a visual representation of a binary search tree:

```
     5
   /   \
  3     7
 / \   / \
1   4 6   8
```

## Variations and Extensions

1. Self-balancing BSTs: AVL trees, Red-Black trees
2. Threaded binary trees: for efficient inorder traversal without recursion
3. Splay trees: for improved performance on frequently accessed elements
4. B-trees: for efficient storage on secondary storage devices

## Implementation Notes

- The implementation uses recursive approaches for simplicity, but iterative versions can be more efficient in practice.
- The deletion operation uses the inorder successor for nodes with two children, but using the inorder predecessor is also valid.
- This implementation does not handle duplicate values. Depending on requirements, duplicates could be allowed by adding an additional data structure at each node or by defining a convention for duplicate placement.

## Limitations

- Performance degrades to O(n) for skewed trees (e.g., inserting sorted data).
- Does not guarantee O(log n) performance without additional balancing mechanisms.
- Not cache-friendly due to pointer-based structure and non-contiguous memory allocation.

## Related Problems

- Validate Binary Search Tree
- Kth Smallest Element in a BST
- Lowest Common Ancestor of a Binary Search Tree
- Convert Sorted Array to Binary Search Tree
- Balanced Binary Tree

Remember that while BSTs provide efficient operations for many scenarios, other data structures like hash tables or balanced trees (e.g., AVL, Red-Black) might be more appropriate depending on specific requirements for insertion/deletion frequency, memory constraints, or guaranteed worst-case performance.
