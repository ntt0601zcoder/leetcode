package findfirstandlastpositionofelementinsortedarray

import (
	"reflect"
	"testing"
	"time"
)

// solutions lists every approach; the table test runs all cases against each.
var solutions = map[string]func(nums []int, target int) []int{
	"binarysearch": searchRange,
	"bruteforce":   searchRangeBruteForce,
}

// call runs fn with a timeout so a non-terminating solution (e.g. an infinite
// loop from a bad binary-search bound update) or a panic reports a clean
// failure instead of hanging or crashing the whole test run.
func call(fn func([]int, int) []int, nums []int, target int) (got []int, timedOut bool) {
	done := make(chan []int, 1)
	go func() {
		defer func() { _ = recover() }()
		done <- fn(nums, target)
	}()
	select {
	case got = <-done:
		return got, false
	case <-time.After(500 * time.Millisecond):
		return nil, true
	}
}

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "leetcode example 1 (range)", nums: []int{5, 7, 7, 8, 8, 10}, target: 8, want: []int{3, 4}},
		{name: "leetcode example 2 (not found)", nums: []int{5, 7, 7, 8, 8, 10}, target: 6, want: []int{-1, -1}},
		{name: "leetcode example 3 (empty)", nums: []int{}, target: 0, want: []int{-1, -1}},
		{name: "single occurrence", nums: []int{5, 7, 7, 8, 8, 10}, target: 5, want: []int{0, 0}},
		{name: "single element found", nums: []int{1}, target: 1, want: []int{0, 0}},
		{name: "single element not found", nums: []int{1}, target: 2, want: []int{-1, -1}},
		{name: "all same", nums: []int{2, 2, 2, 2, 2}, target: 2, want: []int{0, 4}},
		{name: "target at start", nums: []int{3, 3, 5, 7, 9}, target: 3, want: []int{0, 1}},
		{name: "target at end", nums: []int{1, 2, 4, 6, 6, 6}, target: 6, want: []int{3, 5}},
		{name: "target between values", nums: []int{1, 3, 5, 7, 9}, target: 4, want: []int{-1, -1}},
		{name: "below range", nums: []int{2, 4, 6, 8}, target: 1, want: []int{-1, -1}},
		{name: "above range", nums: []int{2, 4, 6, 8}, target: 9, want: []int{-1, -1}},
		{name: "range in middle", nums: []int{1, 2, 3, 3, 3, 4, 5}, target: 3, want: []int{2, 4}},
		{name: "negatives", nums: []int{-5, -5, -3, 0, 0, 0, 2}, target: 0, want: []int{3, 5}},
		{name: "two elements both target", nums: []int{4, 4}, target: 4, want: []int{0, 1}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, timedOut := call(fn, append([]int(nil), tc.nums...), tc.target)
				if timedOut {
					t.Fatalf("searchRange(%v, %d) did not return within 500ms (likely infinite loop or panic), want %v", tc.nums, tc.target, tc.want)
				}
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("searchRange(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkSearchRange(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i / 3 // creates runs of equal values to exercise the range scan
	}
	target := 1500
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums, target)
			}
		})
	}
}
