# 215. Kth Largest Element In An Array

- Difficulty: medium
- Link: https://leetcode.com/problems/kth-largest-element-in-an-array/

## Approach

Maintain a **min-heap of size `k`** holding the `k` largest values seen so far.
Push every number; whenever the heap grows past `k`, pop the smallest. After the
scan the heap contains exactly the `k` largest elements, and its root (the
minimum of those) is the `k`-th largest, returned as `(*h)[0]`. This avoids
sorting the whole array and streams through the input in one pass.

- Time:  O(n log k)
- Space: O(k)
