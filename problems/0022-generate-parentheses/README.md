# 22. Generate Parentheses

- Difficulty: medium
- Link: https://leetcode.com/problems/generate-parentheses/

## Approach

Backtracking. Build the string one character at a time while tracking how many
`(` (open) and `)` (close) have been placed so far. A `)` may be appended only
while `open > close` (so the prefix stays valid), and a `(` may be appended
whenever `open < n`. When both counters reach `n` the string is a complete,
well-formed combination and is recorded. The results come out in an arbitrary
order, so the test sorts both slices before comparing exact answers, verifies
every string is balanced and of length `2n`, and asserts the count equals the
nth Catalan number (e.g. 14 for `n = 4`).

- Time:  O(4^n / sqrt(n)) (the nth Catalan number of valid combinations)
- Space: O(n) recursion depth (excluding the output)
