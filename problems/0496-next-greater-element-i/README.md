# 496. Next Greater Element I

- Difficulty: easy
- Link: https://leetcode.com/problems/next-greater-element-i/

## Approach

Monotonic (decreasing) stack over `nums2`. Scan `nums2` left to right keeping a
stack of values still waiting for a greater element. For each `num`, pop every
stack entry smaller than it and record `greaterMap[popped] = num` — that `num`
is their next greater element. Push `num`. Any value left on the stack never
found a greater element, so it stays absent from the map. Finally map each
element of `nums1` through `greaterMap`, defaulting to `-1` when missing. The
output preserves `nums1`'s order, so the test compares with `reflect.DeepEqual`.

- Time:  O(n + m) where n = len(nums2), m = len(nums1)
- Space: O(n) for the stack and map
