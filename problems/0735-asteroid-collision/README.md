# 735. Asteroid Collision

- Difficulty: medium
- Link: https://leetcode.com/problems/asteroid-collision/

## Approach

Stack of surviving asteroids. A collision happens only when the top of the
stack moves right (positive) and the incoming asteroid moves left (negative).
While that holds, compare absolute sizes: equal sizes annihilate both, a
smaller top is popped and the incoming asteroid keeps checking the new top, and
a smaller incoming asteroid is destroyed. Anything that survives the chain of
collisions is pushed, and the stack is the final answer.

- Time:  O(n)
- Space: O(n)
