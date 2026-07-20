# 206. Reverse Linked List

- Difficulty: easy
- Link: https://leetcode.com/problems/reverse-linked-list/

## Approach

Iterative in-place pointer reversal. Walk the list with a `curr` pointer while
keeping a `prev` pointer (initially `nil`). At each node, stash `curr.Next` in
`next`, repoint `curr.Next` back to `prev`, then advance both `prev` and `curr`
forward. When `curr` reaches `nil`, `prev` is the new head. Handles the empty
list (returns `nil`) and single node naturally, and no auxiliary storage is
allocated since the existing nodes are simply relinked.

- Time:  O(n)
- Space: O(1)
