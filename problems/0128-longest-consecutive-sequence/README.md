# 128. Longest Consecutive Sequence

- Difficulty: medium
- Link: https://leetcode.com/problems/longest-consecutive-sequence/

## Approach

Put every number in a hash set, which also drops duplicates. For each value,
only start counting when it is the beginning of a run (i.e. `num-1` is absent
from the set); otherwise skip it. From a run start, walk upward `num+1`,
`num+2`, ... while the set contains the next value, tracking the longest run
seen. Because each number is visited by the inner walk at most once across the
whole input, the total work is linear despite the nested loop.

- Time:  O(n)
- Space: O(n)
