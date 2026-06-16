// Package twosumiiinputarrayissorted solves LeetCode 167. Two Sum Ii Input Array Is Sorted.
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	result := make([]int, 2)

	for i < j {
		sum := numbers[i] + numbers[j]

		switch {
		case sum > target:
			j--
		case sum < target:
			i++
		case sum == target:
			result[0] = i + 1
			result[1] = j + 1
			return result
		}
	}
	return result
}
