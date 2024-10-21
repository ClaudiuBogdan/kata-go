# Huffman Coding

## Problem Description

Huffman coding is a data compression technique that assigns variable-length codes to characters based on their frequency of occurrence. The goal is to use shorter codes for more frequent characters and longer codes for less frequent ones, resulting in overall data compression.

## Approach: Greedy Algorithm and Priority Queue

Huffman coding uses a greedy algorithm approach combined with a priority queue:

1. Calculate the frequency of each character in the input data.
2. Create a leaf node for each character and add it to a priority queue.
3. While there is more than one node in the queue:
   - Remove the two nodes with the lowest frequency.
   - Create a new internal node with these two nodes as children.
   - Add the new node to the queue.
4. The remaining node is the root of the Huffman tree.
5. Traverse the tree to assign codes (0 for left, 1 for right) to each character.

## Complexity Analysis

- Time Complexity: O(n log n), where n is the number of unique characters. Building the priority queue takes O(n), and we perform n-1 extract-min operations, each taking O(log n).
- Space Complexity: O(n) for storing the Huffman tree and the character frequency map.

## Implementation Details

The package provides the following main functions:

`BuildHuffmanTree(text string) *Node`: Builds a Huffman tree from the input text.
`GenerateHuffmanCodes(root *Node) map[rune]string`: Generates Huffman codes for each character.
`Encode(text string, codes map[rune]string) string`: Encodes the input text using Huffman codes.
`Decode(encoded string, root *Node) string`: Decodes the Huffman-encoded text.

## Usage

```go
text := "hello world"
root := BuildHuffmanTree(text)
codes := GenerateHuffmanCodes(root)
encoded := Encode(text, codes)
decoded := Decode(encoded, root)
fmt.Println(encoded)
fmt.Println(decoded)
```

## Testing

The implementation should include a comprehensive test suite covering various scenarios:

1. Encoding and decoding simple strings
2. Handling empty strings
3. Texts with varying character frequencies
4. Large inputs
5. Special characters and Unicode

To run the tests, use the following command:

```bash
go test
```

## Advantages of Huffman Coding

1. Optimal prefix-free encoding
2. Lossless compression
3. Adaptive to input data characteristics
4. Relatively simple implementation

## Applications

- Text compression
- Image and video compression (as part of larger algorithms)
- Data transmission in communication systems
- File compression utilities

## Visual Representation

Here's a visual representation of a Huffman tree for the string "hello world":

```
        *
      /   \
     *     o
   /   \
  *     l
 / \
h   *
   / \
  e   *
     / \
    w   r
        |
        d
```

## Variations and Extensions

1. Adaptive Huffman coding
2. Canonical Huffman coding
3. Huffman coding with limited code length
4. Huffman coding for streams of data

## Implementation Notes

- The implementation assumes that the input is valid (non-empty string).
- This solution provides both encoding and decoding functionality.
- For very large inputs, consider implementing a streaming version of the algorithm.

## Limitations

- Not efficient for small amounts of data
- Requires knowledge of character frequencies in advance
- Overhead of storing the Huffman tree or codes

## Related Problems

- Run-Length Encoding
- Arithmetic Coding
- LZW Compression
- Burrows-Wheeler Transform
- Shannon-Fano Coding
