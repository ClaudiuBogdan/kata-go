# Binary Tree Lowest Common Ancestor

## Problem Description

Implement a function that finds the lowest common ancestor (LCA) of two nodes in a binary tree. The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).

## Approach

1. If the root is nil or equal to either p or q, return the root (base case).
2. Recursively search for p and q in the left and right subtrees.
3. If both left and right recursive calls return non-nil values, the current node is the LCA.
4. If one side returns nil, return the non-nil side.

## Complexity Analysis

- Time Complexity: O(n), where n is the number of nodes in the binary tree. In the worst case, we might need to visit all nodes of the tree.
- Space Complexity: O(h), where h is the height of the tree. This space is used by the recursion stack. In the worst case (skewed tree), h can be equal to n, so the space complexity could be O(n).

## Usage

```go
root := &TreeNode{Val: 3}
root.Left = &TreeNode{Val: 5}
root.Right = &TreeNode{Val: 1}
root.Left.Left = &TreeNode{Val: 6}
root.Left.Right = &TreeNode{Val: 2}
root.Right.Left = &TreeNode{Val: 0}
root.Right.Right = &TreeNode{Val: 8}
root.Left.Right.Left = &TreeNode{Val: 7}
root.Left.Right.Right = &TreeNode{Val: 4}

lca := LowestCommonAncestor(root, root.Left, root.Right)
fmt.Println(lca.Val) // Output: 3
```

## Implementation Details

The `LowestCommonAncestor` function takes three arguments:
1. `root`: The root of the binary tree
2. `p`: The first node
3. `q`: The second node

It returns the lowest common ancestor of p and q.

The implementation uses a recursive approach to traverse the tree and find the LCA. It handles various cases, including when one node is the ancestor of the other, or when the nodes are in different subtrees.

## Testing

The implementation includes a comprehensive test suite that covers various scenarios:

1. LCA of nodes in different subtrees
2. LCA when one node is the ancestor of the other
3. LCA of the root and another node
4. LCA in a tree with only a root node

To run the tests, use the following command:

```bash
go test
```

The test file uses Go's built-in testing package and includes multiple test cases to ensure the correctness of the `LowestCommonAncestor` function.
