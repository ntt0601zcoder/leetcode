package mincosttoconnectallpoints

import (
	"sort"
)

type Edge struct {
	u, v, w int
}

func minCostConnectPointsKruskal(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}

	edges := []Edge{}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			edge := Edge{
				u: i,
				v: j,
				w: calcDistance([2]int{points[i][0], points[i][1]}, [2]int{points[j][0], points[j][1]}),
			}
			edges = append(edges, edge)
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].w < edges[j].w
	})

	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}

		return x
	}

	rank := make([]int, n)
	union := func(a, b int) bool {
		ra, rb := find(a), find(b)

		if ra == rb {
			return false
		}

		if rank[ra] < rank[rb] {
			ra, rb = rb, ra
		}

		parent[rb] = ra

		if rank[ra] == rank[rb] {
			rank[ra]++
		}

		return true
	}

	total, count := 0, 0

	for _, edge := range edges {
		if union(edge.u, edge.v) {
			count++
			total += edge.w

			if count == n-1 {
				break
			}
		}
	}

	return total
}
