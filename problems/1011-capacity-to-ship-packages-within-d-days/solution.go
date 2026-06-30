// Package capacitytoshippackageswithinddays solves LeetCode 1011. Capacity To Ship Packages Within D Days.
// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
package capacitytoshippackageswithinddays

func shipWithinDays(weights []int, days int) int {
	minW, maxW := 0, 0

	for _, w := range weights {
		maxW += w

		if w > minW {
			minW = w
		}
	}

	for minW <= maxW {
		midW := minW + (maxW-minW)/2
		daysTaken := calcDaysTaken(weights, midW)

		switch {
		case daysTaken > days:
			minW = midW + 1
		default:
			maxW = midW - 1
		}

	}

	return minW
}

func calcDaysTaken(weights []int, cap int) int {
	days, currW := 1, 0

	for _, w := range weights {
		if currW+w > cap {
			days++
			currW = w
		} else {
			currW += w
		}
	}

	return days
}
