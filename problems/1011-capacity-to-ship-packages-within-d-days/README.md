# 1011. Capacity To Ship Packages Within D Days

- Difficulty: medium
- Link: https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/

## Approach

Binary search on the answer (the ship capacity). A capacity is feasible if
greedily filling days in order needs at most `days` days. The feasible
capacities form a monotonic range, so binary-search the smallest one.

Bounds matter: the lower bound is `max(weights)` (a day must hold the heaviest
package), the upper bound is `sum(weights)` (ship everything in one day). And
the day count must include the final in-progress day — start the counter at 1,
not 0, or the search returns the answer for `days + 1` days.

- Time:  O(n · log(sum - max))
- Space: O(1)
