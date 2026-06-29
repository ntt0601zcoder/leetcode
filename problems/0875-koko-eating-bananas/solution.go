// Package kokoeatingbananas solves LeetCode 875. Koko Eating Bananas.
// https://leetcode.com/problems/koko-eating-bananas/
package kokoeatingbananas

import "math"

func minEatingSpeed(piles []int, h int) int {
	minK, maxK := 1, piles[0]

	for i := 1; i < len(piles); i++ {
		if piles[i] > maxK {
			maxK = piles[i]
		}

		if piles[i] < minK {
			minK = piles[i]
		}
	}

	bestTime, bestK := 0, maxK

	for minK <= maxK {
		midK := minK + (maxK-minK)/2
		totalTime := 0

		for i := 0; i < len(piles); i++ {
			totalTime += int(math.Ceil(float64(piles[i]) / float64(midK)))
		}

		switch {
		case totalTime > h:
			minK = midK + 1
		case totalTime >= bestTime:
			bestTime = totalTime
			bestK = midK
			maxK = midK - 1
		}

	}

	return bestK
}
