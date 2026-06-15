// Package p3sum solves LeetCode 15. 3sum.
// https://leetcode.com/problems/3sum/
package p3sum

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var result [][]int

	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, n-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			switch {
			case sum == 0:
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--

				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			case sum > 0:
				k--
			case sum < 0:
				j++
			}
		}

	}

	return result
}
