// Package largestrectangleinhistogram solves LeetCode 84. Largest Rectangle In Histogram.
// https://leetcode.com/problems/largest-rectangle-in-histogram/
package largestrectangleinhistogram

type Node struct {
	index, height int
}

func largestRectangleArea(heights []int) int {
	maxArea, n, stack := 0, len(heights), make([]Node, 0)

	for i, height := range heights {
		start := i
		for len(stack) > 0 && height < stack[len(stack)-1].height {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			width := i - node.index
			area := width * node.height

			maxArea = max(maxArea, area)
			start = node.index
		}

		stack = append(stack, Node{index: start, height: height})
	}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		width := n - node.index
		area := width * node.height

		maxArea = max(maxArea, area)
	}

	return maxArea
}
