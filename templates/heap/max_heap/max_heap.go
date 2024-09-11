package maxheap

import (
	"cmp"
)

type MaxHeap[T cmp.Ordered] struct {
	elements []T
}

func NewMaxHeap[T cmp.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{elements: make([]T, 0)}
}

func (h *MaxHeap[T]) Insert(element T) {

}

func (h *MaxHeap[T]) Extract() (T, error) {

}

func (h *MaxHeap[T]) Peek() (T, error) {

}

func (h *MaxHeap[T]) Size() int {

}

func (h *MaxHeap[T]) IsEmpty() bool {

}
