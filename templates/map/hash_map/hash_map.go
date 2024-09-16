package hashmap

import (
	"fmt"
)

const (
	defaultCapacity = 16
	loadFactor      = 0.75
)

// KeyValue represents a key-value pair in the hash map
type KeyValue struct {
	Key   string
	Value interface{}
}

// HashMap represents a hash map data structure
type HashMap struct {
	buckets []*KeyValue
	size    int
}

// NewHashMap creates a new hash map
func NewHashMap() *HashMap {

}

// hash generates a hash code for the given key
func (hm *HashMap) hash(key string) int {
    
}

// Put adds a key-value pair to the hash map
func (hm *HashMap) Put(key string, value interface{}) {
	
}

// Get retrieves the value associated with the given key
func (hm *HashMap) Get(key string) (interface{}, error) {
	
}

// Remove removes the key-value pair with the given key
func (hm *HashMap) Remove(key string) bool {
	
}

// rehash rehashes the elements after the given index
func (hm *HashMap) rehash(startIndex int) {
	
}

// resize increases the size of the hash map and rehashes all elements
func (hm *HashMap) resize() {
	
}

// Size returns the number of key-value pairs in the hash map
func (hm *HashMap) Size() int {

}

// IsEmpty returns true if the hash map is empty
func (hm *HashMap) IsEmpty() bool {

}

// Clear removes all key-value pairs from the hash map
func (hm *HashMap) Clear() {
    
}

// String returns a string representation of the hash map
func (hm *HashMap) String() string {
	var result string
	for _, kv := range hm.buckets {
		if kv != nil {
			result += fmt.Sprintf("(%s:%v) ", kv.Key, kv.Value)
		}
	}
	return result
}