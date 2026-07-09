# 102. Binary Tree Level Order Traversal

- Difficulty: medium
- Link: https://leetcode.com/problems/binary-tree-level-order-traversal/

## Approach

BFS, one level at a time. Hold the current level's nodes in a queue; for each
level, read every node's value into a row and collect its children into the
next level's queue, then replace the queue with that next level. Append each
row to the result.

Note: the current level must be **drained** each pass — collect children into a
separate slice and swap it in (`queue = next`). Appending children back into the
same queue without removing the processed nodes never empties it, so the loop
never terminates.

- Time:  O(n)
- Space: O(n) (the widest level / the output)
