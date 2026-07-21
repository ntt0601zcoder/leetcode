// Package courseschedule solves LeetCode 207. Course Schedule.
// https://leetcode.com/problems/course-schedule/
package courseschedule

/*
Constaints:
	- 1 <= numCourses <= 2000
	- 0 <= len(prerequisites) <= 5000
	- prerequisites is unable to null
	- len(prerequisites[i]) always equals 2
	- prerequisite pairs are unique
	- courses in the same pairs maybe the same

Expectation:
	- Return true if you can finish all courses. Otherwise, return false. <=> Graph has no cycle.

Example: <skip>

Brute:
	- Migrate prerequisites to directed graph.
	- Find cycle. If had, cannot finish, else OK.

Pattern:
	- Use directed graph because prerequisites is 2D array that illustrate graph.
	- Find cycle. If had, cannot finish, else OK.
	- Use DFS to travel until the end of this road. If face a visiting node -> Has cycle -> Unable to finish

Target:
	- Time: O(N)
	- Space: O(N)
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		inDegree[p[0]]++
	}

	queue := []int{}

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	finished := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		finished++

		for _, next := range graph[node] {
			inDegree[next]--

			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return finished == numCourses
}
