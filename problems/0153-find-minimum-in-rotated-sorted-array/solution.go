// Package findminimuminrotatedsortedarray solves LeetCode 153. Find Minimum In Rotated Sorted Array.
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
package findminimuminrotatedsortedarray

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
