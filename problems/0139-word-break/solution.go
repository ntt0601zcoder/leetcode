// Package wordbreak solves LeetCode 139. Word Break.
// https://leetcode.com/problems/word-break/
package wordbreak

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool, len(wordDict))

	for _, work := range wordDict {
		wordSet[work] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
