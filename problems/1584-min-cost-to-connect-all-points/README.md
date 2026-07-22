# 1584. Min Cost To Connect All Points

- Difficulty: medium
- Link: https://leetcode.com/problems/min-cost-to-connect-all-points/

## Approach

Minimum spanning tree over a complete graph where the edge weight between two
points is their Manhattan distance. Lazy **Prim's** with a min-heap: start from
node 0 at cost 0; repeatedly pop the cheapest edge to an unvisited node, add it
to the tree, and push edges from that node to every still-unvisited node.
Skip popped nodes already in the tree. Total edge cost is the answer.

Note: the heap entry is `{cost, nodeIndex}` — when computing a neighbor's
distance, use the *coordinates* of the node just added (`points[u]`), not the
`{cost, u}` heap entry.

- Time:  O(n^2 log n) (lazy Prim's over a dense graph)
- Space: O(n^2) heap in the worst case
