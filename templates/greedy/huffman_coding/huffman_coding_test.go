package huffmancoding

import (
	"testing"
)

func TestHuffmanCoding(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Simple string", "abcdefg"},
		{"Repeated characters", "aaaaabbbbbccccdddeef"},
		{"Single character", "aaaaaaaaaa"},
		{"Empty string", ""},
		{"Long string", "The quick brown fox jumps over the lazy dog"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compressed, codes := Compress(tt.input)
			decompressed, err := Decompress(compressed, codes)

			if err != nil {
				t.Errorf("Decompression error: %v", err)
			}

			if decompressed != tt.input {
				t.Errorf("Decompressed string does not match input. Got %s, want %s", decompressed, tt.input)
			}

			if len(tt.input) > 0 {
				compressionRatio := float64(len(compressed)) / (float64(len(tt.input)) * 8)
				t.Logf("Compression ratio: %.2f", compressionRatio)
			}

			if !isValidHuffmanCode(codes) {
				t.Errorf("Invalid Huffman code generated")
			}
		})
	}
}

func isValidHuffmanCode(codes map[rune]string) bool {
	for _, code1 := range codes {
		for _, code2 := range codes {
			if code1 != code2 && (strings.HasPrefix(code1, code2) || strings.HasPrefix(code2, code1)) {
				return false
			}
		}
	}
	return true
}

func TestCompressDecompress(t *testing.T) {
	input := "hello world"
	compressed, codes := Compress(input)
	decompressed, err := Decompress(compressed, codes)

	if err != nil {
		t.Errorf("Decompression error: %v", err)
	}

	if decompressed != input {
		t.Errorf("Decompressed string does not match input. Got %s, want %s", decompressed, input)
	}
}

func TestInvalidDecompress(t *testing.T) {
	compressed := "1010101"
	codes := map[rune]string{'a': "0", 'b': "10", 'c': "11"}

	_, err := Decompress(compressed, codes)

	if err == nil {
		t.Error("Expected an error for invalid compressed data, but got nil")
	}
}
