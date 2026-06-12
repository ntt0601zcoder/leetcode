# 217. Contains Duplicate

- Difficulty: easy
- Link: https://leetcode.com/problems/contains-duplicate/

## Approach

Dùng một set (`map[int]...`): duyệt từng phần tử, nếu đã có trong set thì
trả về `true`, ngược lại thêm vào. Hết mảng mà không trùng thì `false`.

- Time:  O(n)
- Space: O(n)
