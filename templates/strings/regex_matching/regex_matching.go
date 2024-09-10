package regexmatching

// IsMatch checks if the text matches the pattern
func IsMatch(text, pattern string) bool {
	return isMatchHelper(text, pattern, 0, 0)
}

func isMatchHelper(text, pattern string, i, j int) bool {
	if j == len(pattern) {
		return i == len(text)
	}

	firstMatch := i < len(text) && (pattern[j] == text[i] || pattern[j] == '.')

	if j+1 < len(pattern) && pattern[j+1] == '*' {
		return isMatchHelper(text, pattern, i, j+2) || (firstMatch && isMatchHelper(text, pattern, i+1, j))
	}

	return firstMatch && isMatchHelper(text, pattern, i+1, j+1)
}

// FindAll returns all non-overlapping matches of the pattern in the text
func FindAll(text, pattern string) []string {
	var matches []string
	for i := 0; i < len(text); i++ {
		if IsMatch(text[i:], pattern) {
			end := i
			for end < len(text) && IsMatch(text[i:end+1], pattern) {
				end++
			}
			matches = append(matches, text[i:end])
			i = end - 1
		}
	}
	return matches
}
