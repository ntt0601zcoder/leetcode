// Package interleavingstring solves LeetCode 97. Interleaving String.
// https://leetcode.com/problems/interleaving-string/
package interleavingstring

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)

	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i > 0 && s1[i-1] == s3[i+j-1] && dp[i-1][j] {
				dp[i][j] = true
			}
			if j > 0 && s2[j-1] == s3[i+j-1] && dp[i][j-1] {
				dp[i][j] = true
			}

		}
	}

	return dp[m][n]
}

func isInterleaveRec(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	var rec func(i, j int) bool
	rec = func(i, j int) bool {
		if len(s1) == i && len(s2) == j {
			return true
		}

		k := i + j
		if i < len(s1) && s1[i] == s3[k] && rec(i+1, j) {
			return true
		}
		if j < len(s2) && s2[j] == s3[k] && rec(i, j+1) {
			return true
		}

		return false
	}

	return rec(0, 0)
}
