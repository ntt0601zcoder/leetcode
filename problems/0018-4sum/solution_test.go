package p4sum

import (
	"reflect"
	"sort"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(nums []int, target int) [][]int{
	"twopointer": fourSum,
}

// lessInts compares two int slices lexicographically.
func lessInts(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

// normalize makes [][]int order-independent: sort within each tuple, then
// sort the list of tuples. The answer may be returned in any order.
func normalize(quads [][]int) [][]int {
	out := make([][]int, len(quads))
	for i, q := range quads {
		c := append([]int(nil), q...)
		sort.Ints(c)
		out[i] = c
	}
	sort.Slice(out, func(i, j int) bool { return lessInts(out[i], out[j]) })
	return out
}

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		// The case you asked about: every quadruplet is [0,0,0,0], and the
		// answer must contain it exactly once.
		{name: "all zeros n=4", nums: []int{0, 0, 0, 0}, target: 0, want: [][]int{{0, 0, 0, 0}}},
		{name: "all zeros n=5", nums: []int{0, 0, 0, 0, 0}, target: 0, want: [][]int{{0, 0, 0, 0}}},
		{
			name: "leetcode example 1", nums: []int{1, 0, -1, 0, -2, 2}, target: 0,
			want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{name: "leetcode example 2 (dedup)", nums: []int{2, 2, 2, 2, 2}, target: 8, want: [][]int{{2, 2, 2, 2}}},
		{name: "four distinct", nums: []int{1, 2, 3, 4}, target: 10, want: [][]int{{1, 2, 3, 4}}},
		{name: "four identical", nums: []int{1, 1, 1, 1}, target: 4, want: [][]int{{1, 1, 1, 1}}},
		{name: "no quadruplet", nums: []int{1, 2, 3, 4}, target: 100, want: [][]int{}},
		{name: "fewer than four", nums: []int{1, 2, 3}, target: 6, want: [][]int{}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				// Copy: fourSum sorts nums in place.
				in := append([]int(nil), tc.nums...)
				got := normalize(fn(in, tc.target))
				want := normalize(tc.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("fourSum(%v, %d) = %v, want %v (any order)", tc.nums, tc.target, got, want)
				}
			})
		}
	}
}
