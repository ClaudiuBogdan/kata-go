package binarytreedfs

// TreeNode represents a node in a binary tree
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// PreorderDFS performs a preorder DFS traversal of the binary tree
func PreorderDFS(root *TreeNode) []int {
    var result []int
    preorderDFSHelper(root, &result)
    return result
}

func preorderDFSHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    *result = append(*result, node.Val)
    preorderDFSHelper(node.Left, result)
    preorderDFSHelper(node.Right, result)
}

// InorderDFS performs an inorder DFS traversal of the binary tree
func InorderDFS(root *TreeNode) []int {
    var result []int
    inorderDFSHelper(root, &result)
    return result
}

func inorderDFSHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    inorderDFSHelper(node.Left, result)
    *result = append(*result, node.Val)
    inorderDFSHelper(node.Right, result)
}

// PostorderDFS performs a postorder DFS traversal of the binary tree
func PostorderDFS(root *TreeNode) []int {
    var result []int
    postorderDFSHelper(root, &result)
    return result
}

func postorderDFSHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    postorderDFSHelper(node.Left, result)
    postorderDFSHelper(node.Right, result)
    *result = append(*result, node.Val)
}
