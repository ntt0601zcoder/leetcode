# 49. Group Anagrams

- Difficulty: medium
- Link: https://leetcode.com/problems/group-anagrams/

## Approach

Hiện tại: gom nhóm bằng cách so từng chuỗi với key đại diện của mỗi nhóm
(`isAnagramMap`). Đơn giản nhưng O(n²·L) vì phải so với mọi nhóm đã có.

Cách phổ biến hơn O(n·L·logL): dùng *chuỗi đã sort* (hoặc count 26 ký tự)
làm key của map → mỗi chuỗi tra key một lần.

Lưu ý test: thứ tự các nhóm trong kết quả là **không xác định** (duyệt map),
nên test so sánh sau khi chuẩn hoá (sort trong nhóm + sort các nhóm).

- Time:  O(n²·L) hiện tại / O(n·L·logL) nếu dùng sorted-key
- Space: O(n·L)
