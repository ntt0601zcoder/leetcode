# 23. Merge K Sorted Lists

- Difficulty: hard
- Link: https://leetcode.com/problems/merge-k-sorted-lists/

## Approach

Divide and conquer by pairwise merging. Each round walks the slice of list
heads two at a time, merging every adjacent pair with a standard two-pointer
`mergeTwoLists` (a trailing unpaired list is merged against `nil`), and replaces
the slice with the halved set of merged heads. After `log k` rounds a single
list remains. `mergeTwoLists` relinks the existing nodes in place, so no new
nodes are allocated. Empty inputs, empty sub-lists, and an all-nil slice all
collapse to `nil`.

- Time:  O(N log k) — N total nodes, each touched once per round over log k rounds
- Space: O(k) for the working slice of list heads (O(1) node-relinking, aside from the output)
