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

func (h *MaxHeap[T]) Insert(element T) {
	h.elements = append(h.elements, element)
	h.heapifyUp(len(h.elements) - 1)
}

func (h *MaxHeap[T]) Extract() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, errors.New("heap is empty")
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
		return zero, errors.New("heap is empty")
	}
	return h.elements[0], nil
}

func (h *MaxHeap[T]) Size() int {
	return len(h.elements)
}

func (h *MaxHeap[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

func (h *MaxHeap[T]) heapifyUp(index int) {
	for index > 0 {
		parentIdx := (index - 1) / 2
		if h.elements[parentIdx] >= h.elements[index] {
			break
		}
		h.swap(index, parentIdx)
		index = parentIdx
	}
}

func (h *MaxHeap[T]) heapifyDown(index int) {
	lastIdx := len(h.elements) - 1
	for {
		leftChildIdx := 2*index + 1
		rightChildIdx := 2*index + 2
		largestIdx := index

		if leftChildIdx <= lastIdx && h.elements[leftChildIdx] > h.elements[largestIdx] {
			largestIdx = leftChildIdx
		}

		if rightChildIdx <= lastIdx && h.elements[rightChildIdx] > h.elements[largestIdx] {
			largestIdx = rightChildIdx
		}

		if largestIdx == index {
			break
		}

		h.swap(index, largestIdx)
		index = largestIdx
	}
}

func (h *MaxHeap[T]) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}