# AVL Tree

## Problem Description

Implement an AVL (Adelson-Velsky and Landis) tree, which is a self-balancing binary search tree. In an AVL tree, the heights of the two child subtrees of any node differ by at most one.

## Approach

1. Implement the basic binary search tree operations: insertion, deletion, and search.
2. Add balancing mechanisms to maintain the AVL property:
   - Keep track of the height of each node.
   - Perform rotations (left rotation, right rotation) when the balance factor of a node becomes greater than 1 or less than -1.
3. Implement four cases of rotations:
   - Left-Left case
   - Right-Right case
   - Left-Right case
   - Right-Left case

## Operations

1. `Insert(key int)`: Inserts a new key into the AVL tree.
2. `Delete(key int)`: Removes a key from the AVL tree.
3. `Search(key int) bool`: Searches for a key in the AVL tree.

## Complexity Analysis

- Time Complexity:
  - Insert: O(log n)
  - Delete: O(log n)
  - Search: O(log n)
- Space Complexity: O(n) for storing n nodes

Where n is the number of nodes in the AVL tree.

## Usage

```go
tree := NewAVLTree()

// Insert keys
tree.Insert(10)
tree.Insert(20)
tree.Insert(30)

// Search for a key
found := tree.Search(20) // true

// Delete a key
tree.Delete(20)

found = tree.Search(20) // false
```

## Testing

Run the tests using the following command:

```bash
go test
```

The test file includes various test cases to ensure the correct implementation of the AVL tree operations and rotations.
