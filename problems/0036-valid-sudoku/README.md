# 36. Valid Sudoku

- Difficulty: medium
- Link: https://leetcode.com/problems/valid-sudoku/

## Approach

Single pass over the 9x9 board using bitmask sets. Three arrays of nine
integers track which digits have been seen in each row, each column, and each
3x3 box, with digit `d` represented by bit `1 << (d - '1')`. The box index is
computed as `(row/3)*3 + col/3`. For every filled cell, if its bit is already
set in the matching row, column, or box mask the board is invalid; otherwise the
bit is recorded. Empty cells (`.`) are skipped. The test covers a valid board
plus duplicates that are detectable only by the row, column, and box checks
respectively.

- Time:  O(1) (fixed 81 cells)
- Space: O(1)
