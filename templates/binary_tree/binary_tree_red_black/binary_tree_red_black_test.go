package redblacktree

import (
    "reflect"
    "testing"
)

func TestRedBlackTree(t *testing.T) {
    tree := NewRedBlackTree()

    // Test insertion
    keys := []int{7, 3, 18, 10, 22, 8, 11, 26, 2, 6, 13}
    for _, key := range keys {
        tree.Insert(key)
    }

    // Test inorder traversal
    expected := []int{2, 3, 6, 7, 8, 10, 11, 13, 18, 22, 26}
    result := tree.InorderTraversal()
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("InorderTraversal() = %v, want %v", result, expected)
    }

    // Test search
    testCases := []struct {
        key      int
        expected bool
    }{
        {7, true},
        {3, true},
        {26, true},
        {15, false},
        {1, false},
    }

    for _, tc := range testCases {
        result := tree.Search(tc.key)
        if result != tc.expected {
            t.Errorf("Search(%d) = %v, want %v", tc.key, result, tc.expected)
        }
    }

    // Test Red-Black tree properties
    if !isValidRedBlackTree(tree) {
        t.Errorf("Tree does not satisfy Red-Black tree properties")
    }
}

func isValidRedBlackTree(tree *RedBlackTree) bool {
    if tree.Root.Color != BLACK {
        return false
    }
    return validateNode(tree.Root, tree.NIL, BLACK, 0) != -1
}

func validateNode(node *Node, nil_node *Node, parentColor Color, blackHeight int) int {
    if node == nil_node {
        return blackHeight + 1
    }

    if node.Color == RED && parentColor == RED {
        return -1
    }

    if node.Color == BLACK {
        blackHeight++
    }

    leftBlackHeight := validateNode(node.Left, nil_node, node.Color, blackHeight)
    rightBlackHeight := validateNode(node.Right, nil_node, node.Color, blackHeight)

    if leftBlackHeight == -1 || rightBlackHeight == -1 || leftBlackHeight != rightBlackHeight {
        return -1
    }

    return leftBlackHeight
}
