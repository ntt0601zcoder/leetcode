// Package p4sum solves LeetCode 18. 4sum.
// https://leetcode.com/problems/4sum/
package p4sum

import (
	"slices"
)

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	result := make([][]int, 0)

	if n < 4 {
		return result
	}

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k, h := j+1, n-1

			for k < h {
				sum := nums[i] + nums[j] + nums[k] + nums[h]
				switch {
				case sum == target:
					result = append(result, []int{nums[i], nums[j], nums[k], nums[h]})
					k++
					h--

					for k < h && nums[k] == nums[k-1] {
						k++
					}
					for k < h && nums[h] == nums[h+1] {
						h--
					}
				case sum > target:
					h--
				case sum < target:
					k++
				}
			}
		}
	}

	return result
}
