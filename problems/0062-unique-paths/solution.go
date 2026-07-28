// Package uniquepaths solves LeetCode 62. Unique Paths.
// https://leetcode.com/problems/unique-paths/
package uniquepaths

import "math/big"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 0

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func uniquePathsS2(m int, n int) int {
	c := new(big.Int).Binomial(int64(m+n-2), int64(m-1))
	return int(c.Int64())
}
