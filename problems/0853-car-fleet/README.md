# 853. Car Fleet

- Difficulty: medium
- Link: https://leetcode.com/problems/car-fleet/

## Approach

A car can never pass the one ahead of it, so two cars form a single fleet
when the rear car would reach the target no later than the car in front. Pair
each car's position with the time it needs to reach the target,
`(target - position) / speed`, then sort the cars by position descending
(front of the road first). Sweep from front to back keeping the running
maximum arrival time as the current fleet's `lead`: a car whose own time is
strictly greater than `lead` cannot catch up, so it starts a new fleet and
becomes the new lead; otherwise it merges into the fleet ahead. The count of
new leads is the number of fleets.

- Time:  O(n log n) for the sort
- Space: O(n) for the position/time pairs

## Notes

The table test in `solution_test.go` copies the input slices before each call
since the solution sorts in place, and covers the LeetCode examples plus edge
cases: a single car, no cars, equal speeds that never merge, two cars that do
versus do not merge, an already-ordered pair, and all cars merging into one
fleet.
