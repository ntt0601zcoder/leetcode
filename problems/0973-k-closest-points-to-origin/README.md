# 973. K Closest Points To Origin

- Difficulty: medium
- Link: https://leetcode.com/problems/k-closest-points-to-origin/

## Approach

Bounded **max-heap** of size k, keyed by distance to the origin. Push each
point; once the heap holds more than k, pop the root — a max-heap's root is the
*farthest*, so the k closest survive. (Alternatives: sort by distance and take
the first k, O(n log n); or quickselect, O(n) average.)

Two details that matter:

- The heap must be a **max**-heap (`Less` uses `>`) and eviction must go through
  `heap.Pop(h)`, not the raw `h.Pop()` method — the method only drops the last
  slice element and skips re-heapifying. A min-heap would evict the closest.
- Only the *ordering* of distances matters, so compare squared distances
  (`x*x + y*y`) — no `Sqrt`/`Pow` needed, exact in integers and much faster.

Test note: the answer may be returned in **any order**, so the test sorts the
list of points before comparing (never within a point — `[x,y]` order matters).

- Time:  O(n log k)
- Space: O(k)
