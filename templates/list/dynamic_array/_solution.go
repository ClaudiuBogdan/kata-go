package dynamicarray

import (
    "errors"
    "fmt"
)

// DynamicArray represents a dynamic array data structure
type DynamicArray struct {
    array []int
    size  int
    capacity int
}

// NewDynamicArray creates a new dynamic array with initial capacity
func NewDynamicArray(initialCapacity int) *DynamicArray {
    return &DynamicArray{
        array: make([]int, initialCapacity),
        size: 0,
        capacity: initialCapacity,
    }
}

// Size returns the number of elements in the array
func (da *DynamicArray) Size() int {
    return da.size
}

// Capacity returns the current capacity of the array
func (da *DynamicArray) Capacity() int {
    return da.capacity
}

// IsEmpty returns true if the array is empty
func (da *DynamicArray) IsEmpty() bool {
    return da.size == 0
}

// At returns the element at the given index
func (da *DynamicArray) At(index int) (int, error) {
    if index < 0 || index >= da.size {
        return 0, errors.New("index out of bounds")
    }
    return da.array[index], nil
}

// Push adds an element to the end of the array
func (da *DynamicArray) Push(element int) {
    if da.size == da.capacity {
        da.resize(da.capacity * 2)
    }
    da.array[da.size] = element
    da.size++
}

// Insert adds an element at the specified index
func (da *DynamicArray) Insert(index int, element int) error {
    if index < 0 || index > da.size {
        return errors.New("index out of bounds")
    }
    if da.size == da.capacity {
        da.resize(da.capacity * 2)
    }
    for i := da.size; i > index; i-- {
        da.array[i] = da.array[i-1]
    }
    da.array[index] = element
    da.size++
    return nil
}

// Prepend adds an element to the beginning of the array
func (da *DynamicArray) Prepend(element int) {
    da.Insert(0, element)
}

// Pop removes and returns the last element
func (da *DynamicArray) Pop() (int, error) {
    if da.IsEmpty() {
        return 0, errors.New("array is empty")
    }
    da.size--
    element := da.array[da.size]
    if da.size > 0 && da.size == da.capacity/4 {
        da.resize(da.capacity / 2)
    }
    return element, nil
}

// Delete removes the element at the specified index
func (da *DynamicArray) Delete(index int) error {
    if index < 0 || index >= da.size {
        return errors.New("index out of bounds")
    }
    for i := index; i < da.size-1; i++ {
        da.array[i] = da.array[i+1]
    }
    da.size--
    if da.size > 0 && da.size == da.capacity/4 {
        da.resize(da.capacity / 2)
    }
    return nil
}

// Remove removes the first occurrence of the element
func (da *DynamicArray) Remove(element int) bool {
    for i := 0; i < da.size; i++ {
        if da.array[i] == element {
            da.Delete(i)
            return true
        }
    }
    return false
}

// Find returns the index of the first occurrence of the element
func (da *DynamicArray) Find(element int) int {
    for i := 0; i < da.size; i++ {
        if da.array[i] == element {
            return i
        }
    }
    return -1
}

// resize changes the capacity of the array
func (da *DynamicArray) resize(newCapacity int) {
    newArray := make([]int, newCapacity)
    copy(newArray, da.array)
    da.array = newArray
    da.capacity = newCapacity
}

// String returns a string representation of the array
func (da *DynamicArray) String() string {
    return fmt.Sprintf("%v", da.array[:da.size])
}
