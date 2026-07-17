package kclosestpointstoorigin

import (
	"reflect"
	"sort"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(points [][]int, k int) [][]int{
	"heap": kClosest,
}

// call guards against a panic so a broken heap reports a clean failure
// instead of crashing the whole test binary.
func call(fn func([][]int, int) [][]int, points [][]int, k int) (got [][]int, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	return fn(points, k), false
}

// clone deep-copies the input so a solution can't affect later cases.
func clone(pts [][]int) [][]int {
	out := make([][]int, len(pts))
	for i, p := range pts {
		out[i] = append([]int(nil), p...)
	}
	return out
}

// normalize sorts the LIST of points so comparison ignores output order (the
// problem allows any order). It never sorts within a point — [x,y] order matters.
func normalize(pts [][]int) [][]int {
	out := clone(pts)
	sort.Slice(out, func(i, j int) bool {
		if out[i][0] != out[j][0] {
			return out[i][0] < out[j][0]
		}
		return out[i][1] < out[j][1]
	})
	return out
}

func TestKClosest(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{name: "leetcode example 1", points: [][]int{{1, 3}, {-2, 2}}, k: 1, want: [][]int{{-2, 2}}},
		{name: "leetcode example 2", points: [][]int{{3, 3}, {5, -1}, {-2, 4}}, k: 2, want: [][]int{{3, 3}, {-2, 4}}},
		{name: "single point", points: [][]int{{0, 1}}, k: 1, want: [][]int{{0, 1}}},
		{name: "k equals n", points: [][]int{{1, 1}, {2, 2}}, k: 2, want: [][]int{{1, 1}, {2, 2}}},
		{name: "origin wins", points: [][]int{{0, 0}, {1, 1}}, k: 1, want: [][]int{{0, 0}}},
		// These two catch a distance formula that isn't symmetric in x and y.
		{name: "y closer than x", points: [][]int{{0, 2}, {3, 0}}, k: 1, want: [][]int{{0, 2}}},
		{name: "x closer than y", points: [][]int{{5, 0}, {0, 4}}, k: 1, want: [][]int{{0, 4}}},
		{name: "negatives", points: [][]int{{-1, 0}, {2, 0}}, k: 1, want: [][]int{{-1, 0}}},
		{name: "keep the two closest", points: [][]int{{1, 0}, {2, 0}, {3, 0}, {4, 0}}, k: 2, want: [][]int{{1, 0}, {2, 0}}},
		{name: "mixed quadrants", points: [][]int{{-5, -5}, {1, 2}, {-1, -1}, {6, 0}}, k: 2, want: [][]int{{1, 2}, {-1, -1}}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, panicked := call(fn, clone(tc.points), tc.k)
				if panicked {
					t.Fatalf("kClosest(%v, %d) panicked, want %v", tc.points, tc.k, tc.want)
				}
				if len(got) != tc.k {
					t.Errorf("kClosest(%v, %d) returned %d points, want %d", tc.points, tc.k, len(got), tc.k)
				}
				if !reflect.DeepEqual(normalize(got), normalize(tc.want)) {
					t.Errorf("kClosest(%v, %d) = %v, want %v (any order)", tc.points, tc.k, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkKClosest(b *testing.B) {
	points := make([][]int, 1000)
	for i := range points {
		points[i] = []int{i%50 - 25, i%37 - 18}
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(clone(points), 50)
			}
		})
	}
}
