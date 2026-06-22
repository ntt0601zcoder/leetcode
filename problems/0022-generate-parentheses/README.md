# 22. Generate Parentheses

- Difficulty: medium
- Link: https://leetcode.com/problems/generate-parentheses/

## Approach

Backtracking. Build the string one character at a time while tracking how many
`(` (open) and `)` (close) have been placed. A `(` may be added whenever `open <
n`, and a `)` may be added only while `open > close` so the string stays valid.
When both counters reach `n` the string is a complete, well-formed combination
and is recorded. The order of the results is irrelevant, so the test sorts both
the produced and expected slices before comparing.

- Time:  O(4^n / sqrt(n)) (the nth Catalan number of valid combinations)
- Space: O(n) recursion depth (excluding the output)
