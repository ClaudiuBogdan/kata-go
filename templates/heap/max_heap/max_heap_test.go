package maxheap

import (
    "testing"
)

func TestMaxHeap(t *testing.T) {
    t.Run("Insert and Extract", func(t *testing.T) {
        heap := NewMaxHeap()
        values := []int{4, 10, 3, 5, 1}
        expected := []int{10, 5, 4, 3, 1}

        for _, val := range values {
            heap.Insert(val)
        }

        for _, exp := range expected {
            val, ok := heap.Extract()
            if !ok {
                t.Errorf("Expected to extract a value, but got false")
            }
            if val != exp {
                t.Errorf("Expected %d, but got %d", exp, val)
            }
        }

        _, ok := heap.Extract()
        if ok {
            t.Errorf("Expected false when extracting from empty heap")
        }
    })

    t.Run("Peek", func(t *testing.T) {
        heap := NewMaxHeap()
        heap.Insert(5)
        heap.Insert(10)
        heap.Insert(3)

        val, ok := heap.Peek()
        if !ok {
            t.Errorf("Expected to peek a value, but got false")
        }
        if val != 10 {
            t.Errorf("Expected peek to return 10, but got %d", val)
        }

        heap.Extract()
        val, ok = heap.Peek()
        if !ok {
            t.Errorf("Expected to peek a value, but got false")
        }
        if val != 5 {
            t.Errorf("Expected peek to return 5, but got %d", val)
        }
    })

    t.Run("Size", func(t *testing.T) {
        heap := NewMaxHeap()
        if heap.Size() != 0 {
            t.Errorf("Expected empty heap to have size 0, but got %d", heap.Size())
        }

        heap.Insert(5)
        heap.Insert(10)
        if heap.Size() != 2 {
            t.Errorf("Expected heap to have size 2, but got %d", heap.Size())
        }

        heap.Extract()
        if heap.Size() != 1 {
            t.Errorf("Expected heap to have size 1, but got %d", heap.Size())
        }
    })
}
