package ringbuffer

import (
    "testing"
)

func TestRingBuffer(t *testing.T) {
    t.Run("Basic Operations", func(t *testing.T) {
        rb := NewRingBuffer(5)

        if !rb.IsEmpty() {
            t.Errorf("Expected empty buffer")
        }

        err := rb.Write(1)
        if err != nil {
            t.Errorf("Unexpected error on write: %v", err)
        }

        if rb.Size() != 1 {
            t.Errorf("Expected size 1, got %d", rb.Size())
        }

        val, err := rb.Peek()
        if err != nil || val != 1 {
            t.Errorf("Expected peek value 1, got %d", val)
        }

        if rb.String() != "[1]" {
            t.Errorf("Unexpected buffer content: %s", rb.String())
        }
    })

    t.Run("Full Buffer", func(t *testing.T) {
        rb := NewRingBuffer(3)

        rb.Write(1)
        rb.Write(2)
        rb.Write(3)

        if !rb.IsFull() {
            t.Errorf("Expected buffer to be full")
        }

        err := rb.Write(4)
        if err == nil {
            t.Errorf("Expected error when writing to full buffer")
        }

        if rb.String() != "[1 2 3]" {
            t.Errorf("Unexpected buffer content: %s", rb.String())
        }
    })

    t.Run("Read and Write", func(t *testing.T) {
        rb := NewRingBuffer(3)

        rb.Write(1)
        rb.Write(2)
        rb.Write(3)

        val, err := rb.Read()
        if err != nil || val != 1 {
            t.Errorf("Expected read value 1, got %d", val)
        }

        rb.Write(4)

        if rb.String() != "[2 3 4]" {
            t.Errorf("Unexpected buffer content: %s", rb.String())
        }
    })

    t.Run("Wrap Around", func(t *testing.T) {
        rb := NewRingBuffer(3)

        rb.Write(1)
        rb.Write(2)
        rb.Write(3)
        rb.Read()
        rb.Read()
        rb.Write(4)
        rb.Write(5)

        if rb.String() != "[3 4 5]" {
            t.Errorf("Unexpected buffer content: %s", rb.String())
        }
    })

    t.Run("Clear", func(t *testing.T) {
        rb := NewRingBuffer(3)

        rb.Write(1)
        rb.Write(2)
        rb.Clear()

        if !rb.IsEmpty() {
            t.Errorf("Expected empty buffer after clear")
        }

        if rb.Size() != 0 {
            t.Errorf("Expected size 0 after clear, got %d", rb.Size())
        }
    })
}
