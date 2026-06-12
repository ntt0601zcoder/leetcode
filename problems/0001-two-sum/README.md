# 1. Two Sum

- Difficulty: easy
- Link: https://leetcode.com/problems/two-sum/

## Approach

Hash map from value to its index. For each `n`, check whether the
complement `target - n` was seen earlier; if so the pair is found in a
single pass.

- Time:  O(n)
- Space: O(n)
