package ringbuffer

import (
    "errors"
    "fmt"
)

// RingBuffer represents a ring buffer data structure
type RingBuffer struct {
    buffer []int
    size   int
    capacity int
    read   int
    write  int
}

// NewRingBuffer creates a new ring buffer with the given capacity
func NewRingBuffer(capacity int) *RingBuffer {
    return &RingBuffer{
        buffer:   make([]int, capacity),
        size:     0,
        capacity: capacity,
        read:     0,
        write:    0,
    }
}

// Write adds an item to the buffer
func (rb *RingBuffer) Write(item int) error {
    if rb.IsFull() {
        return errors.New("buffer is full")
    }
    rb.buffer[rb.write] = item
    rb.write = (rb.write + 1) % rb.capacity
    rb.size++
    return nil
}

// Read removes and returns an item from the buffer
func (rb *RingBuffer) Read() (int, error) {
    if rb.IsEmpty() {
        return 0, errors.New("buffer is empty")
    }
    item := rb.buffer[rb.read]
    rb.read = (rb.read + 1) % rb.capacity
    rb.size--
    return item, nil
}

// Peek returns the next item to be read without removing it
func (rb *RingBuffer) Peek() (int, error) {
    if rb.IsEmpty() {
        return 0, errors.New("buffer is empty")
    }
    return rb.buffer[rb.read], nil
}

// IsEmpty returns true if the buffer is empty
func (rb *RingBuffer) IsEmpty() bool {
    return rb.size == 0
}

// IsFull returns true if the buffer is full
func (rb *RingBuffer) IsFull() bool {
    return rb.size == rb.capacity
}

// Size returns the number of items in the buffer
func (rb *RingBuffer) Size() int {
    return rb.size
}

// Capacity returns the total capacity of the buffer
func (rb *RingBuffer) Capacity() int {
    return rb.capacity
}

// Clear removes all items from the buffer
func (rb *RingBuffer) Clear() {
    rb.size = 0
    rb.read = 0
    rb.write = 0
}

// String returns a string representation of the buffer
func (rb *RingBuffer) String() string {
    items := make([]int, rb.size)
    for i := 0; i < rb.size; i++ {
        items[i] = rb.buffer[(rb.read+i) % rb.capacity]
    }
    return fmt.Sprintf("%v", items)
}
