package movezeroes

import (
	"reflect"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
// moveZeroes mutates its slice in place and returns nothing.
var solutions = map[string]func(nums []int){
	"twopointer": moveZeroes,
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int // zeroes moved to the end, non-zero order preserved
	}{
		{name: "leetcode example 1", nums: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		{name: "leetcode example 2", nums: []int{0}, want: []int{0}},
		{name: "all zeros", nums: []int{0, 0, 0}, want: []int{0, 0, 0}},
		{name: "no zeros", nums: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "single non-zero", nums: []int{5}, want: []int{5}},
		{name: "empty", nums: []int{}, want: []int{}},
		{name: "leading zeros", nums: []int{0, 0, 1, 2}, want: []int{1, 2, 0, 0}},
		{name: "trailing zeros stay", nums: []int{1, 2, 0, 0}, want: []int{1, 2, 0, 0}},
		{name: "interleaved", nums: []int{0, 1, 0, 2, 0, 3}, want: []int{1, 2, 3, 0, 0, 0}},
		{name: "negatives preserved", nums: []int{0, -1, 0, -2}, want: []int{-1, -2, 0, 0}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := make([]int, len(tc.nums))
				copy(got, tc.nums)
				fn(got)
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("moveZeroes(%v) -> %v, want %v", tc.nums, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	base := make([]int, 1000)
	for i := range base {
		if i%2 == 0 {
			base[i] = 0
		} else {
			base[i] = i
		}
	}
	buf := make([]int, len(base))
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				copy(buf, base)
				fn(buf)
			}
		})
	}
}
