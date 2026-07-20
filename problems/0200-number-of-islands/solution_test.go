package numberofislands

import (
	"testing"
)

func toGrid(rows ...string) [][]byte {
	grid := make([][]byte, len(rows))
	for i, r := range rows {
		grid[i] = []byte(r)
	}
	return grid
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "nil grid",
			grid: nil,
			want: 0,
		},
		{
			name: "một ô nước",
			grid: toGrid("0"),
			want: 0,
		},
		{
			name: "một ô đất",
			grid: toGrid("1"),
			want: 1,
		},
		{
			name: "toàn nước",
			grid: toGrid(
				"000",
				"000",
				"000",
			),
			want: 0,
		},
		{
			name: "toàn đất = một đảo lớn",
			grid: toGrid(
				"111",
				"111",
				"111",
			),
			want: 1,
		},
		{
			name: "LeetCode ví dụ 1",
			grid: toGrid(
				"11110",
				"11010",
				"11000",
				"00000",
			),
			want: 1,
		},
		{
			name: "LeetCode ví dụ 2 (3 đảo)",
			grid: toGrid(
				"11000",
				"11000",
				"00100",
				"00011",
			),
			want: 3,
		},
		{
			name: "chéo KHÔNG nối (2 đảo)", // ô chéo không tính là kề
			grid: toGrid(
				"10",
				"01",
			),
			want: 2,
		},
		{
			name: "bàn cờ vua (mỗi ô đất là 1 đảo)",
			grid: toGrid(
				"101",
				"010",
				"101",
			),
			want: 5,
		},
		{
			name: "vòng khép kín có lỗ giữa = 1 đảo", // lỗ ở giữa là nước, viền vẫn liên thông
			grid: toGrid(
				"111",
				"101",
				"111",
			),
			want: 1,
		},
		{
			name: "một hàng duy nhất",
			grid: toGrid("10101"),
			want: 3,
		},
		{
			name: "một cột duy nhất",
			grid: toGrid("1", "0", "1"),
			want: 2,
		},
		{
			name: "hình chữ U (nối vòng dưới)",
			grid: toGrid(
				"10001",
				"10001",
				"11111",
			),
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := numIslands(tc.grid)
			if got != tc.want {
				t.Errorf("numIslands() = %d, muốn %d\ngrid:\n%s", got, tc.want, gridString(tc.grid))
			}
		})
	}
}

// gridString: in grid ra khi test fail, giúp debug nhanh.
func gridString(grid [][]byte) string {
	s := ""
	for _, row := range grid {
		s += string(row) + "\n"
	}
	return s
}
