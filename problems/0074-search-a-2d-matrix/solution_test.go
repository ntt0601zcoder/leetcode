package searcha2dmatrix

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(matrix [][]int, target int) bool{
	"binarysearch": searchMatrix,
}

// call guards against a panic (e.g. an out-of-range index from a bad
// flat-to-2D mapping) so the suite reports a clean failure instead of
// crashing the whole test binary.
func call(fn func([][]int, int) bool, matrix [][]int, target int) (got bool, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	return fn(matrix, target), false
}

func TestSearchMatrix(t *testing.T) {
	mat := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{name: "leetcode example 1 (found)", matrix: mat, target: 3, want: true},
		{name: "leetcode example 2 (not found)", matrix: mat, target: 13, want: false},
		{name: "first cell", matrix: mat, target: 1, want: true},
		{name: "last cell", matrix: mat, target: 60, want: true},
		{name: "below all", matrix: mat, target: 0, want: false},
		{name: "above all", matrix: mat, target: 100, want: false},
		{name: "row start (first column)", matrix: mat, target: 10, want: true},
		{name: "row end (last column)", matrix: mat, target: 20, want: true},
		{name: "middle", matrix: mat, target: 16, want: true},
		{name: "gap not found", matrix: mat, target: 17, want: false},
		{name: "single cell found", matrix: [][]int{{5}}, target: 5, want: true},
		{name: "single cell not found", matrix: [][]int{{5}}, target: 3, want: false},
		{name: "single row found", matrix: [][]int{{1, 3, 5}}, target: 5, want: true},
		{name: "single row not found", matrix: [][]int{{1, 3, 5}}, target: 4, want: false},
		{name: "single column found", matrix: [][]int{{1}, {3}, {5}}, target: 3, want: true},
		{name: "single column not found", matrix: [][]int{{1}, {3}, {5}}, target: 4, want: false},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, panicked := call(fn, tc.matrix, tc.target)
				if panicked {
					t.Fatalf("searchMatrix(%v, %d) panicked, want %v", tc.matrix, tc.target, tc.want)
				}
				if got != tc.want {
					t.Errorf("searchMatrix(%v, %d) = %v, want %v", tc.matrix, tc.target, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkSearchMatrix(b *testing.B) {
	const rows, cols = 100, 100
	mat := make([][]int, rows)
	for r := 0; r < rows; r++ {
		mat[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			mat[r][c] = (r*cols + c) * 2
		}
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(mat, 9998)
			}
		})
	}
}
