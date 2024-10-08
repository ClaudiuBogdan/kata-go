package lrucache

import (
	"fmt"
)

// Node represents a node in the doubly linked list
type Node struct {
    key   int
    value int
    prev  *Node
    next  *Node
}

// LRUCache represents the LRU cache
type LRUCache struct {
    capacity int
    cache    map[int]*Node
    head     *Node
    tail     *Node
}

// NewLRUCache creates a new LRU cache with the given capacity
func NewLRUCache(capacity int) *LRUCache {

}

// Get retrieves a value from the cache
func (lru *LRUCache) Get(key int) (int, error) {

}

// Put adds a key-value pair to the cache
func (lru *LRUCache) Put(key, value int) {

}

// moveToFront moves a node to the front of the list
func (lru *LRUCache) moveToFront(node *Node) {

}

// addToFront adds a node to the front of the list
func (lru *LRUCache) addToFront(node *Node) {

}

// removeNode removes a node from the list
func (lru *LRUCache) removeNode(node *Node) {

}

// removeLRU removes the least recently used item from the cache
func (lru *LRUCache) removeLRU() {

}

// Size returns the current number of items in the cache
func (lru *LRUCache) Size() int {

}

// Clear removes all items from the cache
func (lru *LRUCache) Clear() {

}

// String returns a string representation of the cache
func (lru *LRUCache) String() string {
    var result string
    current := lru.head
    for current != nil {
        result += fmt.Sprintf("(%d:%d) ", current.key, current.value)
        current = current.next
    }
    return result
}
