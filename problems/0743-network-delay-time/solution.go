package networkdelaytime

import "container/heap"

type MinHeap [][2]int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i int, j int) bool {
	return m[i][0] < m[j][0]
}

func (m *MinHeap) Pop() any {
	old := *m
	x := old[len(old)-1]
	*m = old[:len(old)-1]
	return x
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.([2]int))
}

func (m MinHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][][2]int, n+1)
	for _, time := range times {
		graph[time[0]] = append(graph[time[0]], [2]int{time[1], time[2]})
	}

	const INF = 1 << 30
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = INF
	}
	dist[k] = 0

	pq := &MinHeap{{0, k}}

	for pq.Len() > 0 {
		curr := pq.Pop().([2]int)
		d, node := curr[0], curr[1]

		if d > dist[node] {
			continue
		}

		for _, edge := range graph[node] {
			next, w := edge[0], edge[1]

			if d+w < dist[next] {
				dist[next] = d + w
				heap.Push(pq, [2]int{d + w, next})
			}
		}
	}

	ans := 0

	for i := 1; i <= n; i++ {
		if dist[i] == INF {
			return -1
		}

		ans = max(ans, dist[i])
	}

	return ans
}
