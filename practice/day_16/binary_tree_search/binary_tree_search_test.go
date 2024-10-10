package binarytreesearch

import (
	"reflect"
	"testing"
)

func TestBinaryTreeSearchOperations(t *testing.T) {
    var root *TreeNode

    // Test Insert
    values := []int{5, 3, 7, 1, 4, 6, 8}
    for _, val := range values {
        root = Insert(root, val)
    }

    // Test Inorder Traversal
    expected := []int{1, 3, 4, 5, 6, 7, 8}
    result := InorderTraversal(root)
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("InorderTraversal() = %v, want %v", result, expected)
    }

    // Test Search
    testCases := []struct {
        val      int
        expected bool
    }{
        {5, true},
        {3, true},
        {8, true},
        {9, false},
        {0, false},
    }

    for _, tc := range testCases {
        result := Search(root, tc.val)
        if result != tc.expected {
            t.Errorf("Search(%d) = %v, want %v", tc.val, result, tc.expected)
        }
    }

    // Test Delete
    deleteTests := []struct {
        val      int
        expected []int
    }{
        {4, []int{1, 3, 5, 6, 7, 8}},
        {5, []int{1, 3, 6, 7, 8}},
        {7, []int{1, 3, 6, 8}},
        {1, []int{3, 6, 8}},
        {9, []int{3, 6, 8}}, // Deleting non-existent value
    }

    for _, dt := range deleteTests {
        root = Delete(root, dt.val)
        result := InorderTraversal(root)
        if !reflect.DeepEqual(result, dt.expected) {
            t.Errorf("After Delete(%d), InorderTraversal() = %v, want %v", dt.val, result, dt.expected)
        }
    }


	// Test Insert with duplicate values
	root = Insert(root, 6)
	result = InorderTraversal(root)
	if !reflect.DeepEqual(result, []int{3, 6, 8}) {
		t.Errorf("After Insert(6), InorderTraversal() = %v, want %v", result, []int{3, 6, 8})
	}

	// Test Delete root node
	root = Delete(root, 6)
	result = InorderTraversal(root)
	if !reflect.DeepEqual(result, []int{3, 8}) {
		t.Errorf("After Delete(6), InorderTraversal() = %v, want %v", result, []int{3, 8})
	}

	// Test Delete last node
	root = Delete(root, 8)
	root = Delete(root, 3)
	if root != nil {
		t.Errorf("After deleting all nodes, root should be nil, but got %v", root)
	}

	// Test operations on empty tree
	if Search(root, 1) {
		t.Errorf("Search(1) on empty tree should return false")
	}

	root = Insert(root, 1)
	if !Search(root, 1) {
		t.Errorf("Search(1) after inserting 1 should return true")
	}

	root = Delete(root, 1)
	if root != nil {
		t.Errorf("After deleting last node, root should be nil, but got %v", root)
	}
}
