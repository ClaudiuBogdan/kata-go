package queue

import (
    "testing"
)

func TestQueue(t *testing.T) {
    t.Run("Basic Operations", func(t *testing.T) {
        q := NewQueue()

        if !q.IsEmpty() {
            t.Errorf("Expected empty queue")
        }

        q.Enqueue(1)
        q.Enqueue(2)
        q.Enqueue(3)

        if q.Size() != 3 {
            t.Errorf("Expected size 3, got %d", q.Size())
        }

        val, err := q.Front()
        if err != nil || val != 1 {
            t.Errorf("Expected front value 1, got %d", val)
        }

        if q.String() != "[1 2 3]" {
            t.Errorf("Unexpected queue content: %s", q.String())
        }
    })

    t.Run("Dequeue", func(t *testing.T) {
        q := NewQueue()
        q.Enqueue(1)
        q.Enqueue(2)
        q.Enqueue(3)

        val, err := q.Dequeue()
        if err != nil || val != 1 {
            t.Errorf("Expected dequeued value 1, got %d", val)
        }

        if q.Size() != 2 {
            t.Errorf("Expected size 2 after dequeue, got %d", q.Size())
        }

        if q.String() != "[2 3]" {
            t.Errorf("Unexpected queue content after dequeue: %s", q.String())
        }
    })

    t.Run("Empty Queue Operations", func(t *testing.T) {
        q := NewQueue()

        _, err := q.Dequeue()
        if err == nil {
            t.Errorf("Expected error when dequeuing from empty queue")
        }

        _, err = q.Front()
        if err == nil {
            t.Errorf("Expected error when getting front of empty queue")
        }
    })

    t.Run("Clear", func(t *testing.T) {
        q := NewQueue()
        q.Enqueue(1)
        q.Enqueue(2)
        q.Enqueue(3)

        q.Clear()

        if !q.IsEmpty() {
            t.Errorf("Expected empty queue after clear")
        }

        if q.Size() != 0 {
            t.Errorf("Expected size 0 after clear, got %d", q.Size())
        }
    })
}
