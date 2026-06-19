// Package movezeroes solves LeetCode 283. Move Zeroes.
// https://leetcode.com/problems/move-zeroes/
package movezeroes

func moveZeroes(nums []int) {
	i, j, n := 0, 0, len(nums)

	if n < 2 {
		return
	}

	for i < len(nums) {
		switch {
		case j >= n:
			nums[i] = 0
			i++
		case nums[j] != 0:
			nums[i] = nums[j]
			i++
			j++
		default:
			j++
		}
	}
}
