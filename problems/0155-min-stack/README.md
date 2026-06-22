# 155. Min Stack

- Difficulty: medium
- Link: https://leetcode.com/problems/min-stack/

## Approach

Back each stack entry with a `Node` holding its own value plus the minimum of
the whole stack at the time it was pushed. On `Push`, the node's `minValue` is
the smaller of the new value and the current top's `minValue`, so the running
minimum is carried alongside every element. `Pop` just trims the slice, `Top`
returns the top node's value, and `GetMin` returns the top node's `minValue`.
This keeps all four operations O(1) without scanning the stack. The test drives
operation sequences (duplicates of the min, popping to expose a previous min,
single element) and checks `Top`/`GetMin` after each step.

- Time:  O(1) per operation
- Space: O(n)
