// Package dailytemperatures solves LeetCode 739. Daily Temperatures.
// https://leetcode.com/problems/daily-temperatures/
package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	result, stack := make([]int, len(temperatures)), make([]int, 0)

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[prev] = i - prev
		}

		stack = append(stack, i)
	}

	return result
}

func dailyTemperaturesBigOn(temperatures []int) []int {
	result, n := make([]int, len(temperatures)), len(temperatures)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				result[i] = j - i
				break
			}
		}
	}

	return result
}
