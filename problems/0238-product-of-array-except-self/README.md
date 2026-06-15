# 238. Product Of Array Except Self

- Difficulty: medium
- Link: https://leetcode.com/problems/product-of-array-except-self/

## Approach

Hai lượt, không dùng phép chia. Lượt xuôi: `answer[i]` = tích mọi phần tử
*bên trái* `i`. Lượt ngược: nhân thêm `suffix` = tích mọi phần tử *bên phải*
`i`. Mảng `answer` đóng luôn vai trò prefix nên chỉ tốn O(1) bộ nhớ phụ
(không tính mảng kết quả). Xử lý đúng cả khi có số 0.

- Time:  O(n)
- Space: O(1) phụ (ngoài mảng kết quả)
