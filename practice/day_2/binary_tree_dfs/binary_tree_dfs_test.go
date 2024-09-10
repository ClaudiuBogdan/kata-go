package binarytreedfs

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
    // Create a sample binary tree
    //       1
    //      / \
    //     2   3
    //    / \   \
    //   4   5   6
    root := &TreeNode{
        Val: 1,
        Left: &TreeNode{
            Val:   2,
            Left:  &TreeNode{Val: 4},
            Right: &TreeNode{Val: 5},
        },
        Right: &TreeNode{
            Val:   3,
            Right: &TreeNode{Val: 6},
        },
    }

    tests := []struct {
        name     string
        tree     *TreeNode
        dfsFunc  func(*TreeNode) []int
        expected []int
    }{
        {
            name:     "Preorder DFS",
            tree:     root,
            dfsFunc:  PreorderDFS,
            expected: []int{1, 2, 4, 5, 3, 6},
        },
        {
            name:     "Inorder DFS",
            tree:     root,
            dfsFunc:  InorderDFS,
            expected: []int{4, 2, 5, 1, 3, 6},
        },
        {
            name:     "Postorder DFS",
            tree:     root,
            dfsFunc:  PostorderDFS,
            expected: []int{4, 5, 2, 6, 3, 1},
        },
        {
            name:     "Empty tree",
            tree:     nil,
            dfsFunc:  PreorderDFS,
            expected: []int(nil),
        },
        {
            name:     "Tree with only root",
            tree:     &TreeNode{Val: 1},
            dfsFunc:  InorderDFS,
            expected: []int{1},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := tt.dfsFunc(tt.tree)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("%s() = %v, want %v", tt.name, result, tt.expected)
            }
        })
    }
}
