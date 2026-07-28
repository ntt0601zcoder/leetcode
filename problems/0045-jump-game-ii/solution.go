// Package jumpgameii solves LeetCode 45. Jump Game Ii.
// https://leetcode.com/problems/jump-game-ii/
package jumpgameii

func jump(nums []int) int {
	jumps := 0
	farthest := 0
	farthestPos := 0

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])

		if i == farthestPos {
			farthestPos = farthest
			jumps++
		}
	}

	return jumps
}
