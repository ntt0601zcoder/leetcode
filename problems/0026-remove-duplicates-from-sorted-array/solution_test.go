package removeduplicatesfromsortedarray

import (
	"reflect"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(nums []int) int{
	"twopointers": removeDuplicates,
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int // expected dedup prefix nums[:k]
	}{
		{name: "leetcode example 1", nums: []int{1, 1, 2}, want: []int{1, 2}},
		{name: "leetcode example 2", nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: []int{0, 1, 2, 3, 4}},
		{name: "empty", nums: []int{}, want: []int{}},
		{name: "single", nums: []int{7}, want: []int{7}},
		{name: "all duplicates", nums: []int{2, 2, 2, 2}, want: []int{2}},
		{name: "no duplicates", nums: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
		{name: "two distinct", nums: []int{1, 2}, want: []int{1, 2}},
		{name: "two same", nums: []int{1, 1}, want: []int{1}},
		{name: "negatives", nums: []int{-3, -3, -1, 0, 0, 5}, want: []int{-3, -1, 0, 5}},
		{name: "trailing duplicates", nums: []int{1, 2, 3, 3, 3}, want: []int{1, 2, 3}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				nums := append([]int(nil), tc.nums...)
				k := fn(nums)
				if k != len(tc.want) {
					t.Errorf("%s(%v) returned k = %d, want %d", name, tc.nums, k, len(tc.want))
				}
				// Compare element-wise so an empty result (nil slice) matches an
				// empty expectation regardless of the nil-vs-empty distinction.
				got := nums[:k]
				if len(got) != len(tc.want) || (k > 0 && !reflect.DeepEqual(got, tc.want)) {
					t.Errorf("%s(%v): nums[:%d] = %v, want %v", name, tc.nums, k, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	base := make([]int, 1000)
	for i := range base {
		base[i] = i / 3 // each value repeats a few times, kept sorted
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
