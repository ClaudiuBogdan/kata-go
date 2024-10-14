package singlylinkedlist

import (
    "testing"
)

func TestLinkedList(t *testing.T) {
    t.Run("Basic Operations", func(t *testing.T) {
        ll := NewLinkedList()

        if !ll.IsEmpty() {
            t.Errorf("Expected empty list")
        }

        ll.InsertFront(1)
        ll.InsertBack(2)
        ll.InsertFront(0)

        if ll.Size() != 3 {
            t.Errorf("Expected size 3, got %d", ll.Size())
        }

        if ll.String() != "[0 1 2]" {
            t.Errorf("Unexpected list content: %s", ll.String())
        }
    })

    t.Run("Remove", func(t *testing.T) {
        ll := NewLinkedList()
        ll.InsertBack(1)
        ll.InsertBack(2)
        ll.InsertBack(3)

        removed := ll.Remove(2)
        if !removed {
            t.Errorf("Expected true when removing existing element")
        }

        if ll.Size() != 2 {
            t.Errorf("Expected size 2 after remove, got %d", ll.Size())
        }

        if ll.String() != "[1 3]" {
            t.Errorf("Unexpected list content after remove: %s", ll.String())
        }

        removed = ll.Remove(4)
        if removed {
            t.Errorf("Expected false when removing non-existent element")
        }
    })

    t.Run("Get", func(t *testing.T) {
        ll := NewLinkedList()
        ll.InsertBack(1)
        ll.InsertBack(2)
        ll.InsertBack(3)

        val, err := ll.Get(1)
        if err != nil || val != 2 {
            t.Errorf("Expected value 2 at index 1, got %d", val)
        }

        _, err = ll.Get(3)
        if err == nil {
            t.Errorf("Expected error when getting out of bounds index")
        }
    })

    t.Run("Contains", func(t *testing.T) {
        ll := NewLinkedList()
        ll.InsertBack(1)
        ll.InsertBack(2)
        ll.InsertBack(3)

        if !ll.Contains(2) {
            t.Errorf("Expected list to contain 2")
        }

        if ll.Contains(4) {
            t.Errorf("Expected list to not contain 4")
        }
    })

    t.Run("Clear", func(t *testing.T) {
        ll := NewLinkedList()
        ll.InsertBack(1)
        ll.InsertBack(2)
        ll.InsertBack(3)

        ll.Clear()

        if !ll.IsEmpty() {
            t.Errorf("Expected empty list after clear")
        }

        if ll.Size() != 0 {
            t.Errorf("Expected size 0 after clear, got %d", ll.Size())
        }
    })
}
