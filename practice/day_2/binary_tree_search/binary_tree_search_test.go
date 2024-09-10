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
}
