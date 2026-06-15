// Package majorityelement solves LeetCode 169. Majority Element.
// https://leetcode.com/problems/majority-element/
package majorityelement

func majorityElement(nums []int) int {
	n := len(nums)
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++

		if counter[num] > n/2 {
			return num
		}
	}

	return 0
}

func majorityElementOptimal(nums []int) int {
	candidate, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
