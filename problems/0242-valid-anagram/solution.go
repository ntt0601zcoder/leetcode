// Package validanagram solves LeetCode 242. Valid Anagram.
// https://leetcode.com/problems/valid-anagram/
package validanagram

func isAnagramMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cc := make(map[rune]int)

	for _, c := range s {
		cc[c]++
	}

	for _, c := range t {
		if _, ok := cc[c]; !ok {
			return false
		}

		cc[c]--

		if cc[c] == 0 {
			delete(cc, c)
		}
	}

	return len(cc) == 0
}
