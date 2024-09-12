package maxheap

import (
	"cmp"
	"fmt"
)

type MaxHeap[T cmp.Ordered] struct {
	elements []T
}

func NewMaxHeap[T cmp.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{elements: make([]T, 0)}
}

func (h *MaxHeap[T]) Insert(element T) {
	h.elements = append(h.elements, element)
	h.heapifyUp(len(h.elements) - 1)
}

func (h *MaxHeap[T]) Extract() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("heap is empty")
	}

	max := h.elements[0]
	lastIdx := len(h.elements) - 1

	h.elements[0] = h.elements[lastIdx]
	h.elements = h.elements[:lastIdx]

	if !h.IsEmpty() {
		h.heapifyDown(0)
	}

	return max, nil
}

func (h *MaxHeap[T]) Peek() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("heap is empty")
	}

	return h.elements[0], nil
}

func (h *MaxHeap[T]) Size() int {
	return len(h.elements)
}

func (h *MaxHeap[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

func (h *MaxHeap[T]) heapifyUp(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2

		if h.elements[idx] <= h.elements[parentIdx] {
			break
		}

		h.swap(idx, parentIdx)
		idx = parentIdx
	}
}

func (h *MaxHeap[T]) heapifyDown(idx int) {
	for {
		largest := idx
		left := idx*2 + 1
		right := idx*2 + 2

		if left < h.Size() && h.elements[left] > h.elements[largest] {
			largest = left
		}

		if right < h.Size() && h.elements[right] > h.elements[largest] {
			largest = right
		}

		if largest == idx {
			break
		}

		h.swap(largest, idx)
		idx = largest
	}
}

func (h *MaxHeap[T]) swap(a, b int) {
	h.elements[a], h.elements[b] = h.elements[b], h.elements[a]
}
