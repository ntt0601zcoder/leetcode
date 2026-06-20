// Package boatstosavepeople solves LeetCode 881. Boats To Save People.
// https://leetcode.com/problems/boats-to-save-people/
package boatstosavepeople

import (
	"slices"
)

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	i, j, result := 0, len(people)-1, 0

	for i <= j {
		switch {
		case people[i]+people[j] <= limit, i == j:
			i++
			j--
			result++
		case people[i]+people[j] > limit:
			result++
			j--
		}
	}

	return result
}
