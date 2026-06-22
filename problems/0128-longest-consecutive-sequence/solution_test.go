package longestconsecutivesequence

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(nums []int) int{
	"hashSet": longestConsecutive,
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "leetcode example 1", nums: []int{100, 4, 200, 1, 3, 2}, want: 4},
		{name: "leetcode example 2", nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, want: 9},
		{name: "empty", nums: []int{}, want: 0},
		{name: "nil", nums: nil, want: 0},
		{name: "single", nums: []int{5}, want: 1},
		{name: "all duplicates", nums: []int{1, 1, 1, 1}, want: 1},
		{name: "duplicates within run", nums: []int{1, 2, 2, 3}, want: 3},
		{name: "unsorted full run", nums: []int{4, 2, 1, 3, 5}, want: 5},
		{name: "two separate runs", nums: []int{1, 2, 3, 10, 11}, want: 3},
		{name: "negatives", nums: []int{-3, -2, -1, 0, 1}, want: 5},
		{name: "negative and positive gap", nums: []int{-1, 1, 3}, want: 1},
		{name: "two element consecutive", nums: []int{2, 1}, want: 2},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := fn(append([]int(nil), tc.nums...))
				if got != tc.want {
					t.Errorf("longestConsecutive(%v) = %d, want %d", tc.nums, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkLongestConsecutive(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = (i * 7) % 1000 // scrambled but covers a long run
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums)
			}
		})
	}
}
