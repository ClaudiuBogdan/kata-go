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
	var node = root

	for node != nil {
		parent = node

		if node.Val == val {
			return root
		} else if val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	newNode := &TreeNode{
		Val: val,
	}

	if parent == nil {
		return newNode
	} else if val < parent.Val {
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
	var node = root

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

	var child *TreeNode

	if node.Left != nil {
		child = node.Left
	} else {
		child = node.Right
	}

	if parent == nil {
		return child
	} else if node.Val < parent.Val {
		parent.Left = child
	} else {
		parent.Right = child
	}
	return root
}

// InorderTraversal performs an inorder traversal of the binary search tree
func InorderTraversal(root *TreeNode) []int {
	path := []int(nil)
	inorderTraversalHelper(root, &path)
	return path
}

func inorderTraversalHelper(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}

	inorderTraversalHelper(root.Left, path)
	*path = append(*path, root.Val)
	inorderTraversalHelper(root.Right, path)
}
