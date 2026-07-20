# 104. Maximum Depth Of Binary Tree

- Difficulty: easy
- Link: https://leetcode.com/problems/maximum-depth-of-binary-tree/

## Approach

Recursive DFS. A `nil` node has depth 0; any other node's depth is
`1 + max(depth(left), depth(right))`. The recursion bottoms out at the leaves
and each frame returns the taller of its two subtrees plus one for itself, so
the root returns the length of the longest root-to-leaf path.

- Time:  O(n) - each node is visited once.
- Space: O(h) recursion stack, where h is the tree height (O(n) worst case for a skewed tree, O(log n) when balanced).
