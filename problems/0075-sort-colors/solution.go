// Package sortcolors solves LeetCode 75. Sort Colors.
// https://leetcode.com/problems/sort-colors/
package sortcolors

import "slices"

func sortColors(nums []int) {
	left, mid, right := 0, 0, len(nums)-1

	for mid <= right {
		switch {
		case nums[mid] == 0:
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		case nums[mid] == 1:
			mid++
		case nums[mid] == 2:
			nums[right], nums[mid] = nums[mid], nums[right]
			right--
		}
	}
}

func sortColorsW1(nums []int) {
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	R, W, B := counter[0], counter[1], counter[2]

	copy(nums[:R], slices.Repeat([]int{0}, R))
	copy(nums[R:R+W], slices.Repeat([]int{1}, W))
	copy(nums[R+W:], slices.Repeat([]int{2}, B))
}
