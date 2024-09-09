package avltree

import (
    "testing"
)

func TestAVLTree(t *testing.T) {
    tree := NewAVLTree()

    // Test insertion
    keys := []int{10, 20, 30, 40, 50, 25}
    for _, key := range keys {
        tree.Insert(key)
    }

    // Test search
    for _, key := range keys {
        if !tree.Search(key) {
            t.Errorf("Key %d not found in the tree", key)
        }
    }

    if tree.Search(100) {
        t.Errorf("Key 100 should not be in the tree")
    }

    // Test deletion
    tree.Delete(30)
    if tree.Search(30) {
        t.Errorf("Key 30 should have been deleted")
    }

    // Test balance after insertion and deletion
    if getBalance(tree.Root) < -1 || getBalance(tree.Root) > 1 {
        t.Errorf("Tree is not balanced after operations")
    }
}

func TestEmptyTree(t *testing.T) {
    tree := NewAVLTree()

    if tree.Search(10) {
        t.Errorf("Empty tree should not contain any elements")
    }

    tree.Delete(10) // Should not panic
}

func TestDuplicateInsertion(t *testing.T) {
    tree := NewAVLTree()

    tree.Insert(10)
    tree.Insert(10)

    if getHeight(tree.Root) != 1 {
        t.Errorf("Duplicate insertion should not change the tree")
    }
}

func TestRotations(t *testing.T) {
    // Test left-left case
    tree := NewAVLTree()
    tree.Insert(30)
    tree.Insert(20)
    tree.Insert(10)

    if tree.Root.Key != 20 {
        t.Errorf("Left-left rotation failed")
    }

    // Test right-right case
    tree = NewAVLTree()
    tree.Insert(10)
    tree.Insert(20)
    tree.Insert(30)

    if tree.Root.Key != 20 {
        t.Errorf("Right-right rotation failed")
    }

    // Test left-right case
    tree = NewAVLTree()
    tree.Insert(30)
    tree.Insert(10)
    tree.Insert(20)

    if tree.Root.Key != 20 {
        t.Errorf("Left-right rotation failed")
    }

    // Test right-left case
    tree = NewAVLTree()
    tree.Insert(10)
    tree.Insert(30)
    tree.Insert(20)

    if tree.Root.Key != 20 {
        t.Errorf("Right-left rotation failed")
    }
}
