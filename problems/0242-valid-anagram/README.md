# 242. Valid Anagram

- Difficulty: easy
- Link: https://leetcode.com/problems/valid-anagram/

## Approach

Count character frequencies with a `map[rune]int`: add for `s`, subtract for
`t`, deleting on zero. If a character is missing (or left over), it's not an
anagram. Using `rune` handles Unicode too (the problem's follow-up).

- Time:  O(n)
- Space: O(k) where k is the number of distinct characters
