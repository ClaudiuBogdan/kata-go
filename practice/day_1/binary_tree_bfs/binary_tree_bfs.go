package binarytreebfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BFS(root *TreeNode) []int {
	if root == nil {
		return []int(nil)
	}

	result := []int{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}
