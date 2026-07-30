// Package nonoverlappingintervals solves LeetCode 435. Non Overlapping Intervals.
// https://leetcode.com/problems/non-overlapping-intervals/
package nonoverlappingintervals

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	picked := 1
	lastEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= lastEnd {
			picked++
			lastEnd = intervals[i][1]
		}
	}

	return len(intervals) - picked
}
