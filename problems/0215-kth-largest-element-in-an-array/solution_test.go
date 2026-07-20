package kthlargestelementinanarray

import (
	"testing"
	"time"
)

// solutions lists every approach; the table test runs all cases against each.
var solutions = map[string]func(nums []int, k int) int{
	"minheap": findKthLargest,
}

// call runs fn with a timeout so a non-terminating solution (e.g. a heap that
// never makes progress) reports a clean failure instead of hanging the run.
// It also recovers panics (e.g. index-out-of-range on an empty heap).
func call(fn func([]int, int) int, nums []int, k int) (got int, timedOut, panicked bool) {
	done := make(chan int, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
				done <- 0
			}
		}()
		done <- fn(nums, k)
	}()
	select {
	case got = <-done:
		return got, false, panicked
	case <-time.After(500 * time.Millisecond):
		return 0, true, false
	}
}

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "leetcode example 1", nums: []int{3, 2, 1, 5, 6, 4}, k: 2, want: 5},
		{name: "leetcode example 2", nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4, want: 4},
		{name: "single element", nums: []int{1}, k: 1, want: 1},
		{name: "k=1 is the max", nums: []int{7, 10, 4, 3, 20, 15}, k: 1, want: 20},
		{name: "k=len is the min", nums: []int{7, 10, 4, 3, 20, 15}, k: 6, want: 3},
		{name: "all duplicates", nums: []int{2, 2, 2, 2}, k: 3, want: 2},
		{name: "duplicated max k=2", nums: []int{5, 5, 5, 1}, k: 2, want: 5},
		{name: "negatives k=1", nums: []int{-1, -2, -3, -4}, k: 1, want: -1},
		{name: "negatives k=len", nums: []int{-1, -2, -3, -4}, k: 4, want: -4},
		{name: "mixed signs middle", nums: []int{-5, 3, 0, -1, 7, 2}, k: 3, want: 2},
		{name: "two elements k=2", nums: []int{2, 1}, k: 2, want: 1},
		{name: "sorted ascending k=3", nums: []int{1, 2, 3, 4, 5}, k: 3, want: 3},
		{name: "sorted descending k=4", nums: []int{5, 4, 3, 2, 1}, k: 4, want: 2},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, timedOut, panicked := call(fn, append([]int(nil), tc.nums...), tc.k)
				if timedOut {
					t.Fatalf("findKthLargest(%v, %d) did not return within 500ms (likely infinite loop), want %d", tc.nums, tc.k, tc.want)
				}
				if panicked {
					t.Fatalf("findKthLargest(%v, %d) panicked, want %d", tc.nums, tc.k, tc.want)
				}
				if got != tc.want {
					t.Errorf("findKthLargest(%v, %d) = %d, want %d", tc.nums, tc.k, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = (i * 7919) % 10007
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(append([]int(nil), nums...), 100)
			}
		})
	}
}
