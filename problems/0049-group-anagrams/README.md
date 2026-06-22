# 49. Group Anagrams

- Difficulty: medium
- Link: https://leetcode.com/problems/group-anagrams/

## Approach

Current approach: group strings by comparing each one against the
representative key of every existing group (`isAnagramMap`). Simple but
O(n²·L), since each string is compared against all current groups.

A more common O(n·L·logL) approach: use the *sorted string* (or a 26-letter
count) as the map key, so each string is looked up once.

Test note: the order of the groups in the result is **undefined** (map
iteration), so the test compares after normalizing (sort within each group +
sort the groups).

- Time:  O(n²·L) current / O(n·L·logL) with a sorted key
- Space: O(n·L)
