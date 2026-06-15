// Package longestconsecutivesequence solves LeetCode 128. Longest Consecutive Sequence.
// https://leetcode.com/problems/longest-consecutive-sequence/
package longestconsecutivesequence

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		set[n] = struct{}{}
	}

	longest := 0

	for num := range set {
		if _, hasPrev := set[num-1]; hasPrev {
			continue
		}

		length := 1

		for {
			if _, ok := set[num+length]; !ok {
				break
			}

			length++

		}

		if length > longest {
			longest = length
		}
	}

	return longest
}
