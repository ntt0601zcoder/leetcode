package trappingrainwater

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(height []int) int{
	"fullScan":     trapFullScan,
	"prefixSuffix": trapPrefixSuffix,
	"twoPointers":  trapTwoPointers,
}

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{name: "leetcode example 1", height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, want: 6},
		{name: "leetcode example 2", height: []int{4, 2, 0, 3, 2, 5}, want: 9},
		{name: "empty", height: []int{}, want: 0},
		{name: "single bar", height: []int{5}, want: 0},
		{name: "two bars", height: []int{3, 2}, want: 0},
		{name: "flat", height: []int{2, 2, 2, 2}, want: 0},
		{name: "ascending traps nothing", height: []int{1, 2, 3, 4, 5}, want: 0},
		{name: "descending traps nothing", height: []int{5, 4, 3, 2, 1}, want: 0},
		{name: "single peak traps nothing", height: []int{0, 1, 2, 3, 2, 1, 0}, want: 0},
		{name: "single valley", height: []int{3, 0, 3}, want: 3},
		{name: "v shape", height: []int{5, 0, 5}, want: 5},
		{name: "all zeros", height: []int{0, 0, 0}, want: 0},
		{name: "container", height: []int{5, 1, 1, 1, 5}, want: 12},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				in := append([]int(nil), tc.height...)
				got := fn(in)
				if got != tc.want {
					t.Errorf("%s(%v) = %d, want %d", name, tc.height, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkTrap(b *testing.B) {
	height := make([]int, 1000)
	for i := range height {
		if i%2 == 0 {
			height[i] = i % 7
		} else {
			height[i] = 9 - (i % 5)
		}
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(height)
			}
		})
	}
}
