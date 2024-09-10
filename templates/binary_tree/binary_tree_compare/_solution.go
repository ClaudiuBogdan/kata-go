package binarytreecompare

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CompareTrees compares two binary trees
func CompareTrees(p *TreeNode, q *TreeNode) bool {
	// If both nodes are nil, trees are identical up to this point
	if p == nil && q == nil {
		return true
	}
	
	// If one node is nil and the other isn't, trees are not identical
	if p == nil || q == nil {
		return false
	}
	
	// Check if current nodes have the same value
	if p.Val != q.Val {
		return false
	}
	
	// Recursively compare left and right subtrees
	return CompareTrees(p.Left, q.Left) && CompareTrees(p.Right, q.Right)
}

// NewNode creates a new TreeNode with the given value
func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
