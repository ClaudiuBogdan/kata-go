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

}

// Push adds an item to the top of the stack
func (s *Stack) Push(item int) {
}

// Pop removes and returns the item from the top of the stack
func (s *Stack) Pop() (int, error) {

}

// Peek returns the item from the top of the stack without removing it
func (s *Stack) Peek() (int, error) {

}

// IsEmpty returns true if the stack is empty
func (s *Stack) IsEmpty() bool {
    
}

// Size returns the number of items in the stack
func (s *Stack) Size() int {

}

// Clear removes all items from the stack
func (s *Stack) Clear() {

}

// String returns a string representation of the stack
func (s *Stack) String() string {
    return fmt.Sprintf("%v", s.items)
}
