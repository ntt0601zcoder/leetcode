# 18. 4sum

- Difficulty: medium
- Link: https://leetcode.com/problems/4sum/

## Approach

Sort mảng, cố định 2 chỉ số ngoài `i`, `j` (vòng lặp lồng), rồi dùng hai con
trỏ `k`/`h` tìm cặp còn lại cho đủ `target`. Để kết quả **không trùng**: bỏ
qua `nums[i]`, `nums[j]` lặp lại, và sau khi tìm thấy một bộ thì nhảy qua các
`nums[k]`, `nums[h]` bằng nhau.

Lưu ý: kết quả trả về theo **thứ tự bất kỳ**, nên test chuẩn hoá (sort trong
mỗi bộ + sort danh sách) trước khi so.

- Time:  O(n^3)
- Space: O(1) phụ (ngoài mảng kết quả; chưa tính chi phí sort)
