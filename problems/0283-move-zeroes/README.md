# 283. Move Zeroes

- Difficulty: easy
- Link: https://leetcode.com/problems/move-zeroes/

## Approach

In-place two-pointer (write/read) scan. The read pointer `j` walks the array
while the write pointer `i` lags behind: every non-zero value is copied forward
to `i`, and once `j` runs off the end the remaining tail at `i` is filled with
zeros. This preserves the relative order of the non-zero elements and moves all
zeros to the end without allocating a new slice. The function mutates the slice
and returns nothing, so the test copies each input, calls it, then compares the
mutated slice against the expected layout.

- Time:  O(n)
- Space: O(1)
