// Package twosum solves LeetCode 1. Two Sum.
// https://leetcode.com/problems/two-sum/
package twosum

// twoSumBrute checks every pair. O(n^2) time, O(1) space.
func twoSumBrute(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// twoSumHashMap is one pass with a value->index map. O(n) time, O(n) space.
func twoSumHashMap(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
