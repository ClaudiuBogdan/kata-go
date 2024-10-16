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
	return &HashMap{
		buckets: make([]*KeyValue, defaultCapacity),
		size:    0,
	}
}

// hash generates a hash code for the given key
func (hm *HashMap) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = 31*hash + int(key[i])
	}
	return hash & (len(hm.buckets) - 1)
}

// Put adds a key-value pair to the hash map
func (hm *HashMap) Put(key string, value interface{}) {
	currentLoad := float64(hm.size + 1)
	currentCapacity := float64(len(hm.buckets))
	currentLoadFactor := currentLoad / currentCapacity
	if currentLoadFactor > loadFactor {
		hm.resize()
	}

	index := hm.hash(key)

	for i := 0; i < len(hm.buckets); i++ {
		nextIndex := (index + i) & (len(hm.buckets) - 1)

		if hm.buckets[nextIndex] == nil {
			hm.buckets[nextIndex] = &KeyValue{
				Key:   key,
				Value: value,
			}
			hm.size++
			return
		} else if hm.buckets[nextIndex].Key == key {
			hm.buckets[nextIndex].Value = value
			return
		}
	}
}

// Get retrieves the value associated with the given key
func (hm *HashMap) Get(key string) (interface{}, error) {
	var zero interface{}
	if hm.IsEmpty() {
		return zero, fmt.Errorf("hashmap is empty")
	}

	index := hm.hash(key)

	for i := 0; i < len(hm.buckets); i++ {
		nextIndex := (index + i) & (len(hm.buckets) - 1)

		if hm.buckets[nextIndex] == nil {
			return zero, fmt.Errorf("key not found")
		} else if hm.buckets[nextIndex].Key == key {
			return hm.buckets[nextIndex].Value, nil
		}
	}

	return zero, fmt.Errorf("key not found")
}

// Remove removes the key-value pair with the given key
func (hm *HashMap) Remove(key string) bool {
	index := hm.hash(key)

	for i := 0; i < len(hm.buckets); i++ {
		nextIndex := (index + i) & (len(hm.buckets) - 1)
		if hm.buckets[nextIndex] == nil {
			return false
		} else if hm.buckets[nextIndex].Key == key {
			hm.buckets[nextIndex] = nil
			hm.rehash(nextIndex)
			hm.size--
			return true
		}
	}
	return false
}

// rehash rehashes the elements after the given index
func (hm *HashMap) rehash(startIndex int) {
	for i := 1; i < len(hm.buckets); i++ {
		nextIndex := (startIndex + i) & (len(hm.buckets) - 1)

		if hm.buckets[nextIndex] == nil {
			return
		} else {
			pair := hm.buckets[nextIndex]
			hm.buckets[nextIndex] = nil
			hm.size--
			hm.Put(pair.Key, pair.Value)
		}
	}
}

// resize increases the size of the hash map and rehashes all elements
func (hm *HashMap) resize() {
	newCapacity := len(hm.buckets) * 2
	oldBuckets := hm.buckets

	hm.buckets = make([]*KeyValue, newCapacity)
	hm.size = 0

	for i := 0; i < len(oldBuckets); i++ {
		item := oldBuckets[i]
		if item != nil {
			hm.Put(item.Key, item.Value)
		}
	}
}

// Size returns the number of key-value pairs in the hash map
func (hm *HashMap) Size() int {
	return hm.size
}

// IsEmpty returns true if the hash map is empty
func (hm *HashMap) IsEmpty() bool {
	return hm.size == 0
}

// Clear removes all key-value pairs from the hash map
func (hm *HashMap) Clear() {
	hm.buckets = make([]*KeyValue, defaultCapacity)
	hm.size = 0
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
