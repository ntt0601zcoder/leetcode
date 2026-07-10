# 297. Serialize And Deserialize Binary Tree

- Difficulty: hard
- Link: https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

## Approach

Level-order (BFS) with a null marker. **serialize**: seed a queue with the
root and walk it; write each node's value, using a marker (e.g. `-`) for nil,
and enqueue a real node's two children. Join with `,`. **deserialize**: split
by `,`; the first token is the root; keep a queue of nodes still needing
children and consume the remaining tokens in left/right pairs, skipping the
marker. Both directions must agree on the exact format, and `""` maps to `nil`.

Test note: because two different serialized strings can represent the same
tree, the test checks a **round-trip** (`deserialize(serialize(t))` is
structurally equal to `t`) rather than comparing the string.

- Time:  O(n) each direction
- Space: O(n)
