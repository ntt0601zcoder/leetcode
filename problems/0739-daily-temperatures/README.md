# 739. Daily Temperatures

- Difficulty: medium
- Link: https://leetcode.com/problems/daily-temperatures/

## Approach

**Monotonic stack** (`dailyTemperatures`): the stack holds *indices* of days
waiting for a warmer day, with temperatures decreasing from bottom to top. On
a day `i` warmer than the stack top, pop each day `prev` and set
`result[prev] = i - prev`. Each index is pushed/popped once → O(n).

**Brute force** (`dailyTemperaturesBigOn`): for each `i`, scan forward until a
warmer day. Simple but O(n²) (the name "BigOn" is misleading — it's actually
O(n²)).

Note: only a **strictly warmer** day counts (compare with `>`); equal
temperatures don't.

- Time:  O(n) stack / O(n²) brute
- Space: O(n)
