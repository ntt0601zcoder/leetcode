package largestrectangleinhistogram

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(heights []int) int{
	"monotonicStack": largestRectangleArea,
}

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{name: "leetcode example 1", heights: []int{2, 1, 5, 6, 2, 3}, want: 10},
		{name: "leetcode example 2", heights: []int{2, 4}, want: 4},
		{name: "single bar", heights: []int{5}, want: 5},
		{name: "single zero", heights: []int{0}, want: 0},
		{name: "two equal bars", heights: []int{3, 3}, want: 6},
		{name: "increasing", heights: []int{1, 2, 3, 4, 5}, want: 9},
		{name: "decreasing", heights: []int{5, 4, 3, 2, 1}, want: 9},
		{name: "all equal heights", heights: []int{4, 4, 4, 4}, want: 16},
		{name: "valley", heights: []int{6, 2, 5, 4, 5, 1, 6}, want: 12},
		{name: "tall single bar", heights: []int{2, 1, 2}, want: 3},
		{name: "with zero splitting", heights: []int{2, 3, 0, 4, 5}, want: 8},
		{name: "plateau then drop", heights: []int{5, 5, 1, 1}, want: 10},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				in := append([]int(nil), tc.heights...)
				got := fn(in)
				if got != tc.want {
					t.Errorf("%s(%v) = %d, want %d", name, tc.heights, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkLargestRectangleArea(b *testing.B) {
	heights := make([]int, 1000)
	for i := range heights {
		heights[i] = (i * 37) % 101
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(heights)
			}
		})
	}
}
