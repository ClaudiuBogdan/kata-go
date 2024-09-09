package binarytreetraversal

// TreeNode represents a node in a binary tree
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// InorderTraversal performs an inorder traversal of the binary tree
func InorderTraversal(root *TreeNode) []int {
    var result []int
    inorderHelper(root, &result)
    return result
}

func inorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    inorderHelper(node.Left, result)
    *result = append(*result, node.Val)
    inorderHelper(node.Right, result)
}

// PreorderTraversal performs a preorder traversal of the binary tree
func PreorderTraversal(root *TreeNode) []int {
    var result []int
    preorderHelper(root, &result)
    return result
}

func preorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    *result = append(*result, node.Val)
    preorderHelper(node.Left, result)
    preorderHelper(node.Right, result)
}

// PostorderTraversal performs a postorder traversal of the binary tree
func PostorderTraversal(root *TreeNode) []int {
    var result []int
    postorderHelper(root, &result)
    return result
}

func postorderHelper(node *TreeNode, result *[]int) {
    if node == nil {
        return
    }
    postorderHelper(node.Left, result)
    postorderHelper(node.Right, result)
    *result = append(*result, node.Val)
}

// LevelOrderTraversal performs a level-order traversal of the binary tree
func LevelOrderTraversal(root *TreeNode) [][]int {
    var result [][]int
    if root == nil {
        return result
    }

    queue := []*TreeNode{root}
    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := make([]int, 0, levelSize)

        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            currentLevel = append(currentLevel, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, currentLevel)
    }

    return result
}
