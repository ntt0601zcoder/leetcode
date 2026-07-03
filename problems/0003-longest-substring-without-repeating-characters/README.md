# 3. Longest Substring Without Repeating Characters

- Difficulty: medium
- Link: https://leetcode.com/problems/longest-substring-without-repeating-characters/

## Approach

Sliding window with a set of the characters currently in the window. Extend
`right` by one each step; if the new character is already in the window, shrink
from `left` (removing characters) until it's gone, then add it. Track the
largest window size seen.

Note: the shrink loop must re-check membership every iteration (re-read the set
in the loop condition) — reading `ok` only once leaves the loop running until it
indexes past the string and panics.

- Time:  O(n) (each character enters and leaves the window once)
- Space: O(min(n, alphabet))
