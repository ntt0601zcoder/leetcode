// Package removekdigits solves LeetCode 402. Remove K Digits.
// https://leetcode.com/problems/remove-k-digits/
package removekdigits

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, len(num))
	counter := 0

	for i := 0; i < len(num); i++ {
		c := num[i]
		for len(stack) > 0 && stack[len(stack)-1] > c && counter < k {
			stack = stack[:len(stack)-1]
			counter++
		}

		stack = append(stack, c)
	}

	stack = stack[:len(stack)-(k-counter)]

	i := 0
	for i < len(stack) && stack[i] == '0' {
		i++
	}
	stack = stack[i:]

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
