// Package longestsubstringwithoutrepeatingcharacters solves LeetCode 3. Longest Substring Without Repeating Characters.
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	left, maxlen := 0, 0
	set := make(map[byte]struct{}, 0)

	for right, c := range s {
		if _, ok := set[byte(c)]; ok {
			for {
				if _, ok := set[byte(c)]; !ok {
					break
				}
				delete(set, s[left])
				left++
			}
			set[byte(c)] = struct{}{}
		} else {
			set[byte(c)] = struct{}{}
			maxlen = max(maxlen, right-left+1)
		}
	}

	return maxlen
}
