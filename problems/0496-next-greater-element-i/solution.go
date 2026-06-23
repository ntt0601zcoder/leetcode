// Package nextgreaterelementi solves LeetCode 496. Next Greater Element I.
// https://leetcode.com/problems/next-greater-element-i/
package nextgreaterelementi

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0, len(nums2))
	greaterMap := make(map[int]int)

	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			greaterMap[top] = num
		}
		stack = append(stack, num)
	}

	result := make([]int, 0, len(greaterMap))

	for _, num := range nums1 {
		if greaterVal, ok := greaterMap[num]; ok {
			result = append(result, greaterVal)
		} else {
			result = append(result, -1)
		}
	}

	return result
}
