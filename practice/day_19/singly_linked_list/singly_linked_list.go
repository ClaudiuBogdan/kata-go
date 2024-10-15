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
	tail *Node
	size int
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// InsertFront adds a new node at the front of the list
func (ll *LinkedList) InsertFront(data int) {
	newNode := &Node{
		data: data,
		next: ll.head,
	}
	ll.size++
	ll.head = newNode

	if ll.tail == nil {
		ll.tail = newNode
	}
}

// InsertBack adds a new node at the end of the list
func (ll *LinkedList) InsertBack(data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}
	ll.size++
	if ll.tail != nil {
		ll.tail.next = newNode
		ll.tail = newNode
	} else {
		ll.head = newNode
		ll.tail = newNode
	}
}

// RemoveFront removes the first node from the list
func (ll *LinkedList) RemoveFront() (int, error) {
	if ll.head == nil {
		return 0, fmt.Errorf("list is empty")
	}
	ll.size--
	oldHead := ll.head
	ll.head = ll.head.next

	if ll.head == nil {
		ll.tail = nil
	}
	return oldHead.data, nil
}

// Remove removes the first occurrence of the specified data
func (ll *LinkedList) Remove(data int) bool {
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
		ll.head = ll.head.next
	} else if current.next == nil {
		ll.tail = prev
	} else {
		prev.next = current.next
	}

	if ll.head == nil {
		ll.tail = nil
	}

	return true
}

// Contains checks if the list contains the specified data
func (ll *LinkedList) Contains(data int) bool {
	for current := ll.head; current != nil; current = current.next {
		if current.data == data {
			return true
		}
	}
	return false
}

// Get returns the data at the specified index
func (ll *LinkedList) Get(index int) (int, error) {
	currentIndex := 0
	current := ll.head
	for current != nil && currentIndex < index {
		currentIndex++
		current = current.next
	}

	if current == nil {
		return 0, fmt.Errorf("index out of bound")
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
	ll.size = 0
	ll.head = nil
	ll.tail = nil
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
