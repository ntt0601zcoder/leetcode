# 4. Median Of Two Sorted Arrays

- Difficulty: hard
- Link: https://leetcode.com/problems/median-of-two-sorted-arrays/

## Approach

Binary-search partition. Always search over the shorter array and pick a cut
`shortCut` in it; the matching cut in the longer array is fixed by
`longCut = (total+1)/2 - shortCut` so the left side always holds exactly half the
elements. Shrink the search range until the partition is valid
(`shortLeft <= longRight` and `longLeft <= shortRight`), using `math.MinInt` /
`math.MaxInt` sentinels for out-of-range edges. For an odd total the median is
`max(leftEnds)`; for an even total it is the average of `max(leftEnds)` and
`min(rightEnds)`. The file also keeps two fallbacks: an identical binary-search
variant and a simple concat-and-sort version. The result is a `float64`, so the
test compares with a tolerance (`abs diff < 1e-9`) rather than `==`.

- Time:  O(log(min(m, n))) for the binary-search approaches; O((m+n) log(m+n)) for the sort fallback
- Space: O(1) for the binary-search approaches; O(m+n) for the sort fallback
