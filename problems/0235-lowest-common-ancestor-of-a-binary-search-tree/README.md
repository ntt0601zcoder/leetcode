# 235. Lowest Common Ancestor Of A Binary Search Tree

- Difficulty: medium
- Link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

## Approach

Exploit the BST ordering: walk down from the root without any extra
storage. If both `p` and `q` are smaller than the current node, the LCA lies
in the left subtree; if both are larger, it lies in the right subtree.
Otherwise the paths to `p` and `q` diverge here (or one of them *is* this
node), so the current node is the lowest common ancestor. The LCA is unique,
so the test hands the solution pointers to the actual `p` and `q` nodes (found
via a BST lookup) and asserts the returned node's value.

- Time:  O(h) where h is the tree height (O(log n) balanced, O(n) skewed)
- Space: O(1)
