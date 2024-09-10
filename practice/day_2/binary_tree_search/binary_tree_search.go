package binarytreesearch

// TreeNode represents a node in a binary search tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Insert inserts a new value into the binary search tree
func Insert(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	var node *TreeNode = root

	for node != nil {
		parent = node

		if node.Val == val {
			return root
		}

		if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	newNode := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}

	if parent == nil {
		return newNode
	}

	if val < parent.Val {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	return root
}

// Search searches for a value in the binary search tree
func Search(root *TreeNode, val int) bool {
	node := root

	for node != nil {
		if node.Val == val {
			return true
		} else if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return false
}

// Delete removes a value from the binary search tree
func Delete(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	var node *TreeNode = root

	for node != nil && node.Val != val {
		parent = node
		if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	if node == nil {
		return root
	}

	if node.Left != nil && node.Right != nil {
		successor := node.Right
		successorParent := node

		for successor.Left != nil {
			successorParent = successor
			successor = successor.Left
		}

		node.Val = successor.Val
		node = successor
		parent = successorParent
	}

	var childNode *TreeNode

	if node.Left != nil {
		childNode = node.Left
	} else {
		childNode = node.Right
	}

	if parent == nil {
		return nil
	} else if node == parent.Left {
		parent.Left = childNode
	} else {
		parent.Right = childNode
	}

	return root
}

// findMin finds the node with the minimum value in a subtree
func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// InorderTraversal performs an inorder traversal of the binary search tree
func InorderTraversal(root *TreeNode) []int {
	var result []int

	inorderTraversalHelper(root, &result)

	return result
}

func inorderTraversalHelper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inorderTraversalHelper(root.Left, result)
	*result = append(*result, root.Val)
	inorderTraversalHelper(root.Right, result)
}
