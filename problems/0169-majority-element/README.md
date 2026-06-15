# 169. Majority Element

- Difficulty: easy
- Link: https://leetcode.com/problems/majority-element/

## Approach

Đếm tần suất bằng map; ngay khi một phần tử vượt `n/2` thì trả về luôn. Đề
đảm bảo phần tử đa số luôn tồn tại.

Tối ưu hơn: thuật toán bỏ phiếu Boyer-Moore — O(n) thời gian, O(1) bộ nhớ.

- Time:  O(n)
- Space: O(n) (Boyer-Moore: O(1))
