# Anagram Checker

## Problem Description

An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once. This implementation provides a function to check if two given strings are anagrams of each other.

## Implementation Details

The `AreAnagrams` function takes two parameters:
- `s1`: The first string to compare.
- `s2`: The second string to compare.

It returns a boolean value: `true` if the strings are anagrams, `false` otherwise.

### Time Complexity
- O(n), where n is the length of the longer string.

### Space Complexity
- O(k), where k is the number of unique characters in the strings (worst case is O(1) since there's a finite set of characters).

## Usage

```go
result := AreAnagrams("listen", "silent")
fmt.Println(result) // Output: true

result = AreAnagrams("hello", "world")
fmt.Println(result) // Output: false
```

## Testing

The implementation includes a comprehensive test suite covering various scenarios. To run the tests:

```bash
go test
```

## Algorithm Explanation

1. Clean both input strings:
   - Convert to lowercase.
   - Remove non-alphanumeric characters.
2. Check if the cleaned strings have the same length. If not, they can't be anagrams.
3. Create a map to count the occurrences of each character in the first string.
4. Iterate through the second string, decrementing the count for each character.
5. If at any point a character count becomes negative, the strings are not anagrams.
6. If all counts are zero at the end, the strings are anagrams.

## Implementation Notes

- The implementation is case-insensitive and ignores non-alphanumeric characters.
- It uses a map to count character occurrences, which provides O(1) lookup time.
- The `cleanString` function uses `strings.Map` for efficient string manipulation.

## Advantages

1. Efficient: O(n) time complexity.
2. Case-insensitive: Treats uppercase and lowercase letters as the same.
3. Ignores non-alphanumeric characters: Focuses on the essence of the anagram.
4. Handles various edge cases: Empty strings, single characters, repeating characters.

## Limitations

1. Only works for strings that use characters representable as Go runes.
2. May not be suitable for extremely large strings due to memory constraints.

## When to Use

- When you need to determine if two strings are anagrams of each other.
- In applications dealing with word games or linguistic analysis.
- As part of larger algorithms or systems that involve string comparisons.

Remember, while this implementation is efficient for most use cases, always consider the specific requirements of your application, such as handling of special characters or very large strings.
