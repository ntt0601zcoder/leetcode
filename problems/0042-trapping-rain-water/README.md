# 42. Trapping Rain Water

- Difficulty: hard
- Link: https://leetcode.com/problems/trapping-rain-water/

## Approach

Water trapped above bar `i` equals `min(maxLeft, maxRight) - height[i]`, summed
over all bars. Three approaches are provided. The brute force recomputes both
maxes for every index in O(n^2). The prefix/suffix variant precomputes the
running left and right maxima into two arrays for O(n) time at O(n) space. The
two-pointer version converges from both ends, always advancing the side with the
smaller bar, which guarantees the trailing max on that side is the binding limit,
giving O(n) time and O(1) space. The test cross-checks all three against the same
cases, including empty, single bar, flat, ascending/descending (no water), and
container shapes.

- Time:  O(n) (O(n^2) for the brute-force scan)
- Space: O(1) for two pointers, O(n) for prefix/suffix
