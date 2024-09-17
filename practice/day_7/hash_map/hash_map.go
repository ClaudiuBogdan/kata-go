package hashmap

import (
	"errors"
	"fmt"
	"hash/fnv"
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
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) & (len(hm.buckets) - 1)
}

// Put adds a key-value pair to the hash map
func (hm *HashMap) Put(key string, value interface{}) {
	totalItems := float64(hm.size + 1)
	totalCapacity := float64(len(hm.buckets))
	currLoad := totalItems / totalCapacity

	if currLoad > loadFactor {
		hm.resize()
	}

	index := hm.hash(key)

	for i := 0; i < len(hm.buckets); i++ {
		// Get next index. The len is power of 2
		newIndex := (index + i) & (len(hm.buckets) - 1)
		if hm.buckets[newIndex] == nil {
			hm.buckets[newIndex] = &KeyValue{
				Key:   key,
				Value: value,
			}
			hm.size++
			return
		} else if hm.buckets[newIndex].Key == key {
			hm.buckets[newIndex].Value = value
			return
		}
	}
}

// Get retrieves the value associated with the given key
func (hm *HashMap) Get(key string) (interface{}, error) {
	if hm.IsEmpty() {
		return nil, errors.New("hash map is empty")
	}

	index := hm.hash(key)

	for i := 0; i < len(hm.buckets); i++ {
		newIndex := (index + i) & (len(hm.buckets) - 1)
		if hm.buckets[newIndex] == nil {
			return nil, errors.New("not found")
		} else if hm.buckets[newIndex].Key == key {
			return hm.buckets[newIndex].Value, nil
		}
	}
	return nil, errors.New("not found")
}

// Remove removes the key-value pair with the given key
func (hm *HashMap) Remove(key string) bool {
	if hm.IsEmpty() {
		return false
	}

	index := hm.hash(key)

	for i := 0; i < len(hm.buckets); i++ {
		newIndex := (index + i) & (len(hm.buckets) - 1)
		if hm.buckets[newIndex] == nil {
			// Item not found
			return false
		} else if hm.buckets[newIndex].Key == key {
			hm.buckets[newIndex] = nil
			hm.size--
			hm.rehash(newIndex)
			return true
		}
	}
	return false
}

// rehash rehashes the elements after the given index
func (hm *HashMap) rehash(startIndex int) {
	for i := (startIndex + 1) & (len(hm.buckets) - 1); hm.buckets[i] != nil; i = (i + 1) & (len(hm.buckets) - 1) {
		kv := hm.buckets[i]
		hm.buckets[i] = nil
		hm.Put(kv.Key, kv.Value)
	}
}

// resize increases the size of the hash map and rehashes all elements
func (hm *HashMap) resize() {
	oldBuckets := hm.buckets
	hm.buckets = make([]*KeyValue, len(hm.buckets)*2)
	hm.size = 0

	for _, item := range oldBuckets {
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
