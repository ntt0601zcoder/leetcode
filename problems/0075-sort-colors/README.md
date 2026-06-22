# 75. Sort Colors

- Difficulty: medium
- Link: https://leetcode.com/problems/sort-colors/

## Approach

Sort an array of 0s, 1s, and 2s in place. The primary solution is the Dutch
National Flag algorithm: three pointers `left`, `mid`, `right` partition the array
so everything before `left` is 0, everything after `right` is 2, and the middle is
1. Scanning with `mid`, a 0 is swapped to the front and a 2 to the back, sorting in
a single O(n) pass with O(1) space. A second counting-sort variant tallies each
value then overwrites the slice, also O(n) but using a small counter map. Both
mutate `nums` in place and return nothing, so the test copies each input first and
asserts the result equals the expected `[0s, 1s, 2s]` ordering.

- Time:  O(n)
- Space: O(1) (Dutch flag); O(1) extra plus a 3-key counter (counting variant)
