package maxheap

import (
	"cmp"
)

// MaxHeap represents a max heap data structure.
type MaxHeap[T cmp.Ordered] struct {
	elements []T
}

// New creates and returns a new empty MaxHeap.
func New[T cmp.Ordered]() *MaxHeap[T] {

}

// Insert adds a new element to the heap.
func (h *MaxHeap[T]) Insert(element T) {

}

// Extract removes and returns the maximum element from the heap.
// It returns an error if the heap is empty.
func (h *MaxHeap[T]) Extract() (T, error) {

}

// Peek returns the maximum element without removing it from the heap.
// It returns an error if the heap is empty.
func (h *MaxHeap[T]) Peek() (T, error) {

}

// Size returns the number of elements in the heap.
func (h *MaxHeap[T]) Size() int {

}

// IsEmpty returns true if the heap contains no elements.
func (h *MaxHeap[T]) IsEmpty() bool {

}
