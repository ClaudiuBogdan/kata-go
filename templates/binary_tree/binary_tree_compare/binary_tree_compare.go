package binarytreecompare

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// CompareTrees compares two binary trees
func CompareTrees(p *TreeNode, q *TreeNode) bool {
	
}

// NewNode creates a new TreeNode with the given value
func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}
