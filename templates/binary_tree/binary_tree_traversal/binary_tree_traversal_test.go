package binarytreetraversal

import (
	"reflect"
	"testing"
)

func TestBinaryTreeTraversal(t *testing.T) {
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
        traversalFunc interface{}
        expected interface{}
    }{
        {
            name:     "Inorder Traversal",
            tree:     root,
            traversalFunc: InorderTraversal,
            expected: []int{4, 2, 5, 1, 3, 6},
        },
        {
            name:     "Preorder Traversal",
            tree:     root,
            traversalFunc: PreorderTraversal,
            expected: []int{1, 2, 4, 5, 3, 6},
        },
        {
            name:     "Postorder Traversal",
            tree:     root,
            traversalFunc: PostorderTraversal,
            expected: []int{4, 5, 2, 6, 3, 1},
        },
        {
            name:     "Level-order Traversal",
            tree:     root,
            traversalFunc: LevelOrderTraversal,
            expected: [][]int{{1}, {2, 3}, {4, 5, 6}},
        },
        {
            name:     "Empty tree",
            tree:     nil,
            traversalFunc: InorderTraversal,
            expected: []int(nil),
        },
        {
            name:     "Tree with only root",
            tree:     &TreeNode{Val: 1},
            traversalFunc: PreorderTraversal,
            expected: []int{1},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var result interface{}
            switch f := tt.traversalFunc.(type) {
            case func(*TreeNode) []int:
                result = f(tt.tree)
            case func(*TreeNode) [][]int:
                result = f(tt.tree)
            }
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("%s() = %v, want %v", tt.name, result, tt.expected)
            }
        })
    }
}
