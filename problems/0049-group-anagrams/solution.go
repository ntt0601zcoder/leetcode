// Package groupanagrams solves LeetCode 49. Group Anagrams.
// https://leetcode.com/problems/group-anagrams/
package groupanagrams

func groupAnagramsMap(strs []string) [][]string {
	resultM := make(map[string][]string)

	for _, str := range strs {
		exsits := false

		for key, values := range resultM {
			if isAnagramMap(str, key) {
				resultM[key] = append(values, str)
				exsits = true
			}
		}

		if !exsits {
			resultM[str] = []string{str}
		}
	}

	result := make([][]string, 0)

	for _, values := range resultM {
		result = append(result, values)
	}

	return result
}

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
