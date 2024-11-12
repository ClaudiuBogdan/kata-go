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

	ll.head = newNode

	if ll.size == 0 {
		ll.tail = newNode
	}

	ll.size++
}

// InsertBack adds a new node at the end of the list
func (ll *LinkedList) InsertBack(data int) {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if ll.size == 0 {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
	ll.size++
}

// RemoveFront removes the first node from the list
func (ll *LinkedList) RemoveFront() (int, error) {
	if ll.size == 0 {
		return 0, fmt.Errorf("list is empty")
	}
	next := ll.head.next
	value := ll.head.data
	ll.head = next
	ll.size--

	return value, nil
}

// Remove removes the first occurrence of the specified data
func (ll *LinkedList) Remove(data int) bool {
	var prev *Node
	node := ll.head

	for node != nil && node.data != data {
		prev = node
		node = node.next
	}

	if node == nil {
		return false
	}

	if ll.size == 1 {
		ll.head = nil
		ll.tail = nil
		ll.size--
		return true
	}

	if node == ll.head {
		ll.head = ll.head.next
		ll.size--
		return true
	}

	if node == ll.tail {
		ll.tail = prev
		prev.next = nil
		ll.size--
		return true
	}

	prev.next = node.next
	ll.size--
	return true
}

// Contains checks if the list contains the specified data
func (ll *LinkedList) Contains(data int) bool {
	node := ll.head

	for node != nil && node.data != data {
		node = node.next
	}

	if node == nil {
		return false
	}
	return true
}

// Get returns the data at the specified index
func (ll *LinkedList) Get(index int) (int, error) {
	if index < 0 || index >= ll.size {
		return 0, fmt.Errorf("invalid index")
	}

	node := ll.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node.data, nil

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
	ll.tail = nil
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
