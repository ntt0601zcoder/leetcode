package kokoeatingbananas

import (
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(piles []int, h int) int{
	"binarysearch": minEatingSpeed,
}

// call runs fn with a timeout so a non-terminating solution (e.g. a binary
// search whose switch has no branch to move the bounds and therefore spins
// forever) reports a clean failure instead of hanging the whole test run. A
// panic is caught too, so it surfaces as a timeout-style failure rather than
// crashing the whole test binary.
func call(fn func([]int, int) int, piles []int, h int) (got int, timedOut bool) {
	done := make(chan int, 1)
	go func() {
		defer func() { _ = recover() }()
		done <- fn(piles, h)
	}()
	select {
	case got = <-done:
		return got, false
	case <-time.After(500 * time.Millisecond):
		return 0, true
	}
}

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{name: "leetcode example 1", piles: []int{3, 6, 7, 11}, h: 8, want: 4},
		{name: "leetcode example 2", piles: []int{30, 11, 23, 4, 20}, h: 5, want: 30},
		{name: "leetcode example 3", piles: []int{30, 11, 23, 4, 20}, h: 6, want: 23},
		{name: "single pile, one hour", piles: []int{1000000000}, h: 1, want: 1000000000},
		{name: "single pile, ample hours", piles: []int{9}, h: 9, want: 1},
		{name: "single pile, exact division", piles: []int{12}, h: 4, want: 3},
		{name: "h equals len -> max pile", piles: []int{3, 6, 7, 11}, h: 4, want: 11},
		{name: "large h -> speed 1", piles: []int{3, 6, 7, 11}, h: 100, want: 1},
		{name: "all equal piles, h equals len", piles: []int{5, 5, 5, 5}, h: 4, want: 5},
		{name: "all ones, plenty of time", piles: []int{1, 1, 1, 1}, h: 100, want: 1},
		{name: "two piles tight", piles: []int{100, 200}, h: 3, want: 100},
		{name: "ceil boundary non-divisor speed", piles: []int{5, 5}, h: 4, want: 3},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, timedOut := call(fn, append([]int(nil), tc.piles...), tc.h)
				if timedOut {
					t.Fatalf("minEatingSpeed(%v, %d) did not return within 500ms (likely infinite loop or panic), want %d", tc.piles, tc.h, tc.want)
				}
				if got != tc.want {
					t.Errorf("minEatingSpeed(%v, %d) = %d, want %d", tc.piles, tc.h, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkMinEatingSpeed(b *testing.B) {
	piles := make([]int, 10000)
	for i := range piles {
		piles[i] = (i%997)*7 + 1
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(piles, 20000)
			}
		})
	}
}
