// Package trappingrainwater solves LeetCode 42. Trapping Rain Water.
// https://leetcode.com/problems/trapping-rain-water/
package trappingrainwater

func trapFullScan(height []int) int {
	n := len(height)
	total := 0

	for i := 0; i < n; i++ {
		leftMax, rightMax := 0, 0

		for j := 0; j <= i; j++ {
			leftMax = max(leftMax, height[j])
		}

		for k := i; k < n; k++ {
			rightMax = max(rightMax, height[k])
		}

		total += min(leftMax, rightMax) - height[i]
	}

	return total
}

func trapPrefixSuffix(height []int) int {
	n := len(height)

	if n == 0 {
		return 0
	}

	leftMaxs, rigtMaxs := make([]int, n), make([]int, n)

	leftMaxs[0] = height[0]
	for i := 1; i < n; i++ {
		leftMaxs[i] = max(leftMaxs[i-1], height[i])
	}

	rigtMaxs[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rigtMaxs[i] = max(rigtMaxs[i+1], height[i])
	}

	total := 0
	for i := 0; i < n; i++ {
		total += min(leftMaxs[i], rigtMaxs[i]) - height[i]
	}

	return total
}

func trapTwoPointers(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax := 0, 0
	total := 0

	for l < r {
		if height[l] < height[r] {
			if height[l] >= leftMax {
				leftMax = height[l]
			} else {
				total += leftMax - height[l]
			}
			l++
		} else {
			if height[r] >= rightMax {
				rightMax = height[r]
			} else {
				total += rightMax - height[r]
			}
			r--
		}
	}

	return total
}
