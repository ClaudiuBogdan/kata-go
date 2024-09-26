# Invert a Binary Tree using an iterative approach

```go
package binarytreeinvert

// TreeNode represents a node in a binary tree
type TreeNode struct {
 Val   int
 Left  *TreeNode
 Right *TreeNode
}

// InvertTree inverts a binary tree iteratively using a queue.
func InvertTree(root *TreeNode) *TreeNode {
 if root == nil {
  return nil
 }

 queue := []*TreeNode{root}

 for len(queue) > 0 {
  current := queue[0]
  queue = queue[1:] // Dequeue

  // Swap left and right children
  current.Left, current.Right = current.Right, current.Left

  // Enqueue non-nil children
  if current.Left != nil {
   queue = append(queue, current.Left)
  }
  if current.Right != nil {
   queue = append(queue, current.Right)
  }
 }

 return root
}

```

The recursive approach uses the call stack implicitly for its operation which has memory overhead because each recursive call create a new stack frame.  For a significantly deep tree, this could lead to a stack overflow error.  The iterative version, using a queue-based approach has following advantages:

* **Avoids Stack Overflow:** It uses a queue instead of the call stack, preventing stack overflow errors for very deep trees.  The memory use is controlled by the queue size which is related to the width of the tree and is more predictible.
* **Space Efficiency (in some cases):** For wide trees, the queue might hold many nodes. In some of these cases where the recursive approach may not overflow, the iterative solution might still be more space efficient, especially for balanced or wide trees, as the queue's maximum size is proportional to the maximum width of the tree, while the call stack depth in the recursive approach is proportional to the maximum depth of the tree. In balanced binary trees, this difference is not as great as when compared to heavily skewed trees (where the recursive approach will use memory proportional to the number of nodes, while iterative will only use memory proportional to the relatively small width).  For trees that are very deep but not very wide, the iterative method may use similar memory as the recursive.

The changes made ensure that the iterative version mirrors the functionality of the recursive one while providing better memory management for cases where deep trees might cause stack overflows.

## Is it possible to use a stack to invert a tree?

Yes, it's absolutely possible to use a stack to invert a binary tree iteratively. Here's how you can implement it in Go:

```go
package binarytreeinvert

// TreeNode represents a node in a binary tree
type TreeNode struct {
 Val   int
 Left  *TreeNode
 Right *TreeNode
}

// InvertTree inverts a binary tree using a stack.
func InvertTree(root *TreeNode) *TreeNode {
 if root == nil {
  return nil
 }

 stack := []*TreeNode{root}

 for len(stack) > 0 {
  current := stack[len(stack)-1]
  stack = stack[:len(stack)-1] // Pop

  // Swap left and right children
  current.Left, current.Right = current.Right, current.Left

  // Push non-nil children (order doesn't strictly matter here but follows pre-order)
  if current.Left != nil {
   stack = append(stack, current.Left) 
  }
  if current.Right != nil {
   stack = append(stack, current.Right)
  }
 }

 return root
}

```

Here's how this stack-based approach works:

1. **Initialization:** Initialize a stack with the root node.
2. **Iteration:** While the stack is not empty:
   * Pop the top node from the stack.
   * Swap its left and right children.
   * Push the left child (if not nil) onto the stack.
   * Push the right child (if not nil) onto the stack.

The stack-based approach, much like the queue-based iterative approach discussed earlier, also prevents the stack overflow issues that deep recursion can cause, and provides similar memory characteristics to the queue-based iterative version.  The conceptual order of operations will change based on using a queue (level order) vs. stack (pre-order if left is pushed before right, or post-order if left pushed after right).  The final outcome is the same inverted tree regardless of whether stack or queue is used.
