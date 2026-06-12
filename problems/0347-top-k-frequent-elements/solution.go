// Package topkfrequentelements solves LeetCode 347. Top K Frequent Elements.
// https://leetcode.com/problems/top-k-frequent-elements/
package topkfrequentelements

import "sort"

func topKFrequent(nums []int, k int) []int {
	result := make([]int, 0)

	if len(nums) == 0 || k <= 0 {
		return result
	}

	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	for num := range counter {
		result = append(result, num)
	}

	sort.Slice(result, func(i, j int) bool {
		return counter[result[i]] > counter[result[j]]
	})

	return result[:k]
}
