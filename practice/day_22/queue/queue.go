package queue

import (
	"fmt"
)

// Queue represents a queue data structure
type Queue struct {
	items []int
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	return &Queue{
		items: make([]int, 0),
	}
}

// Enqueue adds an item to the back of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Front returns the item at the front of the queue without removing it
func (q *Queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}

	return q.items[0], nil
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

// Clear removes all items from the queue
func (q *Queue) Clear() {
	q.items = make([]int, 0)
}

// String returns a string representation of the queue
func (q *Queue) String() string {
	return fmt.Sprintf("%v", q.items)
}
