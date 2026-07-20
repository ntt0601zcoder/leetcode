package findpeakelement

import (
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(nums []int) int{
	"binarysearch": findPeakElement,
}

// call runs fn with a timeout so a non-terminating solution (e.g. an infinite
// loop from bad bound updates in the binary search) reports a clean failure
// instead of hanging the whole test run. A panic (e.g. an out-of-bounds index)
// is caught by recover and surfaces as got=0, timedOut=false.
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

// isPeak reports whether index i is a valid peak of nums: strictly greater than
// both neighbors, with out-of-bounds neighbors treated as -infinity. The answer
// to this problem is NOT unique (several indices may be valid peaks), so the
// test asserts this property rather than comparing against a fixed index.
func isPeak(nums []int, i int) bool {
	n := len(nums)
	if i < 0 || i >= n {
		return false
	}
	if i > 0 && nums[i] <= nums[i-1] {
		return false
	}
	if i < n-1 && nums[i] <= nums[i+1] {
		return false
	}
	return true
}

func TestFindPeakElement(t *testing.T) {
	// Constraints: adjacent elements are always distinct (nums[i] != nums[i+1]).
	tests := []struct {
		name string
		nums []int
	}{
		{name: "single element", nums: []int{1}},
		{name: "two increasing (peak at last)", nums: []int{1, 2}},
		{name: "two decreasing (peak at first)", nums: []int{2, 1}},
		{name: "leetcode example 1", nums: []int{1, 2, 3, 1}},
		{name: "leetcode example 2 (multiple peaks)", nums: []int{1, 2, 1, 3, 5, 6, 4}},
		{name: "strictly increasing (peak at last)", nums: []int{1, 2, 3, 4, 5}},
		{name: "strictly decreasing (peak at first)", nums: []int{5, 4, 3, 2, 1}},
		{name: "interior peak", nums: []int{1, 3, 2}},
		{name: "peak just after start", nums: []int{1, 4, 3, 2}},
		{name: "valley between two peaks", nums: []int{5, 4, 3, 4, 5}},
		{name: "mountain", nums: []int{1, 2, 3, 4, 3, 2, 1}},
		{name: "negatives increasing", nums: []int{-3, -2, -1}},
		{name: "peak at index 0 with tail", nums: []int{2, 1, 2}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, timedOut := call(fn, append([]int(nil), tc.nums...))
				if timedOut {
					t.Fatalf("findPeakElement(%v) did not return within 500ms (likely infinite loop or panic)", tc.nums)
				}
				if !isPeak(tc.nums, got) {
					t.Errorf("findPeakElement(%v) = %d, which is not a valid peak", tc.nums, got)
				}
			})
		}
	}
}

func BenchmarkFindPeakElement(b *testing.B) {
	// A single mountain: strictly up to the middle, then strictly down.
	nums := make([]int, 10001)
	mid := len(nums) / 2
	for i := range nums {
		if i <= mid {
			nums[i] = i
		} else {
			nums[i] = len(nums) - i
		}
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums)
			}
		})
	}
}
