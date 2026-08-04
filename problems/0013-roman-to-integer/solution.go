// Package romantointeger solves LeetCode 13. Roman To Integer.
// https://leetcode.com/problems/roman-to-integer/
package romantointeger

var RomanToIntMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		cur, _ := RomanToIntMap[s[i]]
		nextVal := 0

		if i+1 < len(s) {
			nextVal = RomanToIntMap[s[i+1]]
		}

		if cur < nextVal {
			result -= cur
		} else {
			result += cur
		}
	}

	return result
}
