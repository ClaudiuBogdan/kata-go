package maxheap

import (
	"cmp"
	"errors"
)

type MaxHeap[T cmp.Ordered] struct {
	elements []T
}

func NewMaxHeap[T cmp.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{elements: make([]T, 0)}
}

// Insert add new element to heap
func (h *MaxHeap[T]) Insert(element T) {
	// Add the element to the back of the list and heapifyUp
	h.elements = append(h.elements, element)
	h.heapifyUp(len(h.elements) - 1)
}

func (h *MaxHeap[T]) Extract() (T, error) {
	// Get first element, move last to first and heapifyDown
	if h.IsEmpty() {
		var zero T
		return zero, errors.New("heap is empty")
	}

	val := h.elements[0]
	idx := len(h.elements) - 1
	h.elements[0] = h.elements[idx]
	h.elements = h.elements[:idx]
	h.heapifyDown(0)

	return val, nil
}

func (h *MaxHeap[T]) Peek() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, errors.New("heap is empty")
	}

	return h.elements[0], nil
}

func (h *MaxHeap[T]) Size() int {
	return len(h.elements)
}

func (h *MaxHeap[T]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *MaxHeap[T]) heapifyUp(idx int) {
	if idx <= 0 {
		return
	}

	parentIdx := (idx - 1) / 2

	if h.elements[parentIdx] < h.elements[idx] {
		h.swap(parentIdx, idx)
		h.heapifyUp(parentIdx)
	}
}

func (h *MaxHeap[T]) heapifyDown(idx int) {
	if idx >= h.Size()-1 {
		return
	}

	leftIdx := idx*2 + 1
	rightIdx := idx*2 + 2

	if leftIdx >= h.Size() {
		return
	}

	maxIdx := leftIdx

	if rightIdx < h.Size() && h.elements[rightIdx] > h.elements[leftIdx] {
		maxIdx = rightIdx
	}

	if h.elements[maxIdx] > h.elements[idx] {
		h.swap(maxIdx, idx)
		h.heapifyDown(maxIdx)
	}
}

func (h *MaxHeap[T]) swap(idxA, idxB int) {
	h.elements[idxA], h.elements[idxB] = h.elements[idxB], h.elements[idxA]
}
