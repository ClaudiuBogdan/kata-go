package stack

import (
	"fmt"
)

// Stack represents a stack data structure
type Stack struct {
	items []int
}

// NewStack creates a new empty stack
func NewStack() *Stack {
	return &Stack{
		items: []int(nil),
	}
}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	idx := len(s.items) - 1
	item := s.items[idx]
	s.items = s.items[:idx]
	return item, nil
}

// Peek returns the item from the top of the stack without removing it
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	idx := len(s.items) - 1
	return s.items[idx], nil
}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Clear removes all items from the stack
func (s *Stack) Clear() {
	s.items = []int(nil)
}

// String returns a string representation of the stack
func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.items)
}
