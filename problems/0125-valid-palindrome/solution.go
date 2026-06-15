// Package validpalindrome solves LeetCode 125. Valid Palindrome.
// https://leetcode.com/problems/valid-palindrome/
package validpalindrome

func isPalindrome(s string) bool {
	ftIdx, ndIdx := 0, len(s)-1

	for ftIdx < ndIdx {
		for ftIdx < ndIdx && !isAlnum(s[ftIdx]) {
			ftIdx++

		}

		for ftIdx < ndIdx && !isAlnum(s[ndIdx]) {
			ndIdx--
		}

		if toLower(s[ftIdx]) != toLower(s[ndIdx]) {
			return false
		}

		ftIdx++
		ndIdx--
	}

	return true
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32

	}
	return c
}

func isAlnum(c byte) bool {
	return c >= '0' && c <= '9' ||
		c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z'
}
