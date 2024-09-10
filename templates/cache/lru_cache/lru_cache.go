package lrucache

import (
    "errors"
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
        cache:    make(map[int]*Node),
        head:     nil,
        tail:     nil,
    }
}

// Get retrieves a value from the cache
func (lru *LRUCache) Get(key int) (int, error) {
    if node, exists := lru.cache[key]; exists {
        lru.moveToFront(node)
        return node.value, nil
    }
    return 0, errors.New("key not found")
}

// Put adds a key-value pair to the cache
func (lru *LRUCache) Put(key, value int) {
    if node, exists := lru.cache[key]; exists {
        node.value = value
        lru.moveToFront(node)
    } else {
        newNode := &Node{key: key, value: value}
        if len(lru.cache) >= lru.capacity {
            lru.removeLRU()
        }
        lru.addToFront(newNode)
        lru.cache[key] = newNode
    }
}

// moveToFront moves a node to the front of the list
func (lru *LRUCache) moveToFront(node *Node) {
    if node == lru.head {
        return
    }
    lru.removeNode(node)
    lru.addToFront(node)
}

// addToFront adds a node to the front of the list
func (lru *LRUCache) addToFront(node *Node) {
    node.prev = nil
    node.next = lru.head
    if lru.head != nil {
        lru.head.prev = node
    }
    lru.head = node
    if lru.tail == nil {
        lru.tail = node
    }
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
    if lru.tail == nil {
        return
    }
    delete(lru.cache, lru.tail.key)
    lru.removeNode(lru.tail)
}

// Size returns the current number of items in the cache
func (lru *LRUCache) Size() int {
    return len(lru.cache)
}

// Clear removes all items from the cache
func (lru *LRUCache) Clear() {
    lru.cache = make(map[int]*Node)
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
