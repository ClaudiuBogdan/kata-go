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

    // Swap the left and right children
    root.Left, root.Right = root.Right, root.Left

    // Recursively invert the left and right subtrees
    InvertTree(root.Left)
    InvertTree(root.Right)

    return root
}
