// Package longestincreasingsubsequence solves LeetCode 300. Longest Increasing Subsequence.
// https://leetcode.com/problems/longest-increasing-subsequence/
package longestincreasingsubsequence

import "sort"

func lengthOfLIS(nums []int) int {
	result := []int{}

	for _, n := range nums {
		pos := sort.SearchInts(result, n)

		if pos == len(result) {
			result = append(result, n)
		} else {
			result[pos] = n
		}
	}

	return len(result)
}

func lengthOfLISSolution2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	result := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}

	return result
}
