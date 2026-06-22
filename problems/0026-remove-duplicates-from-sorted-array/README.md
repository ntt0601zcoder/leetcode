# 26. Remove Duplicates From Sorted Array

- Difficulty: easy
- Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

## Approach

Two pointers over the sorted array. A write index `i` marks where the next
unique value goes and a read index `j` scans forward. Because the array is
sorted, a duplicate is detected by comparing the last written element
(`nums[i-1]`) with `nums[j]`: if they match, skip `j`; otherwise copy `nums[j]`
into `nums[i]` and advance both. Inputs of length 0 or 1 are handled up front.
The function returns the count `k` of unique elements and mutates `nums` in
place, so the test copies each input first, then asserts both the returned `k`
and that `nums[:k]` equals the expected deduplicated prefix.

- Time:  O(n)
- Space: O(1)
