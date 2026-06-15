// Package containerwithmostwater solves LeetCode 11. Container With Most Water.
// https://leetcode.com/problems/container-with-most-water/
package containerwithmostwater

func maxArea(height []int) int {
	best := 0
	i, j := 0, len(height)-1

	for i < j {
		best = max(best, min(height[i], height[j])*(j-i))

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return best
}
