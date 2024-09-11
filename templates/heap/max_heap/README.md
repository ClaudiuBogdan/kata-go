# Max Heap Data Structure Implementation

## Overview

This package provides an implementation of the Max Heap data structure in Go. A Max Heap is a complete binary tree where the parent node is always greater than or equal to its children. The root node contains the maximum element in the heap.

## Features

- Insert: Add an element to the heap
- Extract: Remove and return the maximum element from the heap
- Peek: Return the maximum element without removing it
- Size: Get the current number of elements in the heap
- IsEmpty: Check if the heap is empty

## Implementation Details

The Max Heap is implemented using a slice in Go, which provides dynamic resizing capabilities. The implementation is generic, allowing the heap to hold elements of any comparable type.

### Main Structure

```go
type MaxHeap[T cmp.Ordered] struct {
    elements []T
}
```

### Primary Functions

1. `NewMaxHeap[T cmp.Ordered]() *MaxHeap[T]`: Creates and returns a new empty max heap.
2. `(h *MaxHeap[T]) Insert(element T)`: Adds an element to the heap.
3. `(h *MaxHeap[T]) Extract() (T, error)`: Removes and returns the maximum element from the heap.
4. `(h *MaxHeap[T]) Peek() (T, error)`: Returns the maximum element without removing it.
5. `(h *MaxHeap[T]) Size() int`: Returns the number of elements in the heap.
6. `(h *MaxHeap[T]) IsEmpty() bool`: Returns true if the heap is empty, false otherwise.

## Usage

Here's a basic example of how to use the Max Heap:

```go
heap := NewMaxHeap[int]()

heap.Insert(10)
heap.Insert(5)
heap.Insert(15)
heap.Insert(8)

max, err := heap.Peek()
if err == nil {
    fmt.Println("Maximum element:", max)  // Output: Maximum element: 15
}

extractedMax, err := heap.Extract()
if err == nil {
    fmt.Println("Extracted max:", extractedMax)  // Output: Extracted max: 15
}

fmt.Println("Heap size:", heap.Size())  // Output: Heap size: 3
```

## Error Handling

The `Extract()` and `Peek()` methods return an error if the operation is performed on an empty heap. Always check for errors when using these methods.

## Time Complexity

- Insert: O(log n)
- Extract: O(log n)
- Peek: O(1)
- Size: O(1)
- IsEmpty: O(1)

Where n is the number of elements in the heap.

## Space Complexity

The space complexity is O(n), where n is the number of elements in the heap.

## Implementation Details

The heap is implemented using the following helper methods:

- `heapifyUp()`: Maintains the heap property after insertion by moving the newly inserted element up the heap.
- `heapifyDown()`: Maintains the heap property after extraction by moving the root element down the heap.
- `swap()`: Swaps two elements in the heap.

These methods ensure that the heap property is maintained after each operation.

## Testing

The implementation includes a comprehensive test suite covering all main operations and edge cases. To run the tests, use the following command:

```bash
go test
```

## Applications of Max Heaps

1. Priority queues
2. Scheduling algorithms
3. Graph algorithms (e.g., Dijkstra's shortest path)
4. K-way merge
5. Finding the k largest elements in a collection

## Limitations

- The current implementation doesn't have a maximum size limit. For specific use cases, you might want to add a capacity check.
- As with any slice-based implementation, there's a trade-off between memory usage and performance for very large heaps.

## Future Enhancements

1. Add a max capacity option
2. Implement decrease key operation
3. Add support for custom comparison functions
4. Implement a min-heap version
5. Add concurrency support for thread-safe operations

## Contributing

Contributions to improve the implementation or add new features are welcome. Please ensure that all tests pass before submitting a pull request.

## License

[Specify your license here]

---

Remember to replace [Specify your license here] with the appropriate license for your project. This README provides a comprehensive overview of the Max Heap implementation, its usage, and considerations. It can be adjusted based on specific implementation details or additional features you might add to your Max Heap data structure.
