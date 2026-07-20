package taskscheduler

import "container/heap"

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i int, j int) bool {
	return m[i] > m[j]
}

func (m *MaxHeap) Pop() any {
	old := *m
	val := old[len(old)-1]
	*m = old[:len(old)-1]
	return val
}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m MaxHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func leastInterval(tasks []byte, n int) int {
	counter := [26]int{}
	for _, task := range tasks {
		counter[task-'A']++
	}

	h := &MaxHeap{}
	for _, count := range counter {
		if count > 0 {
			*h = append(*h, count)
		}
	}
	heap.Init(h)

	time := 0
	for h.Len() > 0 {
		leftover := []int{}

		for i := 0; i < n+1; i++ {
			if h.Len() > 0 {
				c := heap.Pop(h).(int)

				if c-1 > 0 {
					leftover = append(leftover, c-1)
				}
				time++
			} else if len(leftover) > 0 {
				time++
			}
		}

		for _, lv := range leftover {
			heap.Push(h, lv)
		}
	}

	return time
}

func leastIntervalMath(tasks []byte, n int) int {
	var count [26]int
	for _, t := range tasks {
		count[t-'A']++
	}

	maxCount := 0
	for _, c := range count {
		maxCount = max(maxCount, c)
	}

	numMax := 0
	for _, c := range count {
		if c == maxCount {
			numMax++
		}
	}

	frame := (maxCount-1)*(n+1) + numMax
	return max(frame, len(tasks))
}
