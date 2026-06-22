// Package asteroidcollision solves LeetCode 735. Asteroid Collision.
// https://leetcode.com/problems/asteroid-collision/
package asteroidcollision

import "math"

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for _, asteroid := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, asteroid)
			continue
		}

		isDetroyed := false

		for len(stack) > 0 && hasCollision(stack[len(stack)-1], asteroid) {
			a1Weight, a2Weight := math.Abs(float64(asteroid)), math.Abs(float64(stack[len(stack)-1]))

			if a1Weight == a2Weight {
				stack = stack[:len(stack)-1]
				isDetroyed = true
				break
			}

			if a1Weight > a2Weight {
				stack = stack[:len(stack)-1]
				continue
			}

			isDetroyed = true
			break
		}

		if isDetroyed {
			continue
		}

		stack = append(stack, asteroid)
	}

	return stack
}

func hasCollision(a1, a2 int) bool {
	return a1 > 0 && a2 < 0
}
