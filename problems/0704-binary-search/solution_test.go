package binarysearch

import (
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(nums []int, target int) int{
	"binarysearch": search,
}

// call runs fn with a timeout so a non-terminating solution (e.g. an infinite
// loop from bad bound updates like right=mid / left=mid) reports a clean
// failure instead of hanging the whole test run.
func call(fn func([]int, int) int, nums []int, target int) (got int, timedOut bool) {
	done := make(chan int, 1)
	go func() {
		defer func() { _ = recover() }()
		done <- fn(nums, target)
	}()
	select {
	case got = <-done:
		return got, false
	case <-time.After(time.Second):
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
		{name: "leetcode example 1 (found)", nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, want: 4},
		{name: "leetcode example 2 (not found)", nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, want: -1},
		{name: "single found", nums: []int{5}, target: 5, want: 0},
		{name: "single not found", nums: []int{5}, target: 3, want: -1},
		{name: "first element", nums: []int{1, 2, 3, 4, 5}, target: 1, want: 0},
		{name: "last element", nums: []int{1, 2, 3, 4, 5}, target: 5, want: 4},
		{name: "two elements, second", nums: []int{1, 2}, target: 2, want: 1},
		{name: "below range", nums: []int{2, 4, 6}, target: 1, want: -1},
		{name: "above range", nums: []int{2, 4, 6}, target: 7, want: -1},
		{name: "even length middle", nums: []int{1, 2, 3, 4}, target: 3, want: 2},
		{name: "negatives", nums: []int{-5, -3, -1, 0, 2}, target: -3, want: 1},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, timedOut := call(fn, append([]int(nil), tc.nums...), tc.target)
				if timedOut {
					t.Fatalf("search(%v, %d) did not return within 1s (likely infinite loop), want %d", tc.nums, tc.target, tc.want)
				}
				if got != tc.want {
					t.Errorf("search(%v, %d) = %d, want %d", tc.nums, tc.target, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i * 2
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums, 9998)
			}
		})
	}
}
