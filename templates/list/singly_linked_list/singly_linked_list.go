package singlylinkedlist

import (
    "errors"
    "fmt"
)

// Node represents a node in the linked list
type Node struct {
    data int
    next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
    head *Node
    size int
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
    return &LinkedList{
        head: nil,
        size: 0,
    }
}

// InsertFront adds a new node at the front of the list
func (ll *LinkedList) InsertFront(data int) {
    newNode := &Node{data: data, next: ll.head}
    ll.head = newNode
    ll.size++
}

// InsertBack adds a new node at the end of the list
func (ll *LinkedList) InsertBack(data int) {
    newNode := &Node{data: data, next: nil}
    if ll.head == nil {
        ll.head = newNode
    } else {
        current := ll.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
    ll.size++
}

// RemoveFront removes the first node from the list
func (ll *LinkedList) RemoveFront() (int, error) {
    if ll.IsEmpty() {
        return 0, errors.New("list is empty")
    }
    data := ll.head.data
    ll.head = ll.head.next
    ll.size--
    return data, nil
}

// Remove removes the first occurrence of the specified data
func (ll *LinkedList) Remove(data int) bool {
    if ll.IsEmpty() {
        return false
    }
    if ll.head.data == data {
        ll.head = ll.head.next
        ll.size--
        return true
    }
    current := ll.head
    for current.next != nil {
        if current.next.data == data {
            current.next = current.next.next
            ll.size--
            return true
        }
        current = current.next
    }
    return false
}

// Contains checks if the list contains the specified data
func (ll *LinkedList) Contains(data int) bool {
    current := ll.head
    for current != nil {
        if current.data == data {
            return true
        }
        current = current.next
    }
    return false
}

// Get returns the data at the specified index
func (ll *LinkedList) Get(index int) (int, error) {
    if index < 0 || index >= ll.size {
        return 0, errors.New("index out of bounds")
    }
    current := ll.head
    for i := 0; i < index; i++ {
        current = current.next
    }
    return current.data, nil
}

// Size returns the number of nodes in the list
func (ll *LinkedList) Size() int {
    return ll.size
}

// IsEmpty returns true if the list is empty
func (ll *LinkedList) IsEmpty() bool {
    return ll.size == 0
}

// Clear removes all nodes from the list
func (ll *LinkedList) Clear() {
    ll.head = nil
    ll.size = 0
}

// String returns a string representation of the list
func (ll *LinkedList) String() string {
    if ll.IsEmpty() {
        return "[]"
    }
    var result string
    current := ll.head
    for current != nil {
        result += fmt.Sprintf("%d ", current.data)
        current = current.next
    }
    return "[" + result[:len(result)-1] + "]"
}
