# 875. Koko Eating Bananas

- Difficulty: medium
- Link: https://leetcode.com/problems/koko-eating-bananas/

## Approach

Binary search on the answer (the eating speed `K`). The feasible speeds range
from `1` up to the largest pile: any speed at or above the max clears each pile
in one hour, and total hours only *decrease* as `K` grows, so feasibility is
monotonic. For a candidate `midK`, the total hours is `sum(ceil(pile/midK))`;
if that exceeds `h` the speed is too slow, so search the upper half, otherwise
record it and search the lower half for a smaller feasible speed. The smallest
`K` whose total hours is `<= h` is returned. The answer is unique, so the test
compares against fixed expected values.

- Time:  O(n log m), where `m` is the largest pile
- Space: O(1)
