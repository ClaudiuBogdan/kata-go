# Iterative implementation of binary tree compare

## My own implementation of binary tree compare

```go
package binarytreecompare

// TreeNode represents a node in a binary tree
type TreeNode struct {
 Val   int
 Left  *TreeNode
 Right *TreeNode
}

// CompareTrees compares two binary trees
func CompareTrees(p *TreeNode, q *TreeNode) bool {

 compare := func(a, b *TreeNode) bool {
  if a == nil && b == nil {
   return true
  }

  if a != nil && b == nil {
   return false
  }

  if a == nil && b != nil {
   return false
  }
  return a.Val == b.Val
 }

 sameRoot := compare(p, q)
 if !sameRoot {
  return false
 }

 if p == nil {
  return true
 }

 pQueue := []*TreeNode{p}
 qQueue := []*TreeNode{q}

 for len(pQueue) > 0 {

  pNode := pQueue[0]
  qNode := qQueue[0]

  pQueue = pQueue[1:]
  qQueue = qQueue[1:]

  haveSameChildren := compare(pNode.Left, qNode.Left) && compare(pNode.Right, qNode.Right)

  if !haveSameChildren {
   return false
  }

  if pNode.Left != nil {
   pQueue = append(pQueue, pNode.Left)
  }

  if pNode.Right != nil {
   pQueue = append(pQueue, pNode.Right)
  }

  if qNode.Left != nil {
   qQueue = append(qQueue, qNode.Left)
  }

  if qNode.Right != nil {
   qQueue = append(qQueue, q.Right)
  }
 }
 return true
}

// NewNode creates a new TreeNode with the given value
func NewNode(val int) *TreeNode {
 return &TreeNode{Val: val}
}
```

## Claude Sonnet 3.5

```go
package binarytreecompare

// TreeNode represents a node in a binary tree
type TreeNode struct {
 Val   int
 Left  *TreeNode
 Right *TreeNode
}

// CompareTrees compares two binary trees iteratively
func CompareTrees(p *TreeNode, q *TreeNode) bool {
 stack := [][2]*TreeNode{{p, q}}

 for len(stack) > 0 {
  nodes := stack[len(stack)-1]
  stack = stack[:len(stack)-1]

  p, q := nodes[0], nodes[1]

  if p == nil && q == nil {
   continue
  }

  if p == nil || q == nil || p.Val != q.Val {
   return false
  }

  stack = append(stack, [2]*TreeNode{p.Right, q.Right})
  stack = append(stack, [2]*TreeNode{p.Left, q.Left})
 }

 return true
}

// NewNode creates a new TreeNode with the given value
func NewNode(val int) *TreeNode {
 return &TreeNode{Val: val}
}
```

## Gemini pro 1.5

```go
package binarytreecompare

// TreeNode represents a node in a binary tree
type TreeNode struct {
 Val   int
 Left  *TreeNode
 Right *TreeNode
}

// CompareTrees compares two binary trees iteratively.
func CompareTrees(p *TreeNode, q *TreeNode) bool {
 var stack []*TreeNode
 stack = append(stack, p, q) // Push both root nodes onto the stack

 for len(stack) > 0 {
  q = stack[len(stack)-1]
  stack = stack[:len(stack)-1] // Pop q
  p = stack[len(stack)-1]
  stack = stack[:len(stack)-1] // Pop p

  if p == nil && q == nil {
   continue
  }
  if p == nil || q == nil {
   return false
  }
  if p.Val != q.Val {
   return false
  }

  stack = append(stack, p.Left, q.Left)  // Push left children
  stack = append(stack, p.Right, q.Right) // Push right children (order matters for mirroring)
 }

 return true
}

// NewNode creates a new TreeNode with the given value
func NewNode(val int) *TreeNode {
 return &TreeNode{Val: val}
}

```

**Improvements:**

* **Iterative Approach:** The original code used recursion, which can lead to stack overflow errors for very deep trees. This revised version uses an explicit stack, making it robust against deep trees and generally more efficient in terms of memory usage for large trees.  Iterative solutions often have better performance characteristics in Go due to the overhead associated with function calls in recursive approaches.

* **Reduced Redundancy:** The original code had redundant `nil` checks. This version streamlines the checks, improving readability and slightly reducing the number of comparisons.

* **Order of Stack Operations:** The order in which nodes are pushed and popped from the stack is crucial for correct comparison. This version ensures the correct mirroring of left and right subtrees during the comparison.

**Explanation of the Iterative Algorithm:**

1. **Initialization:** A stack is used to store pairs of nodes (one from each tree) that need to be compared.  Initially, the root nodes of both trees are pushed onto the stack.

2. **Iteration:** The loop continues as long as the stack is not empty.

3. **Pop and Compare:** In each iteration, a pair of nodes (p, q) is popped from the stack.  Null checks are performed. If both are `nil`, the subtree comparison is successful, and the loop continues. If only one is `nil`, the trees are different, and `false` is returned. If both are non-`nil`, their values are compared. If the values are different, `false` is returned.

4. **Push Children:** If the values are equal, the left children and then the right children of p and q are pushed onto the stack.  The order is important to maintain the correct comparison structure.

5. **Return True:** If the loop completes without returning `false`, it means all corresponding nodes have been compared and found to be equal, so the function returns `true`.

This iterative approach provides a more robust and efficient solution for comparing binary trees, especially for large or deeply nested trees. It avoids the potential stack overflow issues of recursion and generally has better performance characteristics in Go.

## Stack vs Queue for Iterative Binary Tree Comparison

### Stack (Depth-First Search)

#### Advantages

1. Memory efficiency: In a balanced tree, uses O(log n) space.
2. Mimics recursive call stack: Easier to convert recursive solutions to iterative.
3. Depth-first nature: Can find differences faster in some cases.

#### Disadvantages

1. May use O(n) space for skewed trees.
2. Doesn't preserve level-order traversal.

### Queue (Breadth-First Search)

#### Advantages

1. Consistent O(w) space complexity, where w is the maximum width of the tree.
2. Preserves level-order traversal.
3. Can be more intuitive for level-by-level comparison.

#### Disadvantages

1. May use more memory for wide, shallow trees.
2. Slightly more complex to implement for tree comparison.

### Conclusion

For most binary tree comparison scenarios, a stack-based approach (DFS) is preferable:

* It's generally more memory-efficient.
* It's easier to implement, especially when converting from recursive solutions.
* It can terminate earlier in many cases of non-identical trees.

However, queue-based approaches (BFS) can be beneficial in specific scenarios:

* When comparing very wide, shallow trees.
* When level-order traversal is important for the comparison.
* When you need to find the first differing level quickly.
