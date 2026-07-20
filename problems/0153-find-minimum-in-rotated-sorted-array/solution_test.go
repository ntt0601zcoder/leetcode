package findminimuminrotatedsortedarray

import (
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(nums []int) int{
	"binarysearch": findMin,
}

// call runs fn with a timeout so a non-terminating solution (e.g. an infinite
// loop from bad bound updates like left=mid / right=mid) reports a clean
// failure instead of hanging the whole test run.
func call(fn func([]int) int, nums []int) (got int, timedOut bool) {
	done := make(chan int, 1)
	go func() {
		defer func() { _ = recover() }()
		done <- fn(nums)
	}()
	select {
	case got = <-done:
		return got, false
	case <-time.After(500 * time.Millisecond):
		return 0, true
	}
}

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "leetcode example 1", nums: []int{3, 4, 5, 1, 2}, want: 1},
		{name: "leetcode example 2", nums: []int{4, 5, 6, 7, 0, 1, 2}, want: 0},
		{name: "leetcode example 3 (not rotated)", nums: []int{11, 13, 15, 17}, want: 11},
		{name: "single element", nums: []int{1}, want: 1},
		{name: "two elements ascending", nums: []int{1, 2}, want: 1},
		{name: "two elements rotated", nums: []int{2, 1}, want: 1},
		{name: "min at last index", nums: []int{2, 3, 4, 5, 1}, want: 1},
		{name: "min at second index", nums: []int{5, 1, 2, 3, 4}, want: 1},
		{name: "min in middle", nums: []int{4, 5, 6, 1, 2, 3}, want: 1},
		{name: "already sorted long", nums: []int{1, 2, 3, 4, 5, 6, 7}, want: 1},
		{name: "rotated by one", nums: []int{7, 1, 2, 3, 4, 5, 6}, want: 1},
		{name: "rotated by n-1", nums: []int{2, 3, 4, 5, 6, 7, 1}, want: 1},
		{name: "negatives rotated", nums: []int{0, 1, 2, -3, -2, -1}, want: -3},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, timedOut := call(fn, append([]int(nil), tc.nums...))
				if timedOut {
					t.Fatalf("findMin(%v) did not return within 500ms (likely infinite loop), want %d", tc.nums, tc.want)
				}
				if got != tc.want {
					t.Errorf("findMin(%v) = %d, want %d", tc.nums, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkFindMin(b *testing.B) {
	// Large rotated array: [n/2 .. n-1, 0 .. n/2-1] so the minimum (0) sits
	// in the middle, exercising the full binary-search descent.
	const n = 10000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = (i + n/2) % n
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums)
			}
		})
	}
}
