# 84. Largest Rectangle In Histogram

- Difficulty: hard
- Link: https://leetcode.com/problems/largest-rectangle-in-histogram/

## Approach

Maintain a stack of bars in non-decreasing height order. When the current bar is
shorter than the stack top, pop bars and compute the area each popped bar can span:
its height times the width from its recorded start index up to the current index.
The popped bar's start index is carried over to the bar being pushed, which extends
the new bar leftward to where it could still reach. After the scan, any bars left on
the stack extend to the end (`n`). Tracking `maxArea` across every pop yields the
answer in a single O(n) pass. The test covers increasing, decreasing, single bar,
all-equal heights, and zero-split cases.

- Time:  O(n)
- Space: O(n)
