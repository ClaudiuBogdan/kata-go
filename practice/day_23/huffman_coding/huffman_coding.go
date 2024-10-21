package huffmancoding

import (
	"container/heap"
	"sort"
	"strings"
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
	freq := make(map[rune]int, len(text))

	for _, char := range text {
		freq[char]++
	}

	return freq
}

// buildHuffmanTree constructs the Huffman tree based on character frequencies
func buildHuffmanTree(freqMap map[rune]int) *Node {
	pq := make(PriorityQueue, 0, len(freqMap))

	for char, freq := range freqMap {
		pq = append(pq, &Node{Char: char, Freq: freq})
	}

	heap.Init(&pq)
	
	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right :=heap.Pop(&pq).(*Node)

		parent := &Node{Freq: left.Freq + right.Freq}

		if left.Char < right.Char || (left.Char == 0 && right.Char != 0){
			parent.Left = left
			parent.Right = right
		} else {
			parent.Left = right
			parent.Right = left
		}
		heap.Push(&pq, parent)
	}

	if pq.Len() == 0 {
		return nil
	}

	return heap.Pop(&pq).(*Node)
}

// generateCodes traverses the Huffman tree to generate codes for each character
func generateCodes(root *Node) map[rune]string {
	codes := make(map[rune]string)
	var generate func(*Node, string)
	generate = func(node *Node, code string){
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			codes[node.Char] = code
			return
		}
		generate(node.Left, code + "0")
		generate(node.Right, code + "1")
	}

	generate(root, "")
	return codes
}

// encode converts the input text to its Huffman-encoded representation
func encode(text string, codes map[rune]string) string {
	var encoded strings.Builder
	for _, char := range text {
		encoded.WriteString(codes[char])
	}
	return encoded.String()
}

// decode converts the Huffman-encoded string back to the original text
func decode(encoded string, root *Node) string {
	if root == nil {
		return ""
	}

	if root.Left == nil && root.Right == nil {
		// Special case: single character
		return strings.Repeat(string(root.Char),len(encoded))
	}
	var decoded strings.Builder
	current := root
	for _, bit := range encoded {
		if bit == '0'{
			current = current.Left
		} else {
			current = current.Right
		}
		if current.Left == nil && current.Right == nil {
			decoded.WriteRune(current.Char)
			current = root
		}
	}

	return decoded.String()
}

// HuffmanCoding performs the complete Huffman coding process on the input text
func HuffmanCoding(text string) (string, string, map[rune]string) {
	if text == ""{
		return "", "", nil
	}

	freqMap := buildFrequencyMap(text)

	// Special case
	if len(freqMap) == 1 {
		char := rune(text[0])
		return strings.Repeat("0", len(text)), text, map[rune]string{char: "0"}
	}

	root := buildHuffmanTree(freqMap)
	codes := generateCodes(root)
	encoded := encode(text, codes)
	decoded := decode(encoded, root)

	return encoded, decoded, codes
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
