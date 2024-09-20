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
	return &MaxHeap[T]{elements: make([]T, 0)}
}

// Insert adds a new element to the heap.
func (h *MaxHeap[T]) Insert(element T) {
	h.elements = append(h.elements, element)
	h.heapifyUp(h.Size() - 1)
}

// Extract removes and returns the maximum element from the heap.
// It returns an error if the heap is empty.
func (h *MaxHeap[T]) Extract() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("heap is empty")
	}
	lastIdx := len(h.elements) - 1
	max := h.elements[0]
	h.elements[0] = h.elements[lastIdx]
	h.elements = h.elements[:lastIdx]

	if !h.IsEmpty() {
		h.heapifyDown(0)
	}

	return max, nil
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

func (h *MaxHeap[T]) heapifyUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2

		if h.elements[idx] > h.elements[parent] {
			h.swap(idx, parent)
			idx = parent
		} else {
			break
		}
	}
}

func (h *MaxHeap[T]) heapifyDown(idx int) {
	for {
		last := idx
		left := idx*2 + 1
		right := idx*2 + 2


		if left < h.Size() && h.elements[left] > h.elements[last] {
			last = left
		}
		if right < h.Size() && h.elements[right] > h.elements[last] {
			last = right
		}

		if idx == last {
			return
		}

		h.swap(idx, last)
		idx = last
	}
}

func (h *MaxHeap[T]) swap(idxA, idxB int) {
	h.elements[idxA], h.elements[idxB] = h.elements[idxB], h.elements[idxA]
}
