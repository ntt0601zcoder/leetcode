package productofarrayexceptself

import (
	"reflect"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(nums []int) []int{
	"prefixsuffix": productExceptSelf,
}

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "basic", nums: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{name: "two elements", nums: []int{3, 7}, want: []int{7, 3}},
		{name: "all negative", nums: []int{-1, -2, -3, -4}, want: []int{-24, -12, -8, -6}},
		{name: "single zero", nums: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
		{name: "zero at end", nums: []int{1, 2, 3, 0}, want: []int{0, 0, 0, 6}},
		{name: "two zeros", nums: []int{0, 0}, want: []int{0, 0}},
		{name: "mixed signs", nums: []int{2, 3, -2, 4}, want: []int{-24, -16, 24, -12}},
		{name: "with ones", nums: []int{1, 1, 1, 1}, want: []int{1, 1, 1, 1}},
		// n == 1 is below LeetCode's constraint (n >= 2) but the algorithm
		// handles it: product of no other elements is 1.
		{name: "single element", nums: []int{5}, want: []int{1}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				// Copy the input so an in-place solution can't affect later cases.
				in := append([]int(nil), tc.nums...)
				got := fn(in)
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("productExceptSelf(%v) = %v, want %v", tc.nums, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = (i % 9) + 1 // 1..9, avoid zeros and overflow
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums)
			}
		})
	}
}
