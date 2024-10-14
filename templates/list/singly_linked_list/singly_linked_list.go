package singlylinkedlist

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

}

// InsertFront adds a new node at the front of the list
func (ll *LinkedList) InsertFront(data int) {

}

// InsertBack adds a new node at the end of the list
func (ll *LinkedList) InsertBack(data int) {

}

// RemoveFront removes the first node from the list
func (ll *LinkedList) RemoveFront() (int, error) {

}

// Remove removes the first occurrence of the specified data
func (ll *LinkedList) Remove(data int) bool {

}

// Contains checks if the list contains the specified data
func (ll *LinkedList) Contains(data int) bool {

}

// Get returns the data at the specified index
func (ll *LinkedList) Get(index int) (int, error) {

}

// Size returns the number of nodes in the list
func (ll *LinkedList) Size() int {

}

// IsEmpty returns true if the list is empty
func (ll *LinkedList) IsEmpty() bool {

}

// Clear removes all nodes from the list
func (ll *LinkedList) Clear() {

}

// String returns a string representation of the list
func (ll *LinkedList) String() string {

}
