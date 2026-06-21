# 739. Daily Temperatures

- Difficulty: medium
- Link: https://leetcode.com/problems/daily-temperatures/

## Approach

**Monotonic stack** (`dailyTemperatures`): stack giữ *chỉ số* các ngày đang
chờ ngày ấm hơn, nhiệt độ giảm dần từ đáy lên đỉnh. Gặp ngày `i` ấm hơn đỉnh
stack thì pop từng ngày `prev` và gán `result[prev] = i - prev`. Mỗi chỉ số
push/pop đúng một lần → O(n).

**Brute force** (`dailyTemperaturesBigOn`): với mỗi `i`, quét tới khi gặp ngày
ấm hơn. Đơn giản nhưng O(n²) (tên hàm "BigOn" dễ gây nhầm — thực chất O(n²)).

Lưu ý: chỉ tính ngày **ấm hơn hẳn** (so sánh `>`), nhiệt độ bằng nhau không
tính.

- Time:  O(n) stack / O(n²) brute
- Space: O(n)
