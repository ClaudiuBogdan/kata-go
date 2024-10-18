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
    return &LRUCache{
        capacity: capacity,
        cache: make(map[int]*Node, capacity),
        head: nil,
        tail: nil,
    }
}

// Get retrieves a value from the cache
func (lru *LRUCache) Get(key int) (int, error) {
    if lru.Size() == 0 {
        return 0, fmt.Errorf("cache is empty")
    }
    node, ok := lru.cache[key]

    if !ok {
        return 0, fmt.Errorf("key not found")
    }
    // TODO: move to front on get
    lru.moveToFront(node)
    return node.value, nil
}

// Put adds a key-value pair to the cache
func (lru *LRUCache) Put(key, value int) {
    if lru.capacity == 0 {
        return
    }
    if node, ok := lru.cache[key]; ok {
        node.value = value
        lru.moveToFront(node)
        return
    }

    if lru.Size() == lru.capacity {
        lru.removeLRU()
    }

    newNode := &Node{
        key: key,
        value: value,
        next: nil,
        prev: nil,
    }
    lru.addToFront(newNode)
    lru.cache[key] = newNode
}

// moveToFront moves a node to the front of the list
func (lru *LRUCache) moveToFront(node *Node) {
    // TODO: don't update if it is already the head. It will cause a infinite recursion.
    if lru.head == node {
        return
    }
    lru.removeNode(node)
    lru.addToFront(node)
}

// addToFront adds a node to the front of the list
func (lru *LRUCache) addToFront(node *Node) {
    if lru.head == nil {
        lru.head = node
        lru.tail = node
        return
    }

    lru.head.prev = node
    node.next = lru.head
    lru.head = node
}

// removeNode removes a node from the list
func (lru *LRUCache) removeNode(node *Node) {
    if node.prev != nil {
        node.prev.next = node.next
    } else {
        lru.head = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    } else {
        lru.tail = node.prev
    }
}

// removeLRU removes the least recently used item from the cache
func (lru *LRUCache) removeLRU() {
    tail := lru.tail
    if tail == nil{
        return
    }
    lru.removeNode(tail)
    // TODO: delete from map
    delete(lru.cache, tail.key)
}

// Size returns the current number of items in the cache
func (lru *LRUCache) Size() int {
    return len(lru.cache)
}

// Clear removes all items from the cache
func (lru *LRUCache) Clear() {
    lru.cache = map[int]*Node{}
    lru.head = nil
    lru.tail = nil
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
