// Package evaluatereversepolishnotation solves LeetCode 150. Evaluate Reverse Polish Notation.
// https://leetcode.com/problems/evaluate-reverse-polish-notation/
package evaluatereversepolishnotation

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]string, 0, len(tokens))

	for _, token := range tokens {
		number, err := strconv.Atoi(token)

		if err == nil && isValidNumber(number) {
			stack = append(stack, token)
		} else if isValidOperator(token) {
			if len(stack) < 2 {
				break
			}

			num1, _ := strconv.Atoi(stack[len(stack)-2])
			num2, _ := strconv.Atoi(stack[len(stack)-1])

			stack = stack[:len(stack)-2]
			result := calc(num1, num2, token)
			stack = append(stack, strconv.Itoa(result))
		}
	}

	r, _ := strconv.Atoi(stack[0])

	return r
}

func isValidOperator(s string) bool {
	switch s {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}

func isValidNumber(n int) bool {
	return n >= -200 && n <= 200
}

func calc(num1, num2 int, op string) int {
	switch op {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	default:
		return num1 / num2
	}
}
