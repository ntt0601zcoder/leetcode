// Package medianoftwosortedarrays solves LeetCode 4. Median Of Two Sorted Arrays.
// https://leetcode.com/problems/median-of-two-sorted-arrays/
package medianoftwosortedarrays

import (
	"math"
	"slices"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	short, long := nums1, nums2

	if len(nums1) > len(nums2) {
		short, long = long, short
	}

	shortLen, longLen := len(short), len(long)
	totalLen := shortLen + longLen
	leftHalfLen := (totalLen + 1) / 2
	lo, hi := 0, shortLen

	for lo <= hi {
		shortCut := (hi + lo) / 2
		longCut := leftHalfLen - shortCut
		shortLeft, shortRight := math.MinInt, math.MaxInt

		if shortCut > 0 {
			shortLeft = short[shortCut-1]
		}

		if shortCut < shortLen {
			shortRight = short[shortCut]
		}

		longLeft, longRight := math.MinInt, math.MaxInt

		if longCut > 0 {
			longLeft = long[longCut-1]
		}

		if longCut < longLen {
			longRight = long[longCut]
		}

		switch {
		case shortLeft > longRight:
			hi = shortCut - 1
		case longLeft > shortRight:
			lo = shortCut + 1
		default:
			maxLeft := max(shortLeft, longLeft)
			if totalLen%2 == 1 {
				return float64(maxLeft)
			}
			minRight := min(shortRight, longRight)
			return float64(maxLeft+minRight) / 2
		}
	}

	return 0
}

func findMedianSortedArraysBinarySearch(nums1 []int, nums2 []int) float64 {
	short, long := nums1, nums2
	if len(short) > len(long) {
		short, long = long, short
	}

	shortLen, longLen := len(short), len(long)
	totalLen := shortLen + longLen
	leftHalfLen := (totalLen + 1) / 2

	lo, hi := 0, shortLen
	for lo <= hi {
		shortCut := (lo + hi) / 2
		longCut := leftHalfLen - shortCut

		shortLeft, shortRight := math.MinInt, math.MaxInt
		if shortCut > 0 {
			shortLeft = short[shortCut-1]
		}
		if shortCut < shortLen {
			shortRight = short[shortCut]
		}

		longLeft, longRight := math.MinInt, math.MaxInt
		if longCut > 0 {
			longLeft = long[longCut-1]
		}
		if longCut < longLen {
			longRight = long[longCut]
		}

		switch {
		case shortLeft > longRight:
			hi = shortCut - 1
		case longLeft > shortRight:
			lo = shortCut + 1
		default:
			maxLeft := max(shortLeft, longLeft)
			if totalLen%2 == 1 {
				return float64(maxLeft)
			}
			minRight := min(shortRight, longRight)
			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}

func findMedianSortedArraysSort(nums1 []int, nums2 []int) float64 {
	nums := slices.Concat(nums1, nums2)
	slices.Sort(nums)

	n := len(nums)

	if n%2 == 0 {
		return (float64(nums[n/2]) + float64(nums[n/2-1])) / 2
	} else {
		return float64(nums[n/2])
	}
}
