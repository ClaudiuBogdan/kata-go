# Regular Expression Matching

This package implements a basic regular expression matching algorithm in Go. It supports a subset of regular expression features, including:

- `.` (dot) which matches any single character
- `*` (star) which matches zero or more of the preceding element

## Algorithm Description

The algorithm uses a recursive approach to match the pattern against the text. It handles the following cases:

1. Simple character matching
2. Dot (`.`) matching any character
3. Star (`*`) matching zero or more of the preceding element

The implementation provides two main functions:

1. `IsMatch(text, pattern string) bool`: Checks if the entire text matches the pattern
2. `FindAll(text, pattern string) []string`: Finds all non-overlapping matches of the pattern in the text

## Implementation Details

### Time Complexity
- `IsMatch`: O(2^(m+n)) in the worst case, where m and n are the lengths of the text and pattern respectively
- `FindAll`: O(n * 2^(m+n)) in the worst case, where n is the length of the text

### Space Complexity
- O(m+n) due to the recursive call stack

## Usage

```go
text := "aaabbbaaabbb"
pattern := "a*b*"

match := IsMatch(text, pattern)
fmt.Println("Is Match:", match)

matches := FindAll(text, pattern)
fmt.Println("All Matches:", matches)
```

## Testing

The implementation includes a comprehensive test suite covering various scenarios. To run the tests:

```bash
go test
```

## Implementation Notes

- The implementation uses recursion, which may lead to stack overflow for very long patterns or texts
- The `FindAll` function finds non-overlapping matches
- This is a basic implementation and does not support all features of full-fledged regex engines

## Advantages and Limitations

Advantages:
- Simple implementation
- Supports basic regex features (`.` and `*`)
- Can find all matches in a text

Limitations:
- Limited feature set compared to full regex engines
- May be inefficient for complex patterns or long texts due to recursion
- Doesn't support advanced features like capture groups, alternation, etc.

## When to Use

- For simple pattern matching needs where full regex capabilities are not required
- In educational contexts to understand basic regex matching algorithms
- As a starting point for building more complex regex engines

For production use cases requiring full regex capabilities, consider using Go's built-in `regexp` package or other established regex libraries.

## Future Improvements

- Implement memoization or dynamic programming to improve efficiency
- Add support for more regex features (e.g., `+`, `?`, character classes)
- Implement a non-recursive version to handle longer inputs
- Optimize the `FindAll` function for better performance

Remember that this implementation is meant for educational purposes and basic use cases. For complex or performance-critical applications, use well-established regex libraries or Go's built-in `regexp` package.
