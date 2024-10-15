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
	return &DynamicArray{
		array:    make([]int, initialCapacity, initialCapacity),
		size:     0,
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
	if index >= da.size {
		return 0, fmt.Errorf("index out of bound")
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
	if index < 0 || index >= da.size {
		return fmt.Errorf("index out of bound")
	}
	if da.size == da.capacity {
		da.resize(da.capacity * 2)
	}

	for i := da.size; i > index; i-- {
		da.array[i] = da.array[i - 1]
	}

	da.array[index] = element
	da.size++
	return nil
}

// Prepend adds an element to the beginning of the array
func (da *DynamicArray) Prepend(element int) {
	if da.size == da.capacity {
		da.resize(da.capacity * 2)
	}
	for i := da.size; i > 0; i-- {
		da.array[i] = da.array[i-1]
	}
	da.array[0] = element
	da.size++
}

// Pop removes and returns the last element
func (da *DynamicArray) Pop() (int, error) {
	if da.IsEmpty() {
		return 0, fmt.Errorf("array is empty")
	}
	if da.size < da.capacity / 2 {
		da.resize(da.capacity / 2)
	}
	element := da.array[da.size-1]
	da.size--
	return element, nil
}

// Delete removes the element at the specified index
func (da *DynamicArray) Delete(index int) error {
	if index < 0 || index >= da.size {
		return fmt.Errorf("index out of bound")
	}
	if da.size < da.capacity / 2 {
		da.resize(da.capacity / 2)
	}
	for i := index; i < da.size-1; i++ {
		da.array[i] = da.array[i+1]
	}
	da.size--
	return nil
}

// Remove removes the first occurrence of the element
func (da *DynamicArray) Remove(element int) bool {
	index := da.Find(element)
	if index == -1 {
		return false
	}
	if da.size < da.capacity / 2 {
		da.resize(da.capacity / 2)
	}
	err := da.Delete(index)

	if err != nil {
		return false
	}

	return true
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
	oldArr := da.array
	newArr := make([]int, newCapacity, newCapacity)

	for i := 0; i < da.size; i++ {
		newArr[i] = oldArr[i]
	}
	da.capacity = newCapacity
	da.array = newArr
}

// String returns a string representation of the array
func (da *DynamicArray) String() string {
	return fmt.Sprintf("%v", da.array[:da.size])
}
