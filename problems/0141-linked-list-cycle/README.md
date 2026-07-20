# 141. Linked List Cycle

- Difficulty: easy
- Link: https://leetcode.com/problems/linked-list-cycle/

## Approach

Floyd's tortoise-and-hare. Two pointers start one node apart (`curr1` at the
head, `curr2` at `head.Next`); each iteration advances `curr1` by one and
`curr2` by two. If a cycle exists the fast pointer laps the slow one and they
land on the same node, so the loop returns `true` when `curr1 == curr2`. If the
fast pointer reaches the end (`nil` or a node whose `Next` is `nil`), the list
is acyclic and it returns `false`. Detection is by node **identity**
(`curr1 == curr2` compares pointers), so equal `Val`s never cause a false
positive — e.g. `{1,1,1,1}` with no cycle correctly returns `false`. The test
builds cycles with a `buildListCycle(vals, pos)` helper and guards each call
with a goroutine + 500ms timeout so a non-terminating check fails cleanly.

- Time:  O(n)
- Space: O(1)
