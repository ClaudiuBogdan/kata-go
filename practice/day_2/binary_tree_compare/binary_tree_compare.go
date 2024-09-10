package binarytreecompare

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CompareTrees compares two binary trees
func CompareTrees(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && CompareTrees(p.Left, q.Left) && CompareTrees(p.Right, q.Right)
}

// NewNode creates a new TreeNode with the given value
func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
