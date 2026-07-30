// Package mergeintervals solves LeetCode 56. Merge Intervals.
// https://leetcode.com/problems/merge-intervals/
package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}
	lastItem := result[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= lastItem[1] {
			result[len(result)-1][1] = max(lastItem[1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
		}

		lastItem = result[len(result)-1]
	}

	return result
}
