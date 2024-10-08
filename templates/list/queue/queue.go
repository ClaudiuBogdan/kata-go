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

}

// Enqueue adds an item to the back of the queue
func (q *Queue) Enqueue(item int) {

}

// Dequeue removes and returns the item from the front of the queue
func (q *Queue) Dequeue() (int, error) {

}

// Front returns the item at the front of the queue without removing it
func (q *Queue) Front() (int, error) {

}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {

}

// Size returns the number of items in the queue
func (q *Queue) Size() int {

}

// Clear removes all items from the queue
func (q *Queue) Clear() {

}

// String returns a string representation of the queue
func (q *Queue) String() string {
	return fmt.Sprintf("%v", q.items)
}
