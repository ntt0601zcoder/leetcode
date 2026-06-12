# 347. Top K Frequent Elements

- Difficulty: medium
- Link: https://leetcode.com/problems/top-k-frequent-elements/

## Approach

Đếm tần suất bằng map, lấy danh sách các phần tử phân biệt, sort giảm dần
theo tần suất rồi lấy `k` phần tử đầu. (Có thể tối ưu O(n) bằng bucket sort
theo tần suất, hoặc dùng heap kích thước k.)

Lưu ý: đề cho trả về theo **thứ tự bất kỳ** và đáp án **duy nhất**, nên test
so sánh không phụ thuộc thứ tự (sort hai bên rồi so).

- Time:  O(n + m·log m) với m = số phần tử phân biệt (O(n) nếu bucket sort)
- Space: O(m)
