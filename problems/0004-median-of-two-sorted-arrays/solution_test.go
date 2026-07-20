package medianoftwosortedarrays

import (
	"math"
	"testing"
	"time"
)

// solutions lists every approach; the table test runs all cases against each.
var solutions = map[string]func(nums1 []int, nums2 []int) float64{
	"default":      findMedianSortedArrays,
	"binarysearch": findMedianSortedArraysBinarySearch,
	"sort":         findMedianSortedArraysSort,
}

// call runs fn with a timeout so a non-terminating solution (e.g. an infinite
// loop from bad binary-search bound updates) reports a clean failure instead
// of hanging the whole test run. A panic is also recovered and surfaced.
func call(fn func([]int, []int) float64, nums1, nums2 []int) (got float64, timedOut, panicked bool) {
	done := make(chan float64, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
				done <- 0
			}
		}()
		done <- fn(nums1, nums2)
	}()
	select {
	case got = <-done:
		return got, false, panicked
	case <-time.After(500 * time.Millisecond):
		return 0, true, false
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	const eps = 1e-9
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{name: "leetcode example 1 (odd total)", nums1: []int{1, 3}, nums2: []int{2}, want: 2.0},
		{name: "leetcode example 2 (even total)", nums1: []int{1, 2}, nums2: []int{3, 4}, want: 2.5},
		{name: "first empty, one element", nums1: []int{}, nums2: []int{1}, want: 1.0},
		{name: "first empty, even total", nums1: []int{}, nums2: []int{2, 3}, want: 2.5},
		{name: "second empty, one element", nums1: []int{1}, nums2: []int{}, want: 1.0},
		{name: "both single", nums1: []int{1}, nums2: []int{2}, want: 1.5},
		{name: "both single equal", nums1: []int{5}, nums2: []int{5}, want: 5.0},
		{name: "negatives odd total", nums1: []int{-5, -3, -1}, nums2: []int{-2, 0}, want: -2.0},
		{name: "different lengths even total", nums1: []int{1, 2}, nums2: []int{3, 4, 5, 6}, want: 3.5},
		{name: "different lengths odd total", nums1: []int{3}, nums2: []int{1, 2, 4, 5}, want: 3.0},
		{name: "all equal", nums1: []int{2, 2}, nums2: []int{2, 2}, want: 2.0},
		{name: "interleaved even total", nums1: []int{1, 3, 5, 7}, nums2: []int{2, 4, 6, 8}, want: 4.5},
		{name: "disjoint ranges", nums1: []int{1, 2, 3}, nums2: []int{4, 5, 6}, want: 3.5},
		{name: "mixed sign odd total", nums1: []int{-1000000, 0}, nums2: []int{1000000}, want: 0.0},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, timedOut, panicked := call(
					fn,
					append([]int(nil), tc.nums1...),
					append([]int(nil), tc.nums2...),
				)
				switch {
				case timedOut:
					t.Fatalf("%s(%v, %v) did not return within 500ms (likely infinite loop), want %v",
						name, tc.nums1, tc.nums2, tc.want)
				case panicked:
					t.Fatalf("%s(%v, %v) panicked, want %v", name, tc.nums1, tc.nums2, tc.want)
				case math.Abs(got-tc.want) > eps:
					t.Errorf("%s(%v, %v) = %v, want %v", name, tc.nums1, tc.nums2, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	nums1 := make([]int, 500)
	nums2 := make([]int, 700)
	for i := range nums1 {
		nums1[i] = i * 2
	}
	for i := range nums2 {
		nums2[i] = i*2 + 1
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums1, nums2)
			}
		})
	}
}
