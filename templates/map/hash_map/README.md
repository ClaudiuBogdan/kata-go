# Improved HashMap Implementation in Go

## Problem Description

Implement an efficient and flexible HashMap data structure in Go, capable of storing key-value pairs with string keys and values of any type. The implementation should handle collisions, dynamic resizing, and provide basic operations like Put, Get, Remove, and Clear.

## Videos

[Hash Table Data Structure](https://www.youtube.com/watch?v=shs0KM3wKv8)

## Approach

This HashMap implementation uses the following approach:

1. Use an array of buckets to store key-value pairs.
2. Employ the FNV-1a hash function for efficient key hashing.
3. Handle collisions using linear probing.
4. Implement dynamic resizing to maintain performance as the HashMap grows.
5. Provide methods for basic operations: Put, Get, Remove, Clear, Size, and IsEmpty.

## Complexity Analysis

- Time Complexity:
  - Average case for Put, Get, Remove: O(1)
  - Worst case (rare): O(n) due to linear probing in case of many collisions
- Space Complexity: O(n), where n is the number of key-value pairs stored

## Usage

```go
hm := NewHashMap()
hm.Put("key1", 42)
hm.Put("key2", "value")

value, err := hm.Get("key1")
if err == nil {
    fmt.Printf("Value for key1: %v\n", value)
}

hm.Remove("key2")
fmt.Printf("Size of HashMap: %d\n", hm.Size())
```

## Implementation Details

The package provides the following main components:

1. `KeyValue` struct: Represents a key-value pair in the HashMap.
2. `HashMap` struct: Represents the HashMap data structure.
3. `NewHashMap() *HashMap`: Creates a new HashMap.
4. `Put(key string, value interface{})`: Adds or updates a key-value pair.
5. `Get(key string) (interface{}, error)`: Retrieves the value for a given key.
6. `Remove(key string) bool`: Removes a key-value pair from the HashMap.
7. `Size() int`: Returns the number of key-value pairs in the HashMap.
8. `IsEmpty() bool`: Checks if the HashMap is empty.
9. `Clear()`: Removes all key-value pairs from the HashMap.
10. `String() string`: Returns a string representation of the HashMap.

The implementation uses the FNV-1a hash function and handles collisions with linear probing. It also includes dynamic resizing to maintain performance as the HashMap grows.

## Testing

To ensure the correctness of the HashMap implementation, create a comprehensive test suite that covers:

1. Basic operations (Put, Get, Remove)
2. Handling of collisions
3. Dynamic resizing
4. Edge cases (empty HashMap, non-existent keys)
5. Different value types

To run the tests, use the following command:

```bash
go test
```

## Advantages and Limitations

Advantages:

- Efficient average-case time complexity for basic operations
- Flexible value type (can store any type of value)
- Handles collisions and dynamic resizing

Limitations:

- Uses more memory than a fixed-size hash table
- Performance can degrade with a high load factor or poor hash distribution

## Applications

- Caching systems
- Database indexing
- Symbol tables in compilers and interpreters
- Implementing associative arrays
- Counting frequency of items

Remember that while this HashMap implementation is efficient for many use cases, there are specialized implementations (like Go's built-in `map`) that may be more optimized for specific scenarios.

## Further Reading

[Hash Table](https://www.youtube.com/watch?v=KyUTuwz_b7Q)
