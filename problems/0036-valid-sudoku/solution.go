// Package validsudoku solves LeetCode 36. Valid Sudoku.
// https://leetcode.com/problems/valid-sudoku/
package validsudoku

func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9]int

	for rowIdx := 0; rowIdx < 9; rowIdx++ {
		for colIdx := 0; colIdx < 9; colIdx++ {
			c := board[rowIdx][colIdx]

			if c == '.' {
				continue
			}

			bit := 1 << (c - '1')
			b := (rowIdx/3)*3 + colIdx/3

			if (rows[rowIdx]&bit) != 0 || (cols[colIdx]&bit) != 0 || (boxes[b]&bit) != 0 {
				return false
			}

			rows[rowIdx] |= bit
			cols[colIdx] |= bit
			boxes[b] |= bit
		}
	}

	return true
}
