package mincosttoconnectallpoints

import (
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(points [][]int) int{
	"prim": minCostConnectPointsPrim,
}

func clonePoints(pts [][]int) [][]int {
	out := make([][]int, len(pts))
	for i, p := range pts {
		out[i] = append([]int(nil), p...)
	}
	return out
}

// call guards against a panic or a runaway loop (heap-based), reporting a
// clean failure instead of crashing/hanging the test run.
func call(fn func([][]int) int, points [][]int) (got int, outcome string) {
	done := make(chan int, 1)
	panicked := make(chan struct{}, 1)
	go func() {
		defer func() {
			if recover() != nil {
				panicked <- struct{}{}
			}
		}()
		done <- fn(points)
	}()
	select {
	case got = <-done:
		return got, "ok"
	case <-panicked:
		return 0, "panic"
	case <-time.After(time.Second):
		return 0, "timeout"
	}
}

func TestMinCostConnectPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{name: "leetcode example 1", points: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}, want: 20},
		{name: "leetcode example 2", points: [][]int{{3, 12}, {-2, 5}, {-4, 1}}, want: 18},
		{name: "single point", points: [][]int{{0, 0}}, want: 0},
		{name: "empty", points: [][]int{}, want: 0},
		{name: "two points", points: [][]int{{0, 0}, {1, 1}}, want: 2},
		{name: "two points far", points: [][]int{{0, 0}, {3, 4}}, want: 7},
		{name: "collinear", points: [][]int{{0, 0}, {1, 0}, {2, 0}}, want: 2},
		{name: "unit square", points: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, want: 3},
		{name: "square around origin", points: [][]int{{-1, -1}, {1, 1}, {-1, 1}, {1, -1}}, want: 6},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, outcome := call(fn, clonePoints(tc.points))
				switch outcome {
				case "panic":
					t.Fatalf("minCostConnectPoints(%v) panicked, want %d", tc.points, tc.want)
				case "timeout":
					t.Fatalf("minCostConnectPoints(%v) did not return within 1s, want %d", tc.points, tc.want)
				}
				if got != tc.want {
					t.Errorf("minCostConnectPoints(%v) = %d, want %d", tc.points, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkMinCostConnectPoints(b *testing.B) {
	points := make([][]int, 100)
	for i := range points {
		points[i] = []int{i%20 - 10, (i*7)%20 - 10}
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(clonePoints(points))
			}
		})
	}
}
