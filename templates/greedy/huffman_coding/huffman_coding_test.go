package huffmancoding

import (
	"reflect"
	"strings"
	"testing"
)
func TestHuffmanCoding(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Empty string", ""},
		{"Single character", "a"},
		{"Two identical characters", "aa"},
		{"Two different characters", "ab"},
		{"Simple string", "aabbc"},
		{"Longer string", "abracadabra"},
		{"Repeated characters", "aaaaabbbbbcccccddddd"},
	}
	

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, decoded, codes := HuffmanCoding(tt.input)

			// Check if decoded matches the input
			if decoded != tt.input {
				t.Errorf("Decoded text does not match input. Got %q, want %q", decoded, tt.input)
			}

			// Check if encoded only contains '0' and '1'
			if encoded != "" && strings.Trim(encoded, "01") != "" {
				t.Errorf("Encoded text should only contain '0' and '1'. Got: %q", encoded)
			}

			// Check if the encoding is correct
			if tt.input != "" {
				decodedFromEncoded := decode(encoded, buildHuffmanTree(buildFrequencyMap(tt.input)))
				if decodedFromEncoded != tt.input {
					t.Errorf("Encoding is incorrect. Decoded %q from encoded, want %q", decodedFromEncoded, tt.input)
				}
			}

			// Check if codes map is correct
			if tt.input == "" {
				if codes != nil {
					t.Error("Codes map should be nil for empty input")
				}
			} else {
				if codes == nil {
					t.Error("Codes map is nil for non-empty input")
				} else {
					// Check if all characters in input have a code
					for _, ch := range tt.input {
						if _, ok := codes[ch]; !ok {
							t.Errorf("No code found for character %q", ch)
						}
					}
					// Check if all codes are unique
					codeSet := make(map[string]bool)
					for _, code := range codes {
						if codeSet[code] {
							t.Errorf("Duplicate code found: %q", code)
						}
						codeSet[code] = true
					}
					// Check if codes are prefix-free
					for char1, code1 := range codes {
						for char2, code2 := range codes {
							if char1 != char2 && (strings.HasPrefix(code1, code2) || strings.HasPrefix(code2, code1)) {
								t.Errorf("Codes are not prefix-free: %q (%s) and %q (%s)", char1, code1, char2, code2)
							}
						}
					}
				}
			}

			// Check compression ratio for longer inputs
			if len(tt.input) > 10 {
				compressionRatio := float64(len(encoded)) / float64(len(tt.input)*8)
				if compressionRatio >= 1.0 {
					t.Errorf("Compression ratio is not satisfactory: %.2f", compressionRatio)
				}
			}
		})
	}
}

func TestGenerateHuffmanCodes(t *testing.T) {
	input := "aabbbcccc"
	root := buildHuffmanTree(buildFrequencyMap(input))
	codes := generateCodes(root)

	// Check if all characters have a code
	for _, char := range input {
		if _, ok := codes[char]; !ok {
			t.Errorf("No code generated for character %q", char)
		}
	}

	// Check if codes are unique
	usedCodes := make(map[string]bool)
	for char, code := range codes {
		if usedCodes[code] {
			t.Errorf("Duplicate code %q for character %q", code, char)
		}
		usedCodes[code] = true
	}

	// Check if codes are prefix-free
	for char1, code1 := range codes {
		for char2, code2 := range codes {
			if char1 != char2 {
				if strings.HasPrefix(code1, code2) || strings.HasPrefix(code2, code1) {
					t.Errorf("Codes are not prefix-free: %q (%s) and %q (%s)", char1, code1, char2, code2)
				}
			}
		}
	}

	// Check if more frequent characters have shorter codes
	freq := buildFrequencyMap(input)
	for char1, code1 := range codes {
		for char2, code2 := range codes {
			if freq[char1] > freq[char2] && len(code1) > len(code2) {
				t.Errorf("More frequent character %q has longer code than less frequent character %q", char1, char2)
			}
		}
	}

	// Print the generated codes for inspection
	t.Logf("Generated Huffman codes: %v", codes)
}

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"Empty string", ""},
		{"Single character", "a"},
		{"Two identical characters", "aa"},
		{"Two different characters", "ab"},
		{"Simple string", "hello world"},
		{"Longer string", "the quick brown fox jumps over the lazy dog"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, decoded, _ := HuffmanCoding(tt.input)
			if decoded != tt.input {
				t.Errorf("Encode/Decode failed. Got %q, want %q", decoded, tt.input)
			}
			if tt.input != "" && encoded == "" {
				t.Error("Encoded string is empty for non-empty input")
			}
		})
	}
}

func TestPrintHuffmanCodes(t *testing.T) {
	codes := map[rune]string{
		'a': "0",
		'b': "10",
		'c': "110",
		'd': "111",
	}

	expected := []string{
		"a: 0",
		"b: 10",
		"c: 110",
		"d: 111",
	}

	result := PrintHuffmanCodes(codes)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("PrintHuffmanCodes result does not match expected. Got %v, want %v", result, expected)
	}
}
