package binarytreeinvert

import (
    "reflect"
    "testing"
)

func TestInvertTree(t *testing.T) {
    tests := []struct {
        name     string
        input    *TreeNode
        expected *TreeNode
    }{
        {
            name: "Normal binary tree",
            input: &TreeNode{
                Val: 4,
                Left: &TreeNode{
                    Val:   2,
                    Left:  &TreeNode{Val: 1},
                    Right: &TreeNode{Val: 3},
                },
                Right: &TreeNode{
                    Val:   7,
                    Left:  &TreeNode{Val: 6},
                    Right: &TreeNode{Val: 9},
                },
            },
            expected: &TreeNode{
                Val: 4,
                Left: &TreeNode{
                    Val:   7,
                    Left:  &TreeNode{Val: 9},
                    Right: &TreeNode{Val: 6},
                },
                Right: &TreeNode{
                    Val:   2,
                    Left:  &TreeNode{Val: 3},
                    Right: &TreeNode{Val: 1},
                },
            },
        },
        {
            name:     "Empty tree",
            input:    nil,
            expected: nil,
        },
        {
            name:     "Tree with only root",
            input:    &TreeNode{Val: 1},
            expected: &TreeNode{Val: 1},
        },
        {
            name: "Unbalanced tree",
            input: &TreeNode{
                Val: 1,
                Left: &TreeNode{
                    Val: 2,
                    Left: &TreeNode{
                        Val: 3,
                    },
                },
            },
            expected: &TreeNode{
                Val: 1,
                Right: &TreeNode{
                    Val: 2,
                    Right: &TreeNode{
                        Val: 3,
                    },
                },
            },
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := InvertTree(tt.input)
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("InvertTree() = %v, want %v", result, tt.expected)
            }
        })
    }
}

// Helper function to compare trees
func compareTrees(t *testing.T, got, want *TreeNode) bool {
    if got == nil && want == nil {
        return true
    }
    if got == nil || want == nil {
        return false
    }
    if got.Val != want.Val {
        return false
    }
    return compareTrees(t, got.Left, want.Left) && compareTrees(t, got.Right, want.Right)
}
