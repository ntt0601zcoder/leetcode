# 143. Reorder List

- Difficulty: medium
- Link: https://leetcode.com/problems/reorder-list/

## Approach

Reorder `L0 -> L1 -> ... -> Ln` in place to `L0 -> Ln -> L1 -> Ln-1 -> ...`.
Three steps, all O(1) extra space: (1) find the middle with slow/fast pointers,
(2) reverse the second half, (3) splice the two halves together by alternating
nodes. The function mutates the list in place and returns nothing, so the test
builds a fresh list per case and reads the result back with `listVals`.

Note: the current solution does not guard the empty-list case. When `head` is
`nil`, `slowNode` stays `nil` and `reverseList(slowNode.Next)` (solution.go:19)
dereferences a nil pointer and panics. The correct spec is that an empty list
stays empty, so the test's `empty` case fails by design; the minimal fix is an
early `if head == nil { return }` at the top of `reorderList`.

- Time:  O(n)
- Space: O(1) extra
