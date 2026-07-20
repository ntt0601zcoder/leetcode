# 34. Find First And Last Position Of Element In Sorted Array

- Difficulty: medium
- Link: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

## Approach

Two binary searches over the sorted array. `findBound` locates a matching index,
then keeps shrinking toward one side: when `leftmost` is true it moves `right`
past the match to find the first occurrence, otherwise it moves `left` to find the
last. Call it once for each bound; if the leftmost search returns `-1` the target
is absent, so return `[-1, -1]`. The file also keeps a linear `searchRangeBruteForce`
(scan from both ends) that the tests run as a second approach for cross-checking.

- Time:  O(log n) for the binary-search version, O(n) for the brute force
- Space: O(1)
