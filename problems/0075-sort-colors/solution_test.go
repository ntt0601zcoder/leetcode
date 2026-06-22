package sortcolors

import (
	"reflect"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
// Each sorts nums IN PLACE and returns nothing.
var solutions = map[string]func(nums []int){
	"dutchFlag": sortColors,
	"counting":  sortColorsW1,
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "leetcode example 1", nums: []int{2, 0, 2, 1, 1, 0}, want: []int{0, 0, 1, 1, 2, 2}},
		{name: "leetcode example 2", nums: []int{2, 0, 1}, want: []int{0, 1, 2}},
		{name: "single 0", nums: []int{0}, want: []int{0}},
		{name: "single 1", nums: []int{1}, want: []int{1}},
		{name: "single 2", nums: []int{2}, want: []int{2}},
		{name: "all zeros", nums: []int{0, 0, 0}, want: []int{0, 0, 0}},
		{name: "all ones", nums: []int{1, 1, 1}, want: []int{1, 1, 1}},
		{name: "all twos", nums: []int{2, 2, 2}, want: []int{2, 2, 2}},
		{name: "already sorted", nums: []int{0, 0, 1, 1, 2, 2}, want: []int{0, 0, 1, 1, 2, 2}},
		{name: "reverse sorted", nums: []int{2, 2, 1, 1, 0, 0}, want: []int{0, 0, 1, 1, 2, 2}},
		{name: "no zeros", nums: []int{2, 1, 2, 1}, want: []int{1, 1, 2, 2}},
		{name: "no ones", nums: []int{2, 0, 2, 0}, want: []int{0, 0, 2, 2}},
		{name: "no twos", nums: []int{1, 0, 1, 0}, want: []int{0, 0, 1, 1}},
		{name: "two elements swap", nums: []int{1, 0}, want: []int{0, 1}},
		{name: "mixed", nums: []int{1, 2, 0, 1, 2, 0, 1}, want: []int{0, 0, 1, 1, 1, 2, 2}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				in := append([]int(nil), tc.nums...)
				fn(in)
				if !reflect.DeepEqual(in, tc.want) {
					t.Errorf("%s(%v) sorted to %v, want %v", name, tc.nums, in, tc.want)
				}
			})
		}
	}
}

func BenchmarkSortColors(b *testing.B) {
	base := make([]int, 1000)
	for i := range base {
		base[i] = i % 3
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				nums := append([]int(nil), base...)
				fn(nums)
			}
		})
	}
}
