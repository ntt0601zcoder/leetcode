// Package jumpgame solves LeetCode 55. Jump Game.
// https://leetcode.com/problems/jump-game/
package jumpgame

func canJump(nums []int) bool {
	maxReach := 0

	for i, num := range nums {
		if i > maxReach {
			return false
		}

		maxReach = max(maxReach, i+num)
	}

	return true
}
