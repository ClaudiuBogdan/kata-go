package binarytreebfs

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name     string
		tree     *TreeNode
		expected []int
	}{
		{
			name: "Normal binary tree",
			tree: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "Empty tree",
			tree:     nil,
			expected: []int{},
		},
		{
			name:     "Tree with only root",
			tree:     &TreeNode{Val: 1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BFS(tt.tree)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BFS() = %v, want %v", result, tt.expected)
			}
		})
	}
}
