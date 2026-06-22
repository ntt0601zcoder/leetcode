# 167. Two Sum Ii Input Array Is Sorted

- Difficulty: medium
- Link: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

## Approach

The input is sorted, so use two pointers from both ends. If the current sum is
too large, move the right pointer left; if too small, move the left pointer
right; when it equals the target, return the two 1-based indices. This finds
the unique pair in a single pass without extra space. The problem guarantees
exactly one solution, so the tests only cover targets with an unambiguous pair.

- Time:  O(n)
- Space: O(1)
