// Package findpeakelement solves LeetCode 162. Find Peak Element.
// https://leetcode.com/problems/find-peak-element/
package findpeakelement

func findPeakElement(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 0
	}

	if nums[0] > nums[1] {
		return 0
	}

	if nums[n-1] > nums[n-2] {
		return n - 1
	}

	left, right := 1, n-2

	for left <= right {
		mid := left + (right-left)/2

		switch {
		case nums[mid] < nums[mid+1]:
			left = mid + 1
		case nums[mid] < nums[mid-1]:
			right = mid - 1
		case nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1]:
			return mid
		}
	}

	return -1
}
