package stack

import (
    "testing"
)

func TestStack(t *testing.T) {
    t.Run("Basic Operations", func(t *testing.T) {
        s := NewStack()

        if !s.IsEmpty() {
            t.Errorf("Expected empty stack")
        }

        s.Push(1)
        s.Push(2)
        s.Push(3)

        if s.Size() != 3 {
            t.Errorf("Expected size 3, got %d", s.Size())
        }

        val, err := s.Peek()
        if err != nil || val != 3 {
            t.Errorf("Expected peek value 3, got %d", val)
        }

        if s.String() != "[1 2 3]" {
            t.Errorf("Unexpected stack content: %s", s.String())
        }
    })

    t.Run("Pop", func(t *testing.T) {
        s := NewStack()
        s.Push(1)
        s.Push(2)
        s.Push(3)

        val, err := s.Pop()
        if err != nil || val != 3 {
            t.Errorf("Expected popped value 3, got %d", val)
        }

        if s.Size() != 2 {
            t.Errorf("Expected size 2 after pop, got %d", s.Size())
        }

        if s.String() != "[1 2]" {
            t.Errorf("Unexpected stack content after pop: %s", s.String())
        }
    })

    t.Run("Empty Stack Operations", func(t *testing.T) {
        s := NewStack()

        _, err := s.Pop()
        if err == nil {
            t.Errorf("Expected error when popping from empty stack")
        }

        _, err = s.Peek()
        if err == nil {
            t.Errorf("Expected error when peeking empty stack")
        }
    })

    t.Run("Clear", func(t *testing.T) {
        s := NewStack()
        s.Push(1)
        s.Push(2)
        s.Push(3)

        s.Clear()

        if !s.IsEmpty() {
            t.Errorf("Expected empty stack after clear")
        }

        if s.Size() != 0 {
            t.Errorf("Expected size 0 after clear, got %d", s.Size())
        }
    })
}
