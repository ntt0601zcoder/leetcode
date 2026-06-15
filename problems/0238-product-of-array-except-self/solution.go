// Package productofarrayexceptself solves LeetCode 238. Product Of Array Except Self.
// https://leetcode.com/problems/product-of-array-except-self/
package productofarrayexceptself

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}
