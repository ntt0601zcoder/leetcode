package majorityelement

import "testing"

// solutions lists every approach; the test runs all cases against each.
// (e.g. add a Boyer-Moore voting version and put it here too.)
var solutions = map[string]func(nums []int) int{
	"count": majorityElement,
}

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "basic", nums: []int{3, 2, 3}, want: 3},
		{name: "leetcode example", nums: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
		{name: "single", nums: []int{1}, want: 1},
		{name: "all same", nums: []int{5, 5, 5, 5}, want: 5},
		{name: "two same", nums: []int{4, 4}, want: 4},
		{name: "majority at edges", nums: []int{6, 1, 6, 6, 2, 6, 6}, want: 6},
		{name: "negatives", nums: []int{-1, -1, -1, 2}, want: -1},
		{name: "majority is zero", nums: []int{0, 0, 1}, want: 0},
		{name: "majority is last seen", nums: []int{1, 2, 2}, want: 2},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				in := append([]int(nil), tc.nums...)
				if got := fn(in); got != tc.want {
					t.Errorf("majorityElement(%v) = %d, want %d", tc.nums, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkMajorityElement(b *testing.B) {
	nums := make([]int, 1001)
	for i := range nums {
		nums[i] = 7 // overwhelming majority
	}
	nums[0], nums[1] = 1, 2
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums)
			}
		})
	}
}
