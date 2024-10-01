package hashmap

import (
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
	// TODO: learn this hash algorithm. Maybe add a new kata as the has algorithm
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) & (len(hm.buckets) - 1)
}

// Put adds a key-value pair to the hash map
func (hm *HashMap) Put(key string, value interface{}) {
	currLoadFactor := float64(hm.size+1) / float64(len(hm.buckets))
	if currLoadFactor > loadFactor {
		hm.resize()
	}

	index := hm.hash(key)

	for i := 0; i < len(hm.buckets); i++ {
		newIndex := (index + i) & (len(hm.buckets) - 1)

		if hm.buckets[newIndex] == nil {
			hm.buckets[newIndex] = &KeyValue{
				Key:   key,
				Value: value,
			}
			hm.size++
			break
		} else if hm.buckets[newIndex].Key == key {
			hm.buckets[newIndex].Value = value
			break
		}
	}
}

// Get retrieves the value associated with the given key
func (hm *HashMap) Get(key string) (interface{}, error) {
	if hm.IsEmpty() {
		var zero interface{}
		return zero, fmt.Errorf("hashmap is empty")
	}

	index := hm.hash(key)

	for i := 0; i < len(hm.buckets); i++ {
		newIndex := (index + i) & (len(hm.buckets) - 1)

		if hm.buckets[newIndex] == nil {
			break
		}

		if hm.buckets[newIndex].Key == key {
			return hm.buckets[newIndex].Value, nil
		}
	}

	return nil, fmt.Errorf("key not found")
}

// Remove removes the key-value pair with the given key
func (hm *HashMap) Remove(key string) bool {

	index := hm.hash(key)

	for i := 0; i < len(hm.buckets); i++ {
		newIndex := (index + i) & (len(hm.buckets) - 1)

		if hm.buckets[newIndex] == nil {
			return false
		}

		if hm.buckets[newIndex].Key == key {
			hm.buckets[newIndex] = nil
			hm.rehash(newIndex)
			hm.size--
			return true
		}
	}

	return false
}

// rehash rehashes the elements after the given index
func (hm *HashMap) rehash(startIndex int) {

	for i := 0; i < len(hm.buckets); i++ {
		newIndex := (startIndex + i) & (len(hm.buckets) - 1)
		nextIndex := (newIndex + 1) & (len(hm.buckets) - 1)

		if hm.buckets[nextIndex] == nil {
			break
		}

		hm.buckets[newIndex] = hm.buckets[nextIndex]
		hm.buckets[nextIndex] = nil
	}
}

// resize increases the size of the hash map and rehashes all elements
func (hm *HashMap) resize() {
	oldBucket := hm.buckets
	hm.buckets = make([]*KeyValue, len(hm.buckets)*2)
	hm.size = 0

	for i := 0; i < len(oldBucket); i++ {
		item := oldBucket[i]
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
