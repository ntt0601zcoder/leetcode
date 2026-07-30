// Package partitionlabels solves LeetCode 763. Partition Labels.
// https://leetcode.com/problems/partition-labels/
package partitionlabels

func partitionLabels(s string) []int {
	lastIdxMap := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		lastIdxMap[s[i]] = i
	}

	start, end := 0, 0
	result := []int{}

	for i := 0; i < len(s); i++ {
		if lastIdxMap[s[i]] > end {
			end = lastIdxMap[s[i]]
		}

		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}
