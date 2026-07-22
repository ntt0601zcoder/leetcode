package mincosttoconnectallpoints

import "container/heap"

type MinHeap [][2]int

func minCostConnectPointsPrim(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return 0
	}

	inMST := make([]bool, n)
	total := 0
	count := 0
	pq := &MinHeap{{0, 0}}

	for pq.Len() > 0 && count < n {
		cur := heap.Pop(pq).([2]int)
		cost, u := cur[0], cur[1]

		if inMST[u] {
			continue
		}

		inMST[u] = true
		total += cost
		count++

		for i := 0; i < n; i++ {
			if inMST[i] {
				continue
			}

			heap.Push(pq, [2]int{calcDistance([2]int{points[u][0], points[u][1]}, [2]int{points[i][0], points[i][1]}), i})
		}
	}

	return total
}

func calcDistance(point1, point2 [2]int) int {
	return abs(point1[0]-point2[0]) + abs(point1[1]-point2[1])
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.([2]int)) }
func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
