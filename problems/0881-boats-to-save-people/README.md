# 881. Boats To Save People

- Difficulty: medium
- Link: https://leetcode.com/problems/boats-to-save-people/

## Approach

Each boat carries at most two people and a total weight of at most `limit`.
Sort the people by weight, then use two pointers from the lightest (`i`) to
the heaviest (`j`). The heaviest remaining person must always go on a boat;
if the lightest remaining person also fits alongside them
(`people[i] + people[j] <= limit`), pair them up and advance both pointers,
otherwise the heaviest goes alone. Each iteration uses exactly one boat, so
incrementing a counter until the pointers cross gives the minimum number of
boats.

- Time:  O(n log n) for the sort
- Space: O(1) extra (sorts in place)

## Notes

The table test in `solution_test.go` copies the input slice before each call
because the solution sorts in place, and covers the LeetCode examples plus
edge cases: a single person, everyone too heavy to share, all people pairing
up, a heavy person who must go alone while the rest pair, and a pair summing
exactly to the limit.
