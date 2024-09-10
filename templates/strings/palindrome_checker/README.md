# Palindrome Checker

## Problem Description

A palindrome is a word, phrase, number, or other sequence of characters that reads the same forward and backward, disregarding spaces, punctuation, and capitalization. This implementation provides a function to check if a given string is a palindrome.

## Implementation Details

The `IsPalindrome` function takes one parameter:
- `s`: The string to check.

It returns a boolean value: `true` if the string is a palindrome, `false` otherwise.

### Time Complexity
- O(n), where n is the length of the string.

### Space Complexity
- O(n) for the cleaned string, O(1) for the actual comparison.

## Usage

```go
result := IsPalindrome("A man a plan a canal Panama")
fmt.Println(result) // Output: true

result = IsPalindrome("hello world")
fmt.Println(result) // Output: false
```

## Testing

The implementation includes a comprehensive test suite covering various scenarios. To run the tests:

```bash
go test
```

## Algorithm Explanation

1. Clean the input string:
   - Convert to lowercase.
   - Remove non-alphanumeric characters.
2. Use two pointers, one at the start and one at the end of the cleaned string.
3. Move the pointers towards the center, comparing characters at each step.
4. If at any point the characters don't match, return false.
5. If the pointers meet or cross without finding a mismatch, return true.

## Implementation Notes

- The implementation is case-insensitive and ignores non-alphanumeric characters.
- It uses a two-pointer approach for efficient comparison.
- The `cleanString` function uses `strings.Map` for efficient string manipulation.
- Unicode characters are handled correctly.

## Advantages

1. Efficient: O(n) time complexity.
2. Case-insensitive: Treats uppercase and lowercase letters as the same.
3. Ignores non-alphanumeric characters: Focuses on the essence of the palindrome.
4. Handles various edge cases: Empty strings, single characters, unicode characters.

## Limitations

1. Creates a new string for the cleaned version, which uses additional memory.
2. May not be suitable for extremely large strings due to memory constraints of creating a cleaned copy.

## When to Use

- When you need to determine if a string is a palindrome.
- In text processing applications, word games, or linguistic analysis.
- As part of larger algorithms or systems that involve string manipulations.

Remember, while this implementation is efficient for most use cases, always consider the specific requirements of your application, such as memory constraints or the need for in-place algorithms.
