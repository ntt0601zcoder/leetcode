package searchinrotatedsortedarray

import (
	"testing"
	"time"
)

// solutions lists every approach; the table test runs all cases against each.
var solutions = map[string]func(nums []int, target int) int{
	"binarysearch": search,
}

// call runs fn in a goroutine guarded by a timeout so a non-terminating
// solution (e.g. an infinite loop from bad bound updates like left=mid /
// right=mid) or a panic reports a clean failure instead of hanging or
// crashing the whole test run.
func call(fn func([]int, int) int, nums []int, target int) (got int, timedOut bool) {
	done := make(chan int, 1)
	go func() {
		defer func() { _ = recover() }()
		done <- fn(nums, target)
	}()
	select {
	case got = <-done:
		return got, false
	case <-time.After(500 * time.Millisecond):
		return 0, true
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		// LeetCode examples.
		{name: "example 1 rotated found", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{name: "example 2 rotated not found", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
		{name: "example 3 single not found", nums: []int{1}, target: 0, want: -1},
		// Single element.
		{name: "single found", nums: []int{5}, target: 5, want: 0},
		// Not rotated (fully sorted).
		{name: "not rotated found", nums: []int{1, 2, 3, 4, 5, 6, 7}, target: 4, want: 3},
		{name: "not rotated not found", nums: []int{1, 2, 3, 4, 5, 6, 7}, target: 8, want: -1},
		// Target at the pivot (largest element before the wrap) and the min.
		{name: "target at max before pivot", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 7, want: 3},
		{name: "target is pivot (min) value", nums: []int{6, 7, 0, 1, 2, 4, 5}, target: 0, want: 2},
		// First and last element of a rotated array.
		{name: "first element rotated", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 4, want: 0},
		{name: "last element rotated", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 2, want: 6},
		// Target smaller / larger than every element.
		{name: "smaller than all", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: -1, want: -1},
		{name: "larger than all", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 100, want: -1},
		// Two elements, rotated.
		{name: "two elements rotated first", nums: []int{2, 1}, target: 2, want: 0},
		{name: "two elements rotated second", nums: []int{2, 1}, target: 1, want: 1},
		// Rotated by one (pivot at index 1); target lives in the sorted tail.
		{name: "rotate by one found in tail", nums: []int{5, 1, 2, 3, 4}, target: 3, want: 3},
		// Negative values, rotated.
		{name: "negatives rotated found", nums: []int{-1, 0, 1, 2, -5, -4, -3, -2}, target: -4, want: 5},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				// Copy the input in case a solution mutates it in place.
				got, timedOut := call(fn, append([]int(nil), tc.nums...), tc.target)
				if timedOut {
					t.Fatalf("search(%v, %d) did not return within 500ms (likely infinite loop), want %d", tc.nums, tc.target, tc.want)
				}
				if got != tc.want {
					t.Errorf("search(%v, %d) = %d, want %d", tc.nums, tc.target, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	// A large rotated array: [n/2 .. n-1, 0 .. n/2-1].
	const n = 10000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = (i + n/2) % n
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums, 1)
			}
		})
	}
}
