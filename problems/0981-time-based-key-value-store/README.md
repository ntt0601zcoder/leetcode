# 981. Time Based Key Value Store

- Difficulty: medium
- Link: https://leetcode.com/problems/time-based-key-value-store/

## Approach

Keep a map from key to a slice of `(timestamp, value)` entries. `Set` appends;
since the problem guarantees timestamps for `Set` are strictly increasing per
key, the slice stays sorted by timestamp. `Get` binary-searches that slice for
the largest timestamp `<= ts` and returns its value, or `""` if every entry is
newer than `ts` (or the key is unknown).

- Time:  Set O(1); Get O(log n) over the key's entries
- Space: O(total number of Set calls)
