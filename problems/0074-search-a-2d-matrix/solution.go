// Package searcha2dmatrix solves LeetCode 74. Search A 2d Matrix.
// https://leetcode.com/problems/search-a-2d-matrix/
package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix[0]), len(matrix)
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2
		x, y := mid%m, mid/m

		switch {
		case matrix[y][x] == target:
			return true
		case matrix[y][x] > target:
			right = mid - 1
		case matrix[y][x] < target:
			left = mid + 1
		}
	}

	return false
}
