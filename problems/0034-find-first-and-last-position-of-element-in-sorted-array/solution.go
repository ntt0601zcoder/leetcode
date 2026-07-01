// Package findfirstandlastpositionofelementinsortedarray solves LeetCode 34. Find First And Last Position Of Element In Sorted Array.
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	left := findBound(nums, target, true)

	if left == -1 {
		return []int{-1, -1}
	}

	right := findBound(nums, target, false)

	return []int{left, right}
}

func findBound(nums []int, target int, leftmost bool) int {
	left, right := 0, len(nums)-1

	pos := -1
	for left <= right {
		mid := left + (right-left)/2

		switch {
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		default:
			pos = mid

			if leftmost {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return pos
}

func searchRangeBruteForce(nums []int, target int) []int {
	left, right := -1, -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			left = i
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			right = i
			break
		}
	}

	return []int{left, right}
}
