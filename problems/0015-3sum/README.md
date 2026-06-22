# 15. 3Sum

- Difficulty: medium
- Link: https://leetcode.com/problems/3sum/

## Approach

Sort the array, then for each `nums[i]` run a two-pointer scan over the
remaining suffix looking for a pair that sums to `-nums[i]`. Move the left
pointer up when the sum is too small and the right pointer down when it is too
large. Duplicates are skipped both at the anchor `i` and after recording a hit
so every triplet is reported exactly once. The result order is unspecified, so
the test normalizes (sorts each triplet, then sorts the list) before comparing.

- Time:  O(n^2)
- Space: O(1) extra (besides the output; sort is in place)
