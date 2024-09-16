# Hash Map - Notes

## That's an excellent question. Let's compare the linear probing approach used in our HashMap implementation to a linked list-based approach (often called separate chaining). Both are methods for handling collisions in hash tables, but they have different characteristics

## Linear Probing vs. Linked List for Handling Collisions in Hash Tables

When implementing a hash table, one of the key design decisions is how to handle collisions, which occur when two keys hash to the same index. Two common approaches are linear probing and linked list-based chaining.

Let's break down the comparison:

1. Structure:
   - Linear Probing: Uses a single array. When a collision occurs, it looks for the next empty slot in the array.
   - Linked List: Each bucket in the array contains a linked list. When a collision occurs, the new item is added to the linked list at that bucket.

2. Memory Usage:
   - Linear Probing: Generally uses less memory as it only needs to store the elements themselves.
   - Linked List: Requires additional memory for the linked list nodes (next pointers).

3. Performance:

   a. Insertion:
   - Linear Probing: O(1) average case, but can degrade to O(n) in worst case (when many collisions occur).
   - Linked List: O(1) to insert at the head of the list, regardless of collisions.

   b. Lookup:
   - Linear Probing: O(1) average case, O(n) worst case.
   - Linked List: O(1) to find the bucket, then O(k) to search the list, where k is the length of the list at that bucket.

   c. Deletion:
   - Linear Probing: More complex, as you need to rehash elements that were placed after the deleted item.
   - Linked List: Simpler, just remove the node from the list.

4. Cache Performance:
   - Linear Probing: Better cache performance due to array locality.
   - Linked List: Worse cache performance due to scattered memory access.

5. Load Factor:
   - Linear Probing: Performance degrades rapidly as load factor increases.
   - Linked List: Can maintain good performance even with a high load factor.

6. Implementation Complexity:
   - Linear Probing: Simpler to implement, but requires careful handling of deletions.
   - Linked List: Slightly more complex due to linked list management, but easier to handle deletions.

In conclusion, both approaches have their strengths:

- Linear probing (our implementation) is generally faster for small to medium-sized hash maps due to better cache performance and simpler code.
- Linked list-based approach handles high load factors better and simplifies deletion operations.

The choice between these methods often depends on the specific use case, expected data distribution, and performance requirements. Our linear probing implementation is a good general-purpose solution, but for scenarios with a high number of collisions or frequent deletions, a linked list-based approach might be preferable.

## newIndex := (index + i) & (len(hm.buckets) - 1)

This line is a key part of the linear probing mechanism in the hash map implementation. Let's break it down:

```go
newIndex := (index + i) & (len(hm.buckets) - 1)
```

This line calculates the next index to check when searching for an empty slot or a matching key. It's doing two important things:

1. Linear probing: `(index + i)`
   - `index` is the original hash of the key
   - `i` is the current step in the probing sequence (0, 1, 2, ...)
   - This part implements the linear probing by moving `i` steps from the original index

2. Wrapping around: `& (len(hm.buckets) - 1)`
   - This is a bitwise AND operation
   - It's used to wrap the index around to the beginning of the array if it exceeds the array's length

The reason this works is due to a property of binary numbers:

- When the length of `hm.buckets` is a power of 2 (which is typically ensured in hash map implementations), `len(hm.buckets) - 1` will be a number with all bits set to 1.
- For example, if `len(hm.buckets)` is 8 (2^3), then `len(hm.buckets) - 1` is 7, which in binary is 111.

Let's see an example:

Assume `len(hm.buckets)` is 8:

- `len(hm.buckets) - 1` = 7 (in binary: 111)

Now, let's say `index + i` is 10 (in binary: 1010)

```
  1010 (10 in decimal)
& 0111 (7 in decimal, which is len(hm.buckets) - 1)
  ----
  0010 (2 in decimal)
```

So, `10 & 7 = 2`, which means the 10th slot wraps around to the 2nd slot in our 8-slot array.

This operation is equivalent to `(index + i) % len(hm.buckets)`, but it's much faster because bitwise operations are typically more efficient than modulo operations in most processors.

The key benefits of this approach are:

1. It ensures the index always stays within the bounds of the array.
2. It provides an efficient way to wrap around to the beginning of the array when we reach the end.
3. It's very fast, as bitwise operations are typically single-cycle operations on most CPUs.

This technique is a common optimization in hash table implementations, providing both safety (preventing out-of-bounds access) and performance (fast computation of the wrapped index).
