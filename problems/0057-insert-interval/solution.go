// Package insertinterval solves LeetCode 57. Insert Interval.
// https://leetcode.com/problems/insert-interval/
package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	n := len(intervals)
	i := 0

	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	result = append(result, newInterval)

	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}
