package dynamicarray

import "fmt"

// DynamicArray represents a dynamic array data structure
type DynamicArray struct {
	array    []int
	size     int
	capacity int
}

// NewDynamicArray creates a new dynamic array with initial capacity
func NewDynamicArray(initialCapacity int) *DynamicArray {

}

// Size returns the number of elements in the array
func (da *DynamicArray) Size() int {

}

// Capacity returns the current capacity of the array
func (da *DynamicArray) Capacity() int {

}

// IsEmpty returns true if the array is empty
func (da *DynamicArray) IsEmpty() bool {

}

// At returns the element at the given index
func (da *DynamicArray) At(index int) (int, error) {

}

// Push adds an element to the end of the array
func (da *DynamicArray) Push(element int) {

}

// Insert adds an element at the specified index
func (da *DynamicArray) Insert(index int, element int) error {

}

// Prepend adds an element to the beginning of the array
func (da *DynamicArray) Prepend(element int) {

}

// Pop removes and returns the last element
func (da *DynamicArray) Pop() (int, error) {

}

// Delete removes the element at the specified index
func (da *DynamicArray) Delete(index int) error {

}

// Remove removes the first occurrence of the element
func (da *DynamicArray) Remove(element int) bool {

}

// Find returns the index of the first occurrence of the element
func (da *DynamicArray) Find(element int) int {

}

// resize changes the capacity of the array
func (da *DynamicArray) resize(newCapacity int) {

}

// String returns a string representation of the array
func (da *DynamicArray) String() string {
	return fmt.Sprintf("%v", da.array[:da.size])
}
