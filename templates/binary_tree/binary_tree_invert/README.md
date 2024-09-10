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

The package provides the main function:

`InvertTree(root *TreeNode) *TreeNode`: Inverts the given binary tree and returns the root of the inverted tree.

The function takes the root of a binary tree as input and returns the root of the inverted tree. It uses a recursive approach to swap the left and right children of each node in the tree.

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

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the `InvertTree` function under different tree structures.

## Advantages of Tree Inversion

- Useful for creating mirror images of binary trees.
- Can be used to check if two binary trees are mirror images of each other.
- Helps in understanding and implementing tree transformations.

## Applications

- Image processing: Flipping images horizontally.
- Compiler design: Transforming abstract syntax trees.
- Game development: Creating symmetric levels or structures.

## Visual Representation

Here's a visual representation of the tree inversion process:

```
     4                 4
   /   \             /   \
  2     7    =>     7     2
 / \   / \         / \   / \
1   3 6   9       9   6 3   1
```

## Iterative Approach

While the recursive approach is intuitive, an iterative approach using a stack or queue can also be implemented. This can be beneficial for very deep trees to avoid stack overflow issues.

```go
func InvertTreeIterative(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        node.Left, node.Right = node.Right, node.Left
        
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
    
    return root
}
```

## Related Problems

- Symmetric Tree: Check if a binary tree is a mirror of itself.
- Same Tree: Determine if two binary trees are identical.
- Subtree of Another Tree: Check if a binary tree is a subtree of another binary tree.

Remember that while this implementation works on binary trees, the concept of inversion can be extended to other tree-like structures or even graphs, depending on the specific requirements of your project.
