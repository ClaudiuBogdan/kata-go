package binarytreesearch

// TreeNode represents a node in a binary search tree
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Insert inserts a new value into the binary search tree
func Insert(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }

    if val < root.Val {
        root.Left = Insert(root.Left, val)
    } else if val > root.Val {
        root.Right = Insert(root.Right, val)
    }

    return root
}

// Search searches for a value in the binary search tree
func Search(root *TreeNode, val int) bool {
    if root == nil {
        return false
    }

    if val == root.Val {
        return true
    } else if val < root.Val {
        return Search(root.Left, val)
    } else {
        return Search(root.Right, val)
    }
}

// Delete removes a value from the binary search tree
func Delete(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }

    if val < root.Val {
        root.Left = Delete(root.Left, val)
    } else if val > root.Val {
        root.Right = Delete(root.Right, val)
    } else {
        // Node to delete found
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        }

        // Node with two children: Get the inorder successor (smallest
        // in the right subtree)
        minNode := findMin(root.Right)
        root.Val = minNode.Val
        root.Right = Delete(root.Right, root.Val)
    }

    return root
}

// findMin finds the node with the minimum value in a subtree
func findMin(node *TreeNode) *TreeNode {
    current := node
    for current.Left != nil {
        current = current.Left
    }
    return current
}

// InorderTraversal performs an inorder traversal of the binary search tree
func InorderTraversal(root *TreeNode) []int {
    var result []int
    inorderTraversalHelper(root, &result)
    return result
}

func inorderTraversalHelper(node *TreeNode, result *[]int) {
    if node != nil {
        inorderTraversalHelper(node.Left, result)
        *result = append(*result, node.Val)
        inorderTraversalHelper(node.Right, result)
    }
}
