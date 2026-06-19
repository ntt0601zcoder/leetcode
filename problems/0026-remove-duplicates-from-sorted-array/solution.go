// Package removeduplicatesfromsortedarray solves LeetCode 26. Remove Duplicates From Sorted Array.
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	i, j, n := 1, 1, len(nums)

	if n < 2 {
		if n == 0 {
			return 0
		}
		return 1
	}

	for j < n {
		switch {
		case nums[i-1] >= nums[j]:
			j++
		default:
			nums[i] = nums[j]
			i++
			j++
		}
	}

	return i
}
