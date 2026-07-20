# 76. Minimum Window Substring

- Difficulty: hard
- Link: https://leetcode.com/problems/minimum-window-substring/

## Approach

Sliding window with two frequency maps. `need` counts each character of `t`
(so multiplicity matters), and `needCount` is the number of *distinct*
characters still to satisfy. Expand the right edge over `s`, incrementing the
window count; whenever a character's window count exactly reaches its needed
count, one more requirement is `have`-satisfied. While all requirements are met
(`have == needCount`), record the current window if it is shorter than the best
so far, then shrink from the left until a requirement drops. Fast guards return
`""` when `t` is empty or longer than `s`. The answer is unique per LeetCode, so
the test compares the returned substring directly.

- Time:  O(len(s) + len(t)) — each index enters and leaves the window once.
- Space: O(k) where k is the number of distinct characters in `t`.
