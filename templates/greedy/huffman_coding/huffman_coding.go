package huffmancoding

import (
	"sort"
)

type Node struct {
	Char  rune
	Freq  int
	Left  *Node
	Right *Node
}

func buildFrequencyMap(text string) map[rune]int {

}

func buildHuffmanTree(freqMap map[rune]int) *Node {

}

func generateCodes(root *Node) map[rune]string {

}

func encode(text string, codes map[rune]string) string {

}

func decode(encoded string, root *Node) string {

}

func HuffmanCoding(text string) (string, string, map[rune]string) {

}

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
