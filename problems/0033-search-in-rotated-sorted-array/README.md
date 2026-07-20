# 33. Search In Rotated Sorted Array

- Difficulty: medium
- Link: https://leetcode.com/problems/search-in-rotated-sorted-array/

## Approach

Modified binary search in a single pass. At each `mid`, exactly one half is
sorted. If `nums[mid] >= nums[left]` the left half `[left..mid]` is sorted, so
if `target` lies in `[nums[left], nums[mid])` we search left, otherwise right.
Otherwise the right half `[mid..right]` is sorted, so if `target` lies in
`(nums[mid], nums[right]]` we search right, otherwise left. This keeps the
search logarithmic despite the rotation, and returns the found index or `-1`.

- Time:  O(log n)
- Space: O(1)
