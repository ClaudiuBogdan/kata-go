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
	preorderDfsHelper(root, &result)
	return result
}

// InorderDFS performs an inorder DFS traversal of the binary tree
func InorderDFS(root *TreeNode) []int {
	var result []int
	inorderDfsHelper(root, &result)
	return result
}

// PostorderDFS performs a postorder DFS traversal of the binary tree
func PostorderDFS(root *TreeNode) []int {
	var result []int
	postorderDfsHelper(root, &result)
	return result
}

func preorderDfsHelper(root *TreeNode, result *[]int){
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	preorderDfsHelper(root.Left, result)
	preorderDfsHelper(root.Right, result)
}

func inorderDfsHelper(root *TreeNode, result *[]int){
	if root == nil {
		return
	}

	inorderDfsHelper(root.Left, result)
	*result = append(*result, root.Val)
	inorderDfsHelper(root.Right, result)
}

func postorderDfsHelper(root *TreeNode, result *[]int){
	if root == nil {
		return
	}

	postorderDfsHelper(root.Left, result)
	postorderDfsHelper(root.Right, result)
	*result = append(*result, root.Val)
}