package binarytreelca

// TreeNode represents a node in a binary tree
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// LowestCommonAncestor finds the lowest common ancestor of two nodes in a binary tree
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // Base case: if root is nil or equal to either p or q, return root
    if root == nil || root == p || root == q {
        return root
    }

    // Recursively search in left and right subtrees
    left := LowestCommonAncestor(root.Left, p, q)
    right := LowestCommonAncestor(root.Right, p, q)

    // If both left and right are non-nil, root is the LCA
    if left != nil && right != nil {
        return root
    }

    // If one side is nil, return the non-nil side
    if left != nil {
        return left
    }
    return right
}