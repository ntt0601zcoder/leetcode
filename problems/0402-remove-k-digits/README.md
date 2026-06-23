# 402. Remove K Digits

- Difficulty: medium
- Link: https://leetcode.com/problems/remove-k-digits/

## Approach

Greedy monotonic (non-decreasing) stack. Scan the digits left to right; while
the stack top is **greater** than the current digit and we still have removals
left, pop it (a bigger high-order digit hurts more than a smaller one). Push
the current digit. If removals remain after the scan (the number is
non-decreasing), drop them from the end. Finally strip leading zeros, and
return `"0"` if nothing is left.

Note: the pop must be a `while`/`for` loop, not a single `if` — one incoming
small digit can require popping several larger digits at once.

- Time:  O(n)
- Space: O(n)
