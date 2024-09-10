package dynamicarray

import (
    "testing"
)

func TestDynamicArray(t *testing.T) {
    t.Run("Basic Operations", func(t *testing.T) {
        da := NewDynamicArray(2)

        if !da.IsEmpty() {
            t.Errorf("Expected empty array")
        }

        da.Push(1)
        da.Push(2)
        da.Push(3)

        if da.Size() != 3 {
            t.Errorf("Expected size 3, got %d", da.Size())
        }

        if da.Capacity() != 4 {
            t.Errorf("Expected capacity 4, got %d", da.Capacity())
        }

        val, err := da.At(1)
        if err != nil || val != 2 {
            t.Errorf("Expected value 2 at index 1, got %d", val)
        }

        da.Prepend(0)
        if da.String() != "[0 1 2 3]" {
            t.Errorf("Unexpected array content: %s", da.String())
        }
    })

    t.Run("Pop and Delete", func(t *testing.T) {
        da := NewDynamicArray(4)
        da.Push(1)
        da.Push(2)
        da.Push(3)
        da.Push(4)

        val, err := da.Pop()
        if err != nil || val != 4 {
            t.Errorf("Expected popped value 4, got %d", val)
        }

        err = da.Delete(1)
        if err != nil {
            t.Errorf("Unexpected error on delete: %v", err)
        }

        if da.String() != "[1 3]" {
            t.Errorf("Unexpected array content after delete: %s", da.String())
        }
    })

    t.Run("Find and Remove", func(t *testing.T) {
        da := NewDynamicArray(4)
        da.Push(1)
        da.Push(2)
        da.Push(3)
        da.Push(2)

        index := da.Find(2)
        if index != 1 {
            t.Errorf("Expected to find 2 at index 1, got %d", index)
        }

        removed := da.Remove(2)
        if !removed {
            t.Errorf("Expected element 2 to be removed")
        }

        if da.String() != "[1 3 2]" {
            t.Errorf("Unexpected array content after remove: %s", da.String())
        }
    })

    t.Run("Resize", func(t *testing.T) {
        da := NewDynamicArray(2)
        for i := 0; i < 10; i++ {
            da.Push(i)
        }

        if da.Capacity() != 16 {
            t.Errorf("Expected capacity 16 after growth, got %d", da.Capacity())
        }

        for i := 0; i < 7; i++ {
            da.Pop()
        }

        if da.Capacity() != 8 {
            t.Errorf("Expected capacity 8 after shrink, got %d", da.Capacity())
        }
    })
}
