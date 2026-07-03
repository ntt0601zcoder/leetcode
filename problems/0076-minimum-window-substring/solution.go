// Package minimumwindowsubstring solves LeetCode 76. Minimum Window Substring.
// https://leetcode.com/problems/minimum-window-substring/
package minimumwindowsubstring

import "math"

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int, 0)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	needCount, left, resultStart, lenResult := len(need), 0, 0, math.MaxInt
	window, have := make(map[byte]int, 0), 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++

		if count, ok := need[c]; ok && window[c] == count {
			have++
		}

		for have == needCount {
			if right-left+1 < lenResult {
				lenResult = right - left + 1
				resultStart = left
			}

			lc := s[left]
			window[lc]--

			if count, ok := need[lc]; ok && window[lc] < count {
				have--
			}

			left++
		}
	}

	if lenResult == math.MaxInt {
		return ""
	}

	return s[resultStart : resultStart+lenResult]
}
