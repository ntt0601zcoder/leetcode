# 150. Evaluate Reverse Polish Notation

- Difficulty: medium
- Link: https://leetcode.com/problems/evaluate-reverse-polish-notation/

## Approach

Duyệt token với một stack: gặp **số** thì push; gặp **toán tử** thì pop hai
phần tử trên cùng (`num1` sâu hơn, `num2` ở đỉnh), tính `num1 op num2`, push
kết quả lại. Cuối cùng đỉnh stack là đáp án.

Lưu ý: thứ tự toán hạng quan trọng với `-` và `/` (num1 op num2, không phải
ngược lại); phép chia **cắt về 0** — đúng với phép `/` int của Go, vd
`-7 / 3 == -2`.

- Time:  O(n)
- Space: O(n)
