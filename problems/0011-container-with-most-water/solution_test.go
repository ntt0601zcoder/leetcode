package containerwithmostwater

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(height []int) int{
	"twopointer": maxArea,
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{name: "leetcode example 1", height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{name: "leetcode example 2", height: []int{1, 1}, want: 1},
		{name: "two zeros", height: []int{0, 0}, want: 0},
		{name: "ascending", height: []int{1, 2, 3, 4, 5}, want: 6},
		{name: "descending", height: []int{5, 4, 3, 2, 1}, want: 6},
		{name: "tallest at ends", height: []int{9, 1, 1, 1, 9}, want: 36},
		{name: "all equal", height: []int{4, 4, 4, 4}, want: 12},
		{name: "single zero in middle", height: []int{2, 0, 2}, want: 4},
		{name: "leading zero", height: []int{0, 2}, want: 0},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := fn(append([]int(nil), tc.height...))
				if got != tc.want {
					t.Errorf("maxArea(%v) = %d, want %d", tc.height, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkMaxArea(b *testing.B) {
	height := make([]int, 1000)
	for i := range height {
		height[i] = i % 97
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(height)
			}
		})
	}
}
