# 169. Majority Element

- Difficulty: easy
- Link: https://leetcode.com/problems/majority-element/

## Approach

Count frequencies with a map; return as soon as an element exceeds `n/2`. The
problem guarantees a majority element always exists.

More efficient: the Boyer-Moore voting algorithm — O(n) time, O(1) space.

- Time:  O(n)
- Space: O(n) (Boyer-Moore: O(1))
