package kthlargestelementinanarray

import "container/heap"

type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i int, j int) bool {
	return m[i] < m[j]
}

func (m *MinHeap) Pop() any {
	old := *m
	n := len(old)
	val := old[n-1]
	*m = old[:n-1]

	return val
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m MinHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}

	for _, num := range nums {
		heap.Push(h, num)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}
