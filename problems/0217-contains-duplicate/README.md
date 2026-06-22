# 217. Contains Duplicate

- Difficulty: easy
- Link: https://leetcode.com/problems/contains-duplicate/

## Approach

Use a set (`map[int]...`): scan each element, return `true` if it's already in
the set, otherwise add it. If the scan finishes with no repeat, return `false`.

- Time:  O(n)
- Space: O(n)
