// Package containsduplicate solves LeetCode 217. Contains Duplicate.
// https://leetcode.com/problems/contains-duplicate/
package containsduplicate

func containsDuplicate(nums []int) bool {
	m := make(map[int]interface{})

	for _, n := range nums {
		if _, exist := m[n]; exist {
			return true
		}

		m[n] = nil
	}

	return false
}
