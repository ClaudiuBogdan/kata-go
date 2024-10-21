package huffmancoding

import (
	"sort"
)

// Node represents a node in the Huffman tree
type Node struct {
	Char  rune
	Freq  int
	Left  *Node
	Right *Node
}

// buildFrequencyMap creates a map of characters to their frequencies in the input text
func buildFrequencyMap(text string) map[rune]int {

}

// buildHuffmanTree constructs the Huffman tree based on character frequencies
func buildHuffmanTree(freqMap map[rune]int) *Node {

}

// generateCodes traverses the Huffman tree to generate codes for each character
func generateCodes(root *Node) map[rune]string {

}

// encode converts the input text to its Huffman-encoded representation
func encode(text string, codes map[rune]string) string {

}

// decode converts the Huffman-encoded string back to the original text
func decode(encoded string, root *Node) string {

}

// HuffmanCoding performs the complete Huffman coding process on the input text
func HuffmanCoding(text string) (string, string, map[rune]string) {

}

// PrintHuffmanCodes returns a sorted list of character-code pairs for display
func PrintHuffmanCodes(codes map[rune]string) []string {
	chars := make([]rune, 0, len(codes))
	for char := range codes {
		chars = append(chars, char)
	}
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

	result := make([]string, len(chars))
	for i, char := range chars {
		result[i] = string(char) + ": " + codes[char]
	}
	return result
}
