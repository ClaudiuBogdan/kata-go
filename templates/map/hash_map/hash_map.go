package hashmap

import (
    "errors"
    "fmt"
)

const (
    defaultCapacity = 16
    loadFactor      = 0.75
)

// KeyValue represents a key-value pair in the hash map
type KeyValue struct {
    Key   string
    Value int
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
func (hm *HashMap) Put(key string, value int) {
    if float64(hm.size+1)/float64(len(hm.buckets)) > loadFactor {
        hm.resize()
    }

    index := hm.hash(key)
    if hm.buckets[index] == nil {
        hm.buckets[index] = &KeyValue{Key: key, Value: value}
        hm.size++
    } else if hm.buckets[index].Key == key {
        hm.buckets[index].Value = value
    } else {
        // Collision: Linear probing
        for i := 1; i < len(hm.buckets); i++ {
            newIndex := (index + i) & (len(hm.buckets) - 1)
            if hm.buckets[newIndex] == nil {
                hm.buckets[newIndex] = &KeyValue{Key: key, Value: value}
                hm.size++
                break
            } else if hm.buckets[newIndex].Key == key {
                hm.buckets[newIndex].Value = value
                break
            }
        }
    }
}

// Get retrieves the value associated with the given key
func (hm *HashMap) Get(key string) (int, error) {
    index := hm.hash(key)
    if hm.buckets[index] != nil && hm.buckets[index].Key == key {
        return hm.buckets[index].Value, nil
    }

    // Collision: Linear probing
    for i := 1; i < len(hm.buckets); i++ {
        newIndex := (index + i) & (len(hm.buckets) - 1)
        if hm.buckets[newIndex] == nil {
            break
        }
        if hm.buckets[newIndex].Key == key {
            return hm.buckets[newIndex].Value, nil
        }
    }

    return 0, errors.New("key not found")
}

// Remove removes the key-value pair with the given key
func (hm *HashMap) Remove(key string) bool {
    index := hm.hash(key)
    if hm.buckets[index] != nil && hm.buckets[index].Key == key {
        hm.buckets[index] = nil
        hm.size--
        return true
    }

    // Collision: Linear probing
    for i := 1; i < len(hm.buckets); i++ {
        newIndex := (index + i) & (len(hm.buckets) - 1)
        if hm.buckets[newIndex] == nil {
            break
        }
        if hm.buckets[newIndex].Key == key {
            hm.buckets[newIndex] = nil
            hm.size--
            return true
        }
    }

    return false
}

// resize increases the size of the hash map and rehashes all elements
func (hm *HashMap) resize() {
    newBuckets := make([]*KeyValue, len(hm.buckets)*2)
    for _, kv := range hm.buckets {
        if kv != nil {
            index := hm.hash(kv.Key)
            for newBuckets[index] != nil {
                index = (index + 1) & (len(newBuckets) - 1)
            }
            newBuckets[index] = kv
        }
    }
    hm.buckets = newBuckets
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
            result += fmt.Sprintf("(%s:%d) ", kv.Key, kv.Value)
        }
    }
    return result
}
