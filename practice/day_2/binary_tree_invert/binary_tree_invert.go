package binarytreeinvert

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InvertTree inverts a binary tree
func InvertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}
