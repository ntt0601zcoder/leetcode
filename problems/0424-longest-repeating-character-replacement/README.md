# 424. Longest Repeating Character Replacement

- Difficulty: medium
- Link: https://leetcode.com/problems/longest-repeating-character-replacement/

## Approach

Sliding window over `s` with a `[26]int` frequency table for the current
window. Expand `right` one char at a time and keep `maxLen`, the highest single
char frequency seen so far. A window is valid when `(right-left+1) - maxLen <= k`,
i.e. every non dominant char can be turned into the dominant one within `k`
replacements; whenever it is violated, shrink from `left`. The answer is the
largest valid window width. `maxLen` is intentionally never decreased on shrink:
the result only grows, so a stale `maxLen` can never inflate it. Assumes `s` is
uppercase A-Z per the constraints; an empty string returns 0.

- Time:  O(n)
- Space: O(1) (fixed 26 element counter)
