// Package longestsubstringwithoutrepeatingcharacters solves LeetCode 3. Longest Substring Without Repeating Characters.
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	left, maxlen := 0, 0
	set := make(map[byte]struct{})

	for right := 0; right < len(s); right++ {
		c := s[right]
		for {
			if _, ok := set[c]; !ok {
				break
			}
			delete(set, s[left])
			left++
		}
		set[c] = struct{}{}
		maxlen = max(maxlen, right-left+1)
	}
	return maxlen
}
