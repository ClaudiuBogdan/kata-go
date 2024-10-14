package singlylinkedlist

import "fmt"

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
    newNode := &Node{
        data: data,
        next: ll.head,
    }

    ll.head = newNode
    ll.size++
}

// InsertBack adds a new node at the end of the list
func (ll *LinkedList) InsertBack(data int) {
    ll.size++
    newNode := &Node{
        data: data,
        next: nil,
    }
    if ll.head == nil {
        ll.head = newNode
        return
    }

    tail := ll.head

    for tail.next != nil {
        tail = tail.next
    }

    tail.next = newNode
}

// RemoveFront removes the first node from the list
func (ll *LinkedList) RemoveFront() (int, error) {
    head := ll.head
    if head == nil {
        return 0, fmt.Errorf("list is empty")
    }
    ll.size--
    ll.head = head.next
    return head.data, nil
}

// Remove removes the first occurrence of the specified data
func (ll *LinkedList) Remove(data int) bool {
    if ll.head == nil {
        return false
    }

    var prev *Node
    current := ll.head
    for current != nil && current.data != data {
        prev = current
        current = current.next
    }

    if current == nil {
        return false
    }
    ll.size--
    if prev == nil {
        ll.head = current.next
    } else {
        prev.next = current.next
    }
    return true
}

// Contains checks if the list contains the specified data
func (ll *LinkedList) Contains(data int) bool {
    current := ll.head

    for current != nil && current.data != data {
        current = current.next
    }

    if current == nil {
        return false
    }
    return true
}

// Get returns the data at the specified index
func (ll *LinkedList) Get(index int) (int, error) {
    if index >= ll.size{
        return 0, fmt.Errorf("index out of bound")
    }

    current := ll.head
    idx := 0

    for idx < index {
        idx++
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
