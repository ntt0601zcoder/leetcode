# 162. Find Peak Element

- Difficulty: medium
- Link: https://leetcode.com/problems/find-peak-element/

## Approach

A peak is any index whose value is strictly greater than both neighbors, with
out-of-bounds neighbors treated as `-infinity`. The solution first handles the
boundaries: length-1 returns `0`, and if either end already rises above its only
neighbor it is returned directly. Otherwise both ends slope inward, so a peak
must lie in the interior `[1, n-2]`; a binary search walks uphill by moving
toward the larger neighbor of `mid` (adjacent elements are guaranteed distinct),
which always converges on a peak. The answer is NOT unique when multiple peaks
exist, so the test asserts the returned index satisfies the peak property rather
than comparing against a fixed index.

- Time:  O(log n)
- Space: O(1)
