# 18. 4sum

- Difficulty: medium
- Link: https://leetcode.com/problems/4sum/

## Approach

Sort the array, fix the two outer indices `i`, `j` (nested loops), then use a
two-pointer scan `k`/`h` to find the remaining pair that reaches `target`. To
keep the result **duplicate-free**: skip repeated `nums[i]` and `nums[j]`, and
after recording a quadruplet skip equal `nums[k]` / `nums[h]`.

Note: the answer may be returned in **any order**, so the test normalizes
(sort within each tuple + sort the list) before comparing.

- Time:  O(n^3)
- Space: O(1) extra (besides the output; excluding the sort cost)
