# Red-Black Tree

## Problem Description

Implement a Red-Black Tree, which is a self-balancing binary search tree. Each node in the tree has a color attribute, which is either red or black. The tree maintains the following properties to ensure balance:

1. Every node is either red or black.
2. The root is black.
3. Every leaf (NIL) is black.
4. If a node is red, then both its children are black.
5. For each node, all simple paths from the node to descendant leaves contain the same number of black nodes.

These properties ensure that the tree remains approximately balanced during insertions and deletions.

## Operations

1. `Insert(key int)`: Inserts a new key into the Red-Black tree.
2. `Search(key int) bool`: Searches for a key in the Red-Black tree.
3. `InorderTraversal() []int`: Performs an inorder traversal of the Red-Black tree.
4. `PrintTree()`: Prints the tree structure for debugging purposes.

## Approach

### Insertion
1. Perform a standard BST insertion.
2. Color the new node red.
3. Perform a fixup operation to maintain Red-Black tree properties.

### Fixup after Insertion
1. While the parent of the current node is red:
   - If the uncle is red, recolor and move up the tree.
   - If the uncle is black, perform rotations and recolor.
2. Color the root black.

### Search
Perform a standard BST search.

### Inorder Traversal
Perform a standard inorder traversal.

## Complexity Analysis

- Time Complexity:
  - Insert: O(log n)
  - Search: O(log n)
  - Inorder Traversal: O(n)
- Space Complexity:
  - Insert: O(log n) due to recursion in fixup
  - Search: O(log n) due to recursion
  - Inorder Traversal: O(n) to store the result, O(log n) for recursion stack

Where n is the number of nodes in the Red-Black tree.

## Usage

```go
tree := NewRedBlackTree()

// Insert keys
tree.Insert(7)
tree.Insert(3)
tree.Insert(18)

// Search for a key
found := tree.Search(3) // true
notFound := tree.Search(10) // false

// Get inorder traversal
inorder := tree.InorderTraversal() // [3, 7, 18]

// Print tree structure
tree.PrintTree()
```

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:
- Insertion of multiple keys
- Inorder traversal verification
- Search functionality
- Validation of Red-Black tree properties

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of all Red-Black tree operations and properties.
