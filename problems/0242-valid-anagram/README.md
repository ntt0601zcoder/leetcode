# 242. Valid Anagram

- Difficulty: easy
- Link: https://leetcode.com/problems/valid-anagram/

## Approach

Đếm tần suất ký tự bằng `map[rune]int`: cộng cho `s`, trừ cho `t`, xoá khi
về 0. Nếu gặp ký tự không có (hoặc dư) thì không phải anagram. Vì dùng
`rune` nên xử lý được cả Unicode (follow-up của bài).

- Time:  O(n)
- Space: O(k) với k là số ký tự phân biệt
