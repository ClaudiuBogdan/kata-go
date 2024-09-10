# Huffman Coding for Data Compression

## Problem Description

Implement Huffman coding, a data compression algorithm that assigns variable-length codes to characters based on their frequency of occurrence. More frequent characters are assigned shorter codes, resulting in overall compression of the data.

## Approach

The implementation follows these steps:

1. Count the frequency of each character in the input string.
2. Build a Huffman tree using a priority queue (min-heap):
   - Create a leaf node for each character and add it to the priority queue.
   - Repeatedly extract two nodes with the lowest frequency, create a new internal node with these two as children, and add it back to the queue.
   - Continue until only one node remains (the root of the Huffman tree).
3. Generate Huffman codes by traversing the tree:
   - Assign '0' for left branches and '1' for right branches.
   - The code for each character is the path from the root to the leaf node.
4. Compress the input string by replacing each character with its Huffman code.
5. For decompression, use the Huffman codes to reconstruct the original string.

## Complexity Analysis

- Time Complexity: 
  - Building the frequency map: O(n)
  - Building the Huffman tree: O(k log k), where k is the number of unique characters
  - Generating codes: O(k)
  - Compression and decompression: O(n)
  Overall: O(n + k log k), where n is the length of the input string
- Space Complexity: O(k) for the Huffman tree and codes, where k is the number of unique characters

## Usage

```go
input := "hello world"
compressed, codes := Compress(input)
fmt.Printf("Compressed: %s\n", compressed)

decompressed, err := Decompress(compressed, codes)
if err != nil {
    fmt.Printf("Decompression error: %v\n", err)
} else {
    fmt.Printf("Decompressed: %s\n", decompressed)
}
```

## Implementation Details

The package provides the following main components:

1. `Node` struct: Represents a node in the Huffman tree.
2. `PriorityQueue` type: A min-heap implementation for building the Huffman tree.
3. `BuildHuffmanTree(freqMap map[rune]int) *Node`: Builds the Huffman tree from a frequency map.
4. `GenerateHuffmanCodes(root *Node) map[rune]string`: Generates Huffman codes from the tree.
5. `Compress(input string) (string, map[rune]string)`: Compresses the input string and returns the compressed string and Huffman codes.
6. `Decompress(compressed string, codes map[rune]string) (string, error)`: Decompresses the input string using the provided Huffman codes.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Compressing and decompressing simple strings
2. Handling repeated characters
3. Edge cases like single-character strings and empty strings
4. Long strings with varied character frequencies
5. Invalid compressed data

To run the tests, use the following command:

```bash
go test
```

## Advantages and Limitations

Advantages:
- Provides efficient compression for data with varying character frequencies
- Lossless compression (original data can be fully recovered)
- Can be adapted for different types of data (not just text)

Limitations:
- Requires the Huffman tree or codes to be transmitted along with the compressed data
- May not be efficient for small amounts of data or data with uniform character frequencies
- Not suitable for real-time compression of streaming data (entire input needed to build the tree)

## Applications

- Text compression
- Image and video compression (as part of more complex algorithms)
- Data transmission in networks
- Archiving and file compression
- Cryptography (though not for security, but as part of larger cryptographic systems)

Remember that while Huffman coding is a fundamental compression technique, modern compression algorithms often use more advanced methods or combinations of techniques for better performance in specific scenarios.

