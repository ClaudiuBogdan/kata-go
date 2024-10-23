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
    
    queue := []*TreeNode{root}
    result := make([]int, 0)
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        result = append(result, current.Val)

        if current.Left != nil {
            queue = append(queue, current.Left)
        }

        if current.Right != nil {
            queue = append(queue, current.Right)
        }
    }

    return result
}