// Package cheapestflightswithinkstops solves LeetCode 787. Cheapest Flights Within K Stops.
// https://leetcode.com/problems/cheapest-flights-within-k-stops/
package cheapestflightswithinkstops

func findCheapestPrice(n int, flights [][]int, src, dst, k int) int {
	const INF = 1 << 30
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[src] = 0

	for i := 0; i <= k; i++ {
		tmp := make([]int, n)
		copy(tmp, dist)

		for _, f := range flights {
			from, to, cost := f[0], f[1], f[2]

			if dist[from] != INF && dist[from]+cost < tmp[to] {
				tmp[to] = dist[from] + cost
			}
		}

		dist = tmp
	}

	if dist[dst] == INF {
		return -1
	}

	return dist[dst]
}
