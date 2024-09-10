package stringmatching

// NaiveStringMatch performs naive (brute force) string matching
func NaiveStringMatch(text, pattern string) []int {
	var matches []int
	n, m := len(text), len(pattern)

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && text[i+j] == pattern[j] {
			j++
		}
		if j == m {
			matches = append(matches, i)
		}
	}

	return matches
}

// KMPStringMatch performs Knuth-Morris-Pratt string matching
func KMPStringMatch(text, pattern string) []int {
	var matches []int
	n, m := len(text), len(pattern)

	// Compute the failure function
	fail := computeKMPFailure(pattern)

	// Perform the matching
	j := 0 // Index for pattern
	for i := 0; i < n; i++ {
		for j > 0 && pattern[j] != text[i] {
			j = fail[j-1]
		}
		if pattern[j] == text[i] {
			j++
		}
		if j == m {
			matches = append(matches, i-m+1)
			j = fail[j-1]
		}
	}

	return matches
}

// computeKMPFailure computes the failure function for KMP
func computeKMPFailure(pattern string) []int {
	m := len(pattern)
	fail := make([]int, m)
	j := 0
	for i := 1; i < m; i++ {
		for j > 0 && pattern[j] != pattern[i] {
			j = fail[j-1]
		}
		if pattern[j] == pattern[i] {
			j++
		}
		fail[i] = j
	}
	return fail
}
