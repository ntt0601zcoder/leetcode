package cheapestflightswithinkstops

import "testing"

func TestFindCheapestPrice(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		flights     [][]int
		src, dst, k int
		want        int
	}{{
		name:    "LeetCode ví dụ 1 (k=1 → đường 1 stop)",
		n:       4,
		flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}},
		src:     0, dst: 3, k: 1,
		want: 700,
	},
		{
			name:    "cùng đồ thị nhưng k=2 → đường rẻ hơn được phép",
			n:       4,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}},
			src:     0, dst: 3, k: 2,
			want: 400,
		},
		{
			name:    "LeetCode ví dụ 2 (k=1)",
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0, dst: 2, k: 1,
			want: 200,
		},
		{
			name:    "LeetCode ví dụ 3 (k=0 → chỉ đường trực tiếp)",
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0, dst: 2, k: 0,
			want: 500,
		},
		{
			name:    "không có đường tới đích",
			n:       3,
			flights: [][]int{{0, 1, 100}},
			src:     0, dst: 2, k: 5,
			want: -1,
		},
		{
			name:    "src trùng dst",
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}},
			src:     0, dst: 0, k: 1,
			want: 0,
		},
		{
			name:    "đường bay trực tiếp, k=0",
			n:       2,
			flights: [][]int{{0, 1, 50}},
			src:     0, dst: 1, k: 0,
			want: 50,
		},
		{
			name:    "một node duy nhất",
			n:       1,
			flights: [][]int{},
			src:     0, dst: 0, k: 0,
			want: 0,
		},
		{
			name:    "chuỗi dài rẻ bị chặn bởi k (k=1)",
			n:       4,
			flights: [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 10}},
			src:     0, dst: 3, k: 1,
			want: 10,
		},
		{
			name:    "chuỗi dài rẻ được phép (k=2)",
			n:       4,
			flights: [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 10}},
			src:     0, dst: 3, k: 2,
			want: 3,
		},
		{
			name:    "chuỗi 4 cạnh bị chặn (k=2)",
			n:       5,
			flights: [][]int{{0, 1, 5}, {1, 2, 5}, {2, 3, 5}, {3, 4, 5}, {0, 4, 100}},
			src:     0, dst: 4, k: 2,
			want: 100,
		},
		{
			name:    "chuỗi 4 cạnh được phép (k=3)",
			n:       5,
			flights: [][]int{{0, 1, 5}, {1, 2, 5}, {2, 3, 5}, {3, 4, 5}, {0, 4, 100}},
			src:     0, dst: 4, k: 3,
			want: 20,
		},
		{
			name:    "đích không tới được trong giới hạn k",
			n:       4,
			flights: [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}},
			src:     0, dst: 3, k: 1,
			want: -1,
		}}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findCheapestPrice(tc.n, tc.flights, tc.src, tc.dst, tc.k)

			if got != tc.want {
				t.Errorf("findCheapestPrice(n=%d, src=%d, dst=%d, k=%d) = %d, muốn %d",
					tc.n, tc.src, tc.dst, tc.k, got, tc.want)
			}
		})
	}
}
