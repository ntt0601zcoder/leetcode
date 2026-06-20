// Package validparentheses solves LeetCode 20. Valid Parentheses.
// https://leetcode.com/problems/valid-parentheses/
package validparentheses

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))

	for _, c := range s {
		n := len(stack)
		switch {
		case '(' == c:
			stack = append(stack, c)
		case '[' == c:
			stack = append(stack, c)
		case '{' == c:
			stack = append(stack, c)
		case n != 0 && ')' == c && '(' == stack[n-1]:
			stack = stack[:n-1]
		case n != 0 && ']' == c && '[' == stack[n-1]:
			stack = stack[:n-1]
		case n != 0 && '}' == c && '{' == stack[n-1]:
			stack = stack[:n-1]
		default:
			return false
		}
	}
	return len(stack) == 0
}
