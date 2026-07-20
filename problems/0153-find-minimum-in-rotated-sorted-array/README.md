# 153. Find Minimum In Rotated Sorted Array

- Difficulty: medium
- Link: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

## Approach

Binary search on the rotation point. Keep a `[left, right]` window and compare
`nums[mid]` with `nums[right]`: if `nums[mid] > nums[right]` the minimum must lie
strictly to the right, so move `left = mid + 1`; otherwise the minimum is at
`mid` or to its left, so keep `right = mid`. Comparing against the right bound
(not the left) makes the fully-sorted, non-rotated case fall through correctly.
The loop narrows until `left == right`, which points at the minimum.

- Time:  O(log n)
- Space: O(1)
