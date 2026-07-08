// Package longestrepeatingcharacterreplacement solves LeetCode 424. Longest Repeating Character Replacement.
// https://leetcode.com/problems/longest-repeating-character-replacement/
package longestrepeatingcharacterreplacement

func characterReplacement(s string, k int) int {
	var counter [26]int

	left, maxLen, result := 0, 0, 0

	for right := 0; right < len(s); right++ {
		counter[s[right]-'A']++
		maxLen = max(maxLen, counter[s[right]-'A'])

		for (right-left+1)-maxLen > k {
			counter[s[left]-'A']--
			left++
		}

		result = max(result, right-left+1)
	}

	return result
}
