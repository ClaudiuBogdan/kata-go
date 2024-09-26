package binarytreelca

import (
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
    // Create a sample binary tree
    //       3
    //      / \
    //     5   1
    //    / \ / \
    //   6  2 0  8
    //     / \
    //    7   4
    root := &TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 5}
    root.Right = &TreeNode{Val: 1}
    root.Left.Left = &TreeNode{Val: 6}
    root.Left.Right = &TreeNode{Val: 2}
    root.Right.Left = &TreeNode{Val: 0}
    root.Right.Right = &TreeNode{Val: 8}
    root.Left.Right.Left = &TreeNode{Val: 7}
    root.Left.Right.Right = &TreeNode{Val: 4}

    onlyRoot := &TreeNode{Val: 1}

    tests := []struct {
        name     string
        root     *TreeNode
        p        *TreeNode
        q        *TreeNode
        expected *TreeNode
    }{
        {
            name:     "LCA of 5 and 1 is 3",
            root:     root,
            p:        root.Left,
            q:        root.Right,
            expected: root,
        },
        {
            name:     "LCA of 5 and 4 is 5",
            root:     root,
            p:        root.Left,
            q:        root.Left.Right.Right,
            expected: root.Left,
        },
        {
            name:     "LCA of 6 and 4 is 5",
            root:     root,
            p:        root.Left.Left,
            q:        root.Left.Right.Right,
            expected: root.Left,
        },
        {
            name:     "LCA of 3 and 4 is 3",
            root:     root,
            p:        root,
            q:        root.Left.Right.Right,
            expected: root,
        },
        {
            name:     "LCA in a tree with only root",
            root:     onlyRoot,
            p:        onlyRoot,
            q:        onlyRoot,
            expected: onlyRoot,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := LowestCommonAncestor(tt.root, tt.p, tt.q)
            if result != tt.expected {
                t.Errorf("LowestCommonAncestor() = %v, want %v", result.Val, tt.expected.Val)
            }
        })
    }
}