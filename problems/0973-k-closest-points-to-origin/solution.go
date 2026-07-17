package kclosestpointstoorigin

import (
	"container/heap"
)

type DistanceHeap [][]int

func (d DistanceHeap) Len() int {
	return len(d)
}

func (d DistanceHeap) Less(i int, j int) bool {
	return calcDistance(d[i]) > calcDistance(d[j])
}

func (d *DistanceHeap) Pop() any {
	old := *d
	val := old[old.Len()-1]
	*d = old[:old.Len()-1]

	return val
}

func (d *DistanceHeap) Push(x any) {
	*d = append(*d, x.([]int))
}

func (d DistanceHeap) Swap(i int, j int) {
	d[i], d[j] = d[j], d[i]
}

func calcDistance(point []int) int { return point[0]*point[0] + point[1]*point[1] }

func kClosest(points [][]int, k int) [][]int {
	h := &DistanceHeap{}

	for _, point := range points {
		heap.Push(h, point)

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return *h
}
