# 704. Binary Search

- Difficulty: easy
- Link: https://leetcode.com/problems/binary-search/

## Approach

Classic binary search on a sorted array. Keep an inclusive window
`[left, right]`; compute `mid = left + (right-left)/2` (avoids overflow). If
`nums[mid] == target` return `mid`; if it's larger, search left
(`right = mid - 1`); if smaller, search right (`left = mid + 1`). Return -1
when the window empties (`left > right`).

Note: the bounds must step **past** `mid` (`mid ± 1`), not to `mid`, otherwise
the window can stop shrinking and the loop never terminates.

- Time:  O(log n)
- Space: O(1)
