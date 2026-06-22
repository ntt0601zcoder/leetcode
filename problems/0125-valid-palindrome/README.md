# 125. Valid Palindrome

- Difficulty: easy
- Link: https://leetcode.com/problems/valid-palindrome/

## Approach

Two pointers from both ends moving inward. Each side skips characters that are
not alphanumeric, then the two characters are compared case-insensitively (via a
small `toLower` helper). If any pair differs the string is not a palindrome;
otherwise it is. No extra string is built, so it scans the input once in place.
Edge cases: the empty string and strings made entirely of punctuation (e.g.
`".,"`) are palindromes since no comparison ever fails.

- Time:  O(n)
- Space: O(1)
