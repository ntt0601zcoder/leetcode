// Package pacificatlanticwaterflow solves LeetCode 417. Pacific Atlantic Water Flow.
// https://leetcode.com/problems/pacific-atlantic-water-flow/
package pacificatlanticwaterflow

/* THIS IS A STUPID SOLUTION
Contraints:
	- Height: [0, 10^5]
	-  1 <= rows, cols <= 200

Example: <skip>

Brute: Use BFS to travel all neighbors node. If have a node with lower, continue visit its neighbors. If this node can flow to both of seas, store it.

Pattern:
	- Uses BFS because input is a graph.
	- Uses Map to store whether this node can flow to a sea (-1: Pacific Ocean, 0: Atlantic Ocean, 1: Both). -> reduce calculation.

	Target:
	- Time: O(MxN)
	- Space: O(MxN + MxN) (M: numbr of rows, N: Number of columns) (fist MxN for Map that contains calculdated results, second on is for return value)
*/

/*
Constraints:
	- Height: [0, 10^5]
	- 1 <= rows, cols <= 200  → grid never empty (rows,cols >= 1)
	- Water flows from a cell to an adjacent cell with height <= current (EQUAL also flows)
	- Pacific  = top edge (row 0)      + left edge (col 0)
	- Atlantic = bottom edge (last row) + right edge (last col)
	- An EDGE cell always reaches its ocean regardless of its height (it's already at the border)

Expectation:
	- Return every cell whose water can reach BOTH oceans.

Example: <skip>

Brute: From EACH cell, BFS/DFS downhill (to <= neighbors) and check if it reaches both oceans.
	Cost: M*N cells, each doing an O(M*N) traversal → O((M*N)^2) → too slow for 200x200.

Pattern: REVERSE the flow — instead of "from cell X, can water reach the ocean?",
         ask "starting AT the ocean, going UPHILL, which cells can I reach?".
	- BFS/DFS from all Pacific-edge cells, climbing to neighbors with height >= current
	  → mark reachPacific.
	- Same from all Atlantic-edge cells → mark reachAtlantic.
	- Answer = cells where reachPacific && reachAtlantic (intersection).
	- Edge cells are seeded UNCONDITIONALLY (height doesn't matter); the >= condition
	  only applies when EXPANDING inward.
	- This turns M*N traversals into just 2 → brings cost down to O(M*N).
	- (Multi-source BFS/DFS: start from a whole line of border cells at once.)

Target:
	Time:  O(M*N) → each ocean's traversal visits each cell <= 1 time; 2 oceans = O(M*N).
	Space: O(M*N) → two visited matrices (reachPacific, reachAtlantic).
	                Recursion/queue depth <= M*N. Output NOT counted.
	(M = rows, N = cols)
*/

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	rows, cols := len(heights), len(heights[0])
	reachP, reachA := make([][]bool, rows), make([][]bool, rows)

	for i := range reachP {
		reachP[i] = make([]bool, cols)
		reachA[i] = make([]bool, cols)
	}

	var dfs func(r, c int, reach [][]bool)

	dfs = func(r, c int, reach [][]bool) {
		reach[r][c] = true

		for _, d := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := r+d[0], c+d[1]

			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}

			if reach[nr][nc] || heights[nr][nc] < heights[r][c] {
				continue
			}

			dfs(nr, nc, reach)
		}
	}

	for r := 0; r < rows; r++ {
		dfs(r, 0, reachP)
		dfs(r, cols-1, reachA)
	}
	for c := 0; c < cols; c++ {
		dfs(0, c, reachP)
		dfs(rows-1, c, reachA)
	}

	res := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if reachP[r][c] && reachA[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}
