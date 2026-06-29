# 74. Search A 2d Matrix

- Difficulty: medium
- Link: https://leetcode.com/problems/search-a-2d-matrix/

## Approach

Because each row is sorted and every row's first value exceeds the previous
row's last, the matrix reads as one sorted array of `rows*cols` values. Binary
search over flat indices `0..rows*cols-1`, mapping each `mid` back to a cell
with `row = mid / cols`, `col = mid % cols` (here `m` is the column count).

Note: the row mapping must be integer division `mid / cols` — using
`ceil(mid/cols) - 1` is off by one at the first column and yields `-1` (a
panic) when `mid == 0`.

- Time:  O(log(m·n))
- Space: O(1)
