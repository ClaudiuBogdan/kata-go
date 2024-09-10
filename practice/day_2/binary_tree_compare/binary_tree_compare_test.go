package binarytreecompare

import (
	"testing"
)

func TestCompareTrees(t *testing.T) {
	tests := []struct {
		name     string
		tree1    *TreeNode
		tree2    *TreeNode
		expected bool
	}{
		{
			name:     "Identical trees",
			tree1:    buildTree(1, 2, 3, 4, 5),
			tree2:    buildTree(1, 2, 3, 4, 5),
			expected: true,
		},
		{
			name:     "Different values",
			tree1:    buildTree(1, 2, 3),
			tree2:    buildTree(1, 2, 4),
			expected: false,
		},
		{
			name:     "Different structures",
			tree1:    buildTree(1, 2, 3, 4),
			tree2:    buildTree(1, 2, 3, nil, 4),
			expected: false,
		},
		{
			name:     "One tree nil",
			tree1:    buildTree(1, 2, 3),
			tree2:    nil,
			expected: false,
		},
		{
			name:     "Both trees nil",
			tree1:    nil,
			tree2:    nil,
			expected: true,
		},
		{
			name:     "Single node trees",
			tree1:    buildTree(1),
			tree2:    buildTree(1),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareTrees(tt.tree1, tt.tree2)
			if result != tt.expected {
				t.Errorf("CompareTrees() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Helper function to build a tree from a slice of values
func buildTree(values ...interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := NewNode(values[0].(int))
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		if i < len(values) && values[i] != nil {
			node.Left = NewNode(values[i].(int))
			queue = append(queue, node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = NewNode(values[i].(int))
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
