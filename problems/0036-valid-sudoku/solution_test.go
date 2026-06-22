package validsudoku

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(board [][]byte) bool{
	"bitmask": isValidSudoku,
}

// parse turns a 9-row string board into [][]byte for convenient test fixtures.
func parse(rows [9]string) [][]byte {
	board := make([][]byte, 9)
	for i, r := range rows {
		board[i] = []byte(r)
	}
	return board
}

func TestIsValidSudoku(t *testing.T) {
	valid := parse([9]string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	})
	// Same board but top-left '5' changed to '8', creating a duplicate '8' in
	// the first column (and its 3x3 box).
	invalidCol := parse([9]string{
		"83..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	})
	// Duplicate within the very first row.
	invalidRow := parse([9]string{
		"55..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	})
	// Duplicate inside the top-left 3x3 box: two '5's in the box that share
	// neither a row nor a column, so only the box check can catch it.
	invalidBox := parse([9]string{
		"5........",
		".5.......",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
	})
	empty := parse([9]string{
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
		".........",
	})

	tests := []struct {
		name  string
		board [][]byte
		want  bool
	}{
		{name: "leetcode example 1 valid", board: valid, want: true},
		{name: "leetcode example 2 invalid column", board: invalidCol, want: false},
		{name: "invalid row", board: invalidRow, want: false},
		{name: "invalid box", board: invalidBox, want: false},
		{name: "all empty", board: empty, want: true},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				if got := fn(tc.board); got != tc.want {
					t.Errorf("%s(%s) = %v, want %v", name, tc.name, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkIsValidSudoku(b *testing.B) {
	board := parse([9]string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	})
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(board)
			}
		})
	}
}
