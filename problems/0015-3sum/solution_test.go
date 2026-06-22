package p3sum

import (
	"reflect"
	"sort"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(nums []int) [][]int{
	"twopointer": threeSum,
}

// normalize makes [][]int order-independent: sort within each triplet, then
// sort the list of triplets, so the test does not depend on the order in
// which solutions emit results.
func normalize(triplets [][]int) [][]int {
	out := make([][]int, len(triplets))
	for i, t := range triplets {
		c := append([]int(nil), t...)
		sort.Ints(c)
		out[i] = c
	}
	sort.Slice(out, func(i, j int) bool {
		for k := 0; k < len(out[i]) && k < len(out[j]); k++ {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return len(out[i]) < len(out[j])
	})
	return out
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "leetcode example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{name: "leetcode example 2 (no triplet)", nums: []int{0, 1, 1}, want: nil},
		{name: "leetcode example 3 (all zeros)", nums: []int{0, 0, 0}, want: [][]int{{0, 0, 0}}},
		{name: "empty", nums: []int{}, want: nil},
		{name: "single element", nums: []int{0}, want: nil},
		{name: "two elements", nums: []int{0, 0}, want: nil},
		{name: "no solution all positive", nums: []int{1, 2, 3}, want: nil},
		{
			name: "many duplicates dedup",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "negatives and positives",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			want: [][]int{
				{-4, -2, 6},
				{-4, 0, 4},
				{-4, 1, 3},
				{-4, 2, 2},
				{-2, -2, 4},
				{-2, 0, 2},
			},
		},
		{
			name: "four zeros only one triplet",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := normalize(fn(append([]int(nil), tc.nums...)))
				want := normalize(tc.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("threeSum(%v) = %v, want %v (order-independent)", tc.nums, got, want)
				}
			})
		}
	}
}

func BenchmarkThreeSum(b *testing.B) {
	nums := make([]int, 300)
	for i := range nums {
		nums[i] = i%21 - 10 // values in [-10, 10]
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(append([]int(nil), nums...))
			}
		})
	}
}
