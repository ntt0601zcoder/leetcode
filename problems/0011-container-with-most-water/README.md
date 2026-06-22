# 11. Container With Most Water

- Difficulty: medium
- Link: https://leetcode.com/problems/container-with-most-water/

## Approach

Two pointers from both ends. The area is bounded by the shorter of the two
walls times the width, so we always advance the pointer at the shorter wall:
moving the taller one could only shrink the width without ever exceeding the
current limiting height. Track the best area seen along the way until the
pointers meet.

- Time:  O(n)
- Space: O(1)
