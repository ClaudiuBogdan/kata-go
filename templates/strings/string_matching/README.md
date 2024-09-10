# String Matching Algorithms

This package implements two string matching algorithms:
1. Naive (Brute Force) String Matching
2. Knuth-Morris-Pratt (KMP) String Matching

## Algorithms Description

### Naive String Matching

The naive approach involves checking for a match at every position in the text. It's simple but can be inefficient for large texts or patterns.

### Knuth-Morris-Pratt (KMP)

KMP improves upon the naive approach by utilizing information about the pattern to skip unnecessary comparisons. It precomputes a failure function to determine how much to shift the pattern when a mismatch occurs.

## Implementation Details

Both algorithms return a slice of integers representing the starting indices of all occurrences of the pattern in the text.

### Time Complexity
- Naive: O(nm), where n is the length of the text and m is the length of the pattern
- KMP: O(n + m) for matching, with an additional O(m) preprocessing step

### Space Complexity
- Naive: O(1) (excluding the space for storing matches)
- KMP: O(m) for the failure function

## Usage

```go
text := "ABABDABACDABABCABAB"
pattern := "ABABCABAB"

naiveMatches := NaiveStringMatch(text, pattern)
kmpMatches := KMPStringMatch(text, pattern)

fmt.Println("Naive matches:", naiveMatches)
fmt.Println("KMP matches:", kmpMatches)
```

## Testing

The implementation includes a comprehensive test suite covering various scenarios. To run the tests:

```bash
go test
```

## Implementation Notes

- Both algorithms handle edge cases such as empty strings and patterns longer than the text.
- The KMP algorithm includes a separate function to compute the failure function.
- The implementations return all matches, not just the first one.

## Advantages and Limitations

### Naive Algorithm
Advantages:
- Simple to understand and implement
- Works well for short texts and patterns

Limitations:
- Inefficient for long texts or patterns
- Performs unnecessary comparisons

### KMP Algorithm
Advantages:
- Efficient for long texts and patterns
- Avoids unnecessary comparisons
- Linear time complexity in the worst case

Limitations:
- More complex to implement
- Requires additional space for the failure function
- Overkill for very short texts or patterns

## When to Use

- Use the Naive algorithm for simple cases or when simplicity is more important than efficiency.
- Use KMP for large texts or patterns, or when multiple searches will be performed with the same pattern.

Remember to consider factors like the size of your data, frequency of searches, and implementation complexity when choosing between these algorithms.
