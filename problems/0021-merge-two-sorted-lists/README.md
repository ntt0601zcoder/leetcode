# 21. Merge Two Sorted Lists

- Difficulty: easy
- Link: https://leetcode.com/problems/merge-two-sorted-lists/

## Approach

Iterative merge with a dummy head node. Walk both lists at once, always
splicing the smaller current node onto the tail (`curr`), then advance that
list. Using `<=` keeps the merge stable, so nodes with equal values from
`list1` come first. When either list runs out, append whatever remains of the
other in one step. No new nodes are allocated: existing nodes are re-linked in
place, and `dummy.Next` is the merged head.

- Time:  O(n + m)
- Space: O(1) extra (besides the reused input nodes)
