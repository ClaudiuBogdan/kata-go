package maxheap

import (
	"cmp"
	"fmt"
)

// MaxHeap represents a max heap data structure.
type MaxHeap[T cmp.Ordered] struct {
	elements []T
}

// New creates and returns a new empty MaxHeap.
func New[T cmp.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{
		elements: make([]T, 0),
	}
}

// Insert adds a new element to the heap.
func (h *MaxHeap[T]) Insert(element T) {
	h.elements = append(h.elements, element)
	h.bubbleUp(len(h.elements) - 1)
}

// Extract removes and returns the maximum element from the heap.
// It returns an error if the heap is empty.
func (h *MaxHeap[T]) Extract() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("heap is empty")
	}

	maxVal := h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	if h.Size() > 0 {
		h.bubbleDown(0)
	}
	return maxVal, nil
}

// Peek returns the maximum element without removing it from the heap.
// It returns an error if the heap is empty.
func (h *MaxHeap[T]) Peek() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("heap is empty")
	}
	return h.elements[0], nil
}

// Size returns the number of elements in the heap.
func (h *MaxHeap[T]) Size() int {
	return len(h.elements)
}

// IsEmpty returns true if the heap contains no elements.
func (h *MaxHeap[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

func (h *MaxHeap[T]) bubbleUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2

		if h.elements[index] > h.elements[parent] {
			h.swap(index, parent)
			index = parent
		} else {
			break
		}
	}
}

func (h *MaxHeap[T]) bubbleDown(index int) {
	for {
		if index >= h.Size() {
			return
		}
		left := index*2 + 1
		right := index*2 + 2

		maxIndex := index

		if left < h.Size() && h.elements[left] > h.elements[maxIndex] {
			maxIndex = left
		}

		if right < h.Size() && h.elements[right] > h.elements[maxIndex] {
			maxIndex = right
		}

		if maxIndex == index {
			break
		}

		h.swap(index, maxIndex)
		index = maxIndex
	}
}

func (h *MaxHeap[T]) swap(a, b int) {
	h.elements[a], h.elements[b] = h.elements[b], h.elements[a]
}
