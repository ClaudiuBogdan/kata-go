# Stack Data Structure Implementation

## Overview

This package provides an implementation of the Stack data structure in Go. A stack is a Last-In-First-Out (LIFO) data structure that supports two primary operations: push (add an element) and pop (remove the most recently added element).

## Features

- Push: Add an element to the top of the stack
- Pop: Remove and return the top element from the stack
- Peek: Return the top element without removing it
- IsEmpty: Check if the stack is empty
- Size: Get the current number of elements in the stack

## Implementation Details

The stack is implemented using a slice in Go, which provides dynamic resizing capabilities. The implementation is generic, allowing the stack to hold elements of any type.

### Main Structure

```go
type Stack[T any] struct {
    elements []T
}
```

### Primary Functions

1. `NewStack[T any]() *Stack[T]`: Creates and returns a new empty stack.
2. `(s *Stack[T]) Push(element T)`: Adds an element to the top of the stack.
3. `(s *Stack[T]) Pop() (T, error)`: Removes and returns the top element from the stack.
4. `(s *Stack[T]) Peek() (T, error)`: Returns the top element without removing it.
5. `(s *Stack[T]) IsEmpty() bool`: Returns true if the stack is empty, false otherwise.
6. `(s *Stack[T]) Size() int`: Returns the number of elements in the stack.

## Usage

Here's a basic example of how to use the Stack:

```go
stack := NewStack[int]()

stack.Push(1)
stack.Push(2)
stack.Push(3)

top, err := stack.Peek()
if err == nil {
    fmt.Println("Top element:", top)  // Output: Top element: 3
}

popped, err := stack.Pop()
if err == nil {
    fmt.Println("Popped element:", popped)  // Output: Popped element: 3
}

fmt.Println("Stack size:", stack.Size())  // Output: Stack size: 2
```

## Error Handling

The `Pop()` and `Peek()` methods return an error if the operation is performed on an empty stack. Always check for errors when using these methods.

## Time Complexity

- Push: O(1) amortized
- Pop: O(1)
- Peek: O(1)
- IsEmpty: O(1)
- Size: O(1)

## Space Complexity

The space complexity is O(n), where n is the number of elements in the stack.

## Testing

The implementation includes a comprehensive test suite covering all main operations and edge cases. To run the tests, use the following command:

```bash
go test
```

## Applications of Stacks

1. Function call management (call stack)
2. Undo mechanisms in text editors
3. Expression evaluation and syntax parsing
4. Backtracking algorithms
5. Browser history (back button functionality)

## Limitations

- The current implementation doesn't have a maximum size limit. For specific use cases, you might want to add a capacity check.
- As with any slice-based implementation, there's a trade-off between memory usage and performance for very large stacks.

## Future Enhancements

1. Add a max capacity option
2. Implement iterator functionality
3. Add thread-safe operations for concurrent use

## Contributing

Contributions to improve the implementation or add new features are welcome. Please ensure that all tests pass before submitting a pull request.

## License

[Specify your license here]

---

Remember to replace [Specify your license here] with the appropriate license for your project. This README provides a comprehensive overview of the Stack implementation, its usage, and considerations. It can be adjusted based on specific implementation details or additional features you might add to your Stack data structure.
