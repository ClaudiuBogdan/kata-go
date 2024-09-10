package binarytreelca

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LowestCommonAncestor finds the lowest common ancestor of two nodes in a binary tree
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	leftAncestor := LowestCommonAncestor(root.Left, p, q)
	rightAncestor := LowestCommonAncestor(root.Right, p, q)

	if leftAncestor != nil && rightAncestor != nil {
		return root
	}

	if leftAncestor != nil {
		return leftAncestor
	}

	return rightAncestor
}
