package topkfrequentelements

import (
	"reflect"
	"sort"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(nums []int, k int) []int{
	"freqsort": topKFrequent,
}

// sortedCopy returns a sorted copy so results compare order-independently
// (the problem allows the answer in any order).
func sortedCopy(a []int) []int {
	c := append([]int(nil), a...)
	sort.Ints(c)
	return c
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "basic", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{name: "single element", nums: []int{1}, k: 1, want: []int{1}},
		{name: "two distinct k=2", nums: []int{1, 2}, k: 2, want: []int{1, 2}},
		{name: "k=1 clear winner", nums: []int{4, 4, 4, 5, 5, 6}, k: 1, want: []int{4}},
		{name: "negatives", nums: []int{-1, -1, -2, -2, -2, 3}, k: 2, want: []int{-1, -2}},
		{name: "with zero", nums: []int{3, 0, 1, 0}, k: 1, want: []int{0}},
		{name: "all distinct k=all", nums: []int{5, 6, 7}, k: 3, want: []int{5, 6, 7}},
		{name: "all tied k=all", nums: []int{1, 1, 2, 2, 3, 3}, k: 3, want: []int{1, 2, 3}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := fn(tc.nums, tc.k)
				if len(got) != tc.k {
					t.Errorf("topKFrequent(%v, %d) returned %d elements, want %d", tc.nums, tc.k, len(got), tc.k)
				}
				if !reflect.DeepEqual(sortedCopy(got), sortedCopy(tc.want)) {
					t.Errorf("topKFrequent(%v, %d) = %v, want %v (any order)", tc.nums, tc.k, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkTopKFrequent(b *testing.B) {
	nums := make([]int, 0, 2000)
	for v := 0; v < 1000; v++ {
		nums = append(nums, v, v) // each value appears twice
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums, 100)
			}
		})
	}
}
