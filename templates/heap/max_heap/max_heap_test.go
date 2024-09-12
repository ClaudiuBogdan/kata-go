package maxheap

import (
	"testing"
)

func TestMaxHeap(t *testing.T) {
	t.Run("NewMaxHeap", func(t *testing.T) {
		heap := New[int]()
		if !heap.IsEmpty() {
			t.Errorf("New heap should be empty")
		}
		if heap.Size() != 0 {
			t.Errorf("New heap should have size 0, got %d", heap.Size())
		}
	})

	t.Run("Insert and Peek", func(t *testing.T) {
		heap := New[int]()
		heap.Insert(5)
		heap.Insert(3)
		heap.Insert(7)

		max, err := heap.Peek()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if max != 7 {
			t.Errorf("Expected max to be 7, got %d", max)
		}
		if heap.Size() != 3 {
			t.Errorf("Expected size to be 3, got %d", heap.Size())
		}
	})

	t.Run("Extract", func(t *testing.T) {
		heap := New[int]()
		heap.Insert(5)
		heap.Insert(3)
		heap.Insert(7)
		heap.Insert(1)

		max, err := heap.Extract()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if max != 7 {
			t.Errorf("Expected max to be 7, got %d", max)
		}
		if heap.Size() != 3 {
			t.Errorf("Expected size to be 3, got %d", heap.Size())
		}

		max, err = heap.Peek()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if max != 5 {
			t.Errorf("Expected new max to be 5, got %d", max)
		}
	})

	t.Run("EmptyHeapOperations", func(t *testing.T) {
		heap := New[int]()

		_, err := heap.Peek()
		if err == nil {
			t.Errorf("Expected error on Peek() for empty heap")
		}

		_, err = heap.Extract()
		if err == nil {
			t.Errorf("Expected error on Extract() for empty heap")
		}
	})

	t.Run("LargeNumberOfInserts", func(t *testing.T) {
		heap := New[int]()
		for i := 0; i < 1000; i++ {
			heap.Insert(i)
		}

		if heap.Size() != 1000 {
			t.Errorf("Expected size to be 1000, got %d", heap.Size())
		}

		max, _ := heap.Peek()
		if max != 999 {
			t.Errorf("Expected max to be 999, got %d", max)
		}
	})
}