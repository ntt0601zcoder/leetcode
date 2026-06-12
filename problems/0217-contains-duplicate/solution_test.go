package containsduplicate

import "testing"

// solutions lists every approach; the test runs all cases against each.
// Add another approach (e.g. sort-then-scan) by writing the func and
// adding a line here.
var solutions = map[string]func(nums []int) bool{
	"set": containsDuplicate,
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "has duplicate", nums: []int{1, 2, 3, 1}, want: true},
		{name: "all distinct", nums: []int{1, 2, 3, 4}, want: false},
		{name: "many duplicates", nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, want: true},
		{name: "empty", nums: []int{}, want: false},
		{name: "single", nums: []int{1}, want: false},
		{name: "two same", nums: []int{1, 1}, want: true},
		{name: "two distinct", nums: []int{1, 2}, want: false},
		{name: "duplicate at ends", nums: []int{5, 2, 3, 4, 5}, want: true},
		{name: "negatives with dup", nums: []int{-1, -2, -1}, want: true},
		{name: "negatives distinct", nums: []int{-1, -2, -3}, want: false},
		{name: "zero duplicate", nums: []int{0, 1, 0}, want: true},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				if got := fn(tc.nums); got != tc.want {
					t.Errorf("containsDuplicate(%v) = %v, want %v", tc.nums, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	// Worst case: all distinct, so the whole slice is scanned.
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums)
			}
		})
	}
}
