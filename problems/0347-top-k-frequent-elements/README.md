# 347. Top K Frequent Elements

- Difficulty: medium
- Link: https://leetcode.com/problems/top-k-frequent-elements/

## Approach

Count frequencies with a map, take the list of distinct elements, sort it by
frequency descending, then take the first `k`. (Can be optimized to O(n) with
a bucket sort by frequency, or a size-k heap.)

Note: the problem returns the answer in **any order** and the answer is
**unique**, so the test compares order-independently (sort both sides).

- Time:  O(n + m·log m) where m = number of distinct elements (O(n) with bucket sort)
- Space: O(m)
