// Package generateparentheses solves LeetCode 22. Generate Parentheses.
// https://leetcode.com/problems/generate-parentheses/
package generateparentheses

func generate(ans *[]string, s string, open int, close int, n int) {
	if open == n && close == n {
		*ans = append(*ans, s)
		return
	}

	if open > close {
		generate(ans, s+")", open, close+1, n)
	}

	if open < n {
		generate(ans, s+"(", open+1, close, n)
	}
}

func generateParenthesis(n int) []string {
	ans := []string{}
	generate(&ans, "", 0, 0, n)
	return ans
}
