# Longest Common Subsequence (Dynamic Programming Solution)

## Problem Description

The Longest Common Subsequence (LCS) problem is to find the longest subsequence present in given two sequences in the same order, i.e., find the longest sequence which can be obtained from both the original sequences by deleting some elements without changing the order of the remaining elements.

## Approach

This solution uses dynamic programming to solve the Longest Common Subsequence problem efficiently. The algorithm builds a table where each entry `dp[i][j]` represents the length of the LCS of the prefixes `s1[0:i]` and `s2[0:j]`.

### Algorithm Steps:

1. Create a 2D array `dp` of size `(m+1) x (n+1)`, where `m` and `n` are the lengths of the input strings.
2. Initialize the first row and column of `dp` with 0.
3. For each character in s1 (i) and s2 (j):
   - If the characters match, `dp[i][j] = dp[i-1][j-1] + 1`
   - If they don't match, `dp[i][j] = max(dp[i-1][j], dp[i][j-1])`
4. The value in `dp[m][n]` gives the length of the longest common subsequence.
5. To find the actual LCS string, backtrack through the dp table.

## Complexity Analysis

- Time Complexity: O(m * n), where m and n are the lengths of the input strings.
- Space Complexity: O(m * n) for the dp table.

## Usage

```go
s1 := "ABCDGH"
s2 := "AEDFHR"

length := LongestCommonSubsequence(s1, s2)
fmt.Println("Length of LCS:", length)

lcs := LongestCommonSubsequenceString(s1, s2)
fmt.Println("Longest Common Subsequence:", lcs)
```

## Implementation Details

The package provides two main functions:

1. `LongestCommonSubsequence(s1, s2 string) int`: Returns the length of the longest common subsequence.
2. `LongestCommonSubsequenceString(s1, s2 string) string`: Returns the actual longest common subsequence as a string.

## Testing

The implementation includes a test suite that covers various scenarios:

1. Basic LCS problems
2. Strings with no common subsequence
3. Cases with empty strings
4. Identical strings
5. Longer strings with multiple possible LCS

To run the tests, use the following command:

```bash
go test
```

## Advantages of Dynamic Programming

- Optimal solution: Guarantees to find the longest common subsequence.
- Efficiency: Avoids redundant calculations by storing intermediate results.
- Versatility: Can be easily modified to solve variations of the LCS problem.

## Applications

- DNA sequence analysis in bioinformatics
- File comparison tools (e.g., diff utility)
- Version control systems for tracking changes in files
- Spell checking and correction algorithms
- Speech recognition systems

This dynamic programming solution to the Longest Common Subsequence problem provides an efficient and optimal approach to a fundamental string comparison problem, with applications in various fields including bioinformatics, text processing, and version control systems.
