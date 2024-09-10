package huffmancoding

import (
	"container/heap"
	"fmt"
	"strings"
)

// Node represents a node in the Huffman tree
type Node struct {
	Char  rune
	Freq  int
	Left  *Node
	Right *Node
}

// PriorityQueue implements heap.Interface and holds Nodes
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Freq < pq[j].Freq }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// BuildHuffmanTree builds the Huffman tree from the given frequency map
func BuildHuffmanTree(freqMap map[rune]int) *Node {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for char, freq := range freqMap {
		heap.Push(&pq, &Node{Char: char, Freq: freq})
	}

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*Node)
		right := heap.Pop(&pq).(*Node)
		parent := &Node{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}
		heap.Push(&pq, parent)
	}

	return heap.Pop(&pq).(*Node)
}

// GenerateHuffmanCodes generates Huffman codes for each character
func GenerateHuffmanCodes(root *Node) map[rune]string {
	codes := make(map[rune]string)
	generateCodesRecursive(root, "", codes)
	return codes
}

func generateCodesRecursive(node *Node, code string, codes map[rune]string) {
	if node == nil {
		return
	}
	if node.Char != 0 {
		codes[node.Char] = code
	}
	generateCodesRecursive(node.Left, code+"0", codes)
	generateCodesRecursive(node.Right, code+"1", codes)
}

// Compress compresses the input string using Huffman coding
func Compress(input string) (string, map[rune]string) {
	if len(input) == 0 {
		return "", nil
	}

	freqMap := make(map[rune]int)
	for _, char := range input {
		freqMap[char]++
	}

	root := BuildHuffmanTree(freqMap)
	codes := GenerateHuffmanCodes(root)

	var compressed strings.Builder
	for _, char := range input {
		compressed.WriteString(codes[char])
	}

	return compressed.String(), codes
}

// Decompress decompresses the input string using the Huffman codes
func Decompress(compressed string, codes map[rune]string) (string, error) {
	if len(compressed) == 0 {
		return "", nil
	}

	reverseMap := make(map[string]rune)
	for char, code := range codes {
		reverseMap[code] = char
	}

	var decompressed strings.Builder
	currentCode := ""
	for _, bit := range compressed {
		currentCode += string(bit)
		if char, found := reverseMap[currentCode]; found {
			decompressed.WriteRune(char)
			currentCode = ""
		}
	}

	if currentCode != "" {
		return "", fmt.Errorf("invalid compressed data")
	}

	return decompressed.String(), nil
}
