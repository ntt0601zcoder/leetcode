package kclosestpointstoorigin

import "container/heap"

type MinHeap [][]int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i int, j int) bool {
	return calcDistance(m[i]) < calcDistance(m[j])
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	val := old[n-1]
	*m = old[:n-1]

	return val
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.([]int))
}

func (m MinHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func kClosest(points [][]int, k int) [][]int {
	h := &MinHeap{}

	for _, point := range points {
		heap.Push(h, point)
	}

	var result [][]int

	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).([]int))
	}

	return result
}

func calcDistance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
