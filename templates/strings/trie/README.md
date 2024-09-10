# Trie (Prefix Tree) Implementation

## Data Structure Description

A Trie, also called prefix tree, is an efficient information retrieval data structure. It is commonly used to store and search strings in a space and time efficient way. The trie stores a dynamic set or associative array where the keys are usually strings.

## Implementation Details

This implementation provides three main operations:

1. `Insert(word string)`: Adds a word to the trie.
2. `Search(word string) bool`: Returns true if the word is in the trie, false otherwise.
3. `StartsWith(prefix string) bool`: Returns true if there is any word in the trie that starts with the given prefix.

### Time Complexity
- Insert: O(m), where m is the length of the word
- Search: O(m), where m is the length of the word
- StartsWith: O(m), where m is the length of the prefix

### Space Complexity
- O(n*m), where n is the number of words and m is the average length of words

## Usage

```go
trie := NewTrie()
trie.Insert("apple")
fmt.Println(trie.Search("apple"))    // Output: true
fmt.Println(trie.Search("app"))      // Output: false
fmt.Println(trie.StartsWith("app"))  // Output: true
```

## Testing

The implementation includes a comprehensive test suite covering various scenarios. To run the tests:

```bash
go test
```

## Implementation Notes

- The trie uses a map for children nodes, allowing for efficient lookup and insertion of any Unicode character.
- Each node in the trie contains a boolean flag `isEnd` to mark the end of a word.
- The `StartsWith` method can check for prefixes efficiently.

## Advantages of Trie

1. Efficient prefix-based operations: Searching and prefix matching are very fast.
2. Space-efficient for common prefixes: Saves space when storing many strings with similar prefixes.
3. Predictable performance: O(m) time complexity for basic operations, where m is the key length.

## Limitations

1. Can use more space than other data structures when storing many words with few common prefixes.
2. Not as efficient for exact string matching compared to hash tables.

## When to Use a Trie

- When you need fast prefix-based retrieval, such as in autocomplete systems.
- In spell checkers or dictionaries where prefix-based searches are common.
- For implementing dictionaries with efficient key lookup.
- In IP routing tables, where the trie can efficiently store and retrieve IP addresses.

Remember, while tries are very efficient for certain operations, they may not always be the best choice depending on your specific use case. Consider factors like the nature of your data, the types of operations you'll be performing most frequently, and your memory constraints when deciding whether to use a trie.
