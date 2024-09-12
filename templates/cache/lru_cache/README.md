# LRU Cache

## Problem Description

Implement a data structure for a Least Recently Used (LRU) cache. The cache should support fast access, insertion, and deletion operations.

## Approach: Hash Table + Doubly Linked List

The LRU cache is implemented using a combination of a hash table and a doubly linked list:

1. Hash Table: Provides O(1) access to cache items
2. Doubly Linked List: Maintains the order of items based on their recent use

## Implementation Details

The LRU cache should support the following operations:

- `Get(key)`: Get the value associated with the key if it exists in the cache, or return -1 if it doesn't. This operation should also mark the key as recently used.
- `Put(key, value)`: Insert a key-value pair into the cache. If the key already exists, update its value. If the cache is at capacity, remove the least recently used item before inserting the new item.

## Complexity Analysis

- Time Complexity: O(1) for both Get and Put operations
- Space Complexity: O(capacity) to store up to 'capacity' number of items in the cache

## Usage

```go
cache := Constructor(2)  // Create a cache with capacity of 2

cache.Put(1, 1)
cache.Put(2, 2)
fmt.Println(cache.Get(1))  // Output: 1
cache.Put(3, 3)            // This will evict key 2
fmt.Println(cache.Get(2))  // Output: -1 (not found)
cache.Put(4, 4)            // This will evict key 1
fmt.Println(cache.Get(1))  // Output: -1 (not found)
fmt.Println(cache.Get(3))  // Output: 3
fmt.Println(cache.Get(4))  // Output: 4
```

## Testing

The implementation should include a comprehensive test suite covering various scenarios:

1. Basic functionality of Get and Put
2. Handling capacity limits
3. Updating existing keys
4. Eviction of least recently used items
5. Edge cases (e.g., empty cache, accessing non-existent keys)

To run the tests, use the following command:

```bash
go test
```

## Advantages of the Hash Table + Doubly Linked List Approach

1. Efficiency: O(1) time complexity for all operations
2. Flexible capacity: Easy to implement with different cache sizes
3. Clear eviction policy: Least recently used item is always at the tail of the list
4. Fast updates: Moving an item to the front of the list is O(1) with a doubly linked list

## Applications

- Database caching
- Web server caching
- File system caching
- Image processing and computer graphics
- Network routing tables

## Visual Representation

Here's a visual representation of the LRU Cache structure:

```
Hash Table:
key1 -> Node1
key2 -> Node2
key3 -> Node3

Doubly Linked List:
[HEAD] <-> [Node1] <-> [Node2] <-> [Node3] <-> [TAIL]
(most recently used)          (least recently used)
```

## Variations and Extensions

1. LFU Cache: Least Frequently Used cache
2. Time-based cache expiration
3. Distributed LRU Cache
4. Write-through and write-back policies
5. LRU Cache with size-aware eviction (where items have different sizes)

## Implementation Notes

- The implementation should use a combination of a hash map (for fast lookups) and a doubly linked list (for fast reordering).
- Be careful with edge cases, such as accessing or updating an item that's already the most recently used.
- Consider thread safety if the cache will be used in a concurrent environment.

## Limitations

- Fixed capacity: Once set, changing the capacity is not trivial
- No built-in persistence: Data is lost when the program terminates
- Potential for cache pollution with one-time-use data

## Related Problems

- First Unique Character in a String
- LFU Cache
- Design In-Memory File System
- Design Search Autocomplete System
- Min Stack

Remember that while this implementation is efficient for most use cases, there might be scenarios where other caching strategies or data structures are more appropriate. Always consider the specific requirements and constraints of your system when choosing a caching solution.
