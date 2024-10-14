# Singly Linked List

## Overview

A Singly Linked List is a linear data structure consisting of nodes where each node points to the next node in the sequence. It provides an efficient way to store and manipulate a collection of elements, offering constant time insertions and deletions at the head, but slower operations for searching and accessing elements compared to an array.

## Implementation Details

The package provides several operations to manage the singly linked list:

### Structs

- `Node`: Represents an individual element with the following fields:
  - `data int`: The value stored in the node.
  - `next *Node`: A pointer to the next node in the list.

- `LinkedList`: Represents the linked list structure with the following fields:
  - `head *Node`: A pointer to the first node of the list.
  - `size int`: The number of elements in the list.

### Functions

- `NewLinkedList() *LinkedList`: Creates and returns a new empty linked list.

- `InsertFront(data int)`: Inserts a new node with the specified data at the front of the list.

- `InsertBack(data int)`: Appends a new node with the specified data to the end of the list.

- `RemoveFront() (int, error)`: Removes and returns the first node from the list. Returns an error if the list is empty.

- `Remove(data int) bool`: Removes the first node containing the specified data. Returns true if successful.

- `Contains(data int) bool`: Checks if the list contains a node with the specified data. Returns true if found.

- `Get(index int) (int, error)`: Returns the data of the node at the specified index. Returns an error if the index is out of bounds.

- `Size() int`: Returns the number of nodes in the list.

- `IsEmpty() bool`: Returns true if the list has no nodes.

- `Clear()`: Removes all nodes from the list.

- `String() string`: Returns a string representation of the list in the format `[elem1 elem2 ...]`.

## Usage

```go
ll := singlylinkedlist.NewLinkedList()
ll.InsertFront(10)
ll.InsertBack(20)
fmt.Println(ll.String()) // Output: "[10 20]"

if ll.Contains(20) {
    fmt.Println("List contains 20")
}

data, err := ll.RemoveFront()
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println("Removed:", data)
}
```

## Complexity Analysis

- Insertions/Deletions at Front: O(1)
- Insertions at End: O(n)
- Deletions by Value: O(n)
- Search: O(n)
- Access by Index: O(n)

## Applications

- Implementing stacks and queues
- Undo functionality in applications
- Memory management systems
- Dynamic data storage

## Advantages and Disadvantages

### Advantages

- Dynamic size: Easily grow and shrink as needed.
- Efficient insertions/deletions at head: Constant time operations.

### Disadvantages

- Access by index is slow: Linear time complexity.
- Increased memory usage: Requires extra space for pointers.

## Testing

The implementation should include tests for:

1. Insertions and deletions
2. Edge cases (empty list operations)
3. Access violations (out-of-bounds indices)
4. Comprehensive tests for all functions

To run the tests, use:

```bash
go test
```

## Limitations

- Inefficient for indexed access: Consider arrays for frequently accessed elements by index.
- No backward traversal: If needed, consider Doubly Linked Lists.

## Related Data Structures

- Doubly Linked List
- Circular Linked List
- Arrays
- Trees

Understanding the practical applications and trade-offs of a singly linked list is crucial when deciding to use it in software design. It provides an excellent resource for dynamic data storage but should be used judiciously based on the application’s requirements.
