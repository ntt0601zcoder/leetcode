package nextgreaterelementi

import (
	"reflect"
	"testing"
)

// solutions lists every approach. The table test runs all cases against
// each one, so adding a new approach is just one line here.
var solutions = map[string]func(nums1 []int, nums2 []int) []int{
	"stack": nextGreaterElement,
}

// call runs fn on copies of the inputs (the solution is free to mutate) and
// converts any panic into a clean test failure instead of crashing the suite.
func call(t *testing.T, name string, fn func([]int, []int) []int, nums1, nums2 []int) (out []int) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("%s(%v, %v) panicked: %v", name, nums1, nums2, r)
		}
	}()
	a := append([]int(nil), nums1...)
	b := append([]int(nil), nums2...)
	return fn(a, b)
}

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{name: "example1", nums1: []int{4, 1, 2}, nums2: []int{1, 3, 4, 2}, want: []int{-1, 3, -1}},
		{name: "example2", nums1: []int{2, 4}, nums2: []int{1, 2, 3, 4}, want: []int{3, -1}},
		{name: "single both", nums1: []int{1}, nums2: []int{1}, want: []int{-1}},
		{name: "single has greater", nums1: []int{2}, nums2: []int{2, 3}, want: []int{3}},
		{name: "single no greater", nums1: []int{4}, nums2: []int{1, 3, 4, 2}, want: []int{-1}},
		{name: "descending all -1", nums1: []int{3, 2, 1}, nums2: []int{3, 2, 1}, want: []int{-1, -1, -1}},
		{name: "nums1 == nums2 sorted desc", nums1: []int{4, 3, 2, 1}, nums2: []int{4, 3, 2, 1}, want: []int{-1, -1, -1, -1}},
		{name: "ascending all greater", nums1: []int{1, 2, 3}, nums2: []int{1, 2, 3, 4}, want: []int{2, 3, 4}},
		{name: "shared max element", nums1: []int{2, 1, 3}, nums2: []int{4, 3, 2, 1, 5}, want: []int{5, 5, 5}},
		{name: "valley then peak", nums1: []int{6, 5, 4}, nums2: []int{6, 5, 4, 7}, want: []int{7, 7, 7}},
		{name: "mixed order subset", nums1: []int{2, 4, 1}, nums2: []int{1, 3, 4, 2, 5}, want: []int{5, 5, 3}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := call(t, name, fn, tc.nums1, tc.nums2)
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tc.nums1, tc.nums2, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkNextGreaterElement(b *testing.B) {
	nums2 := make([]int, 1000)
	for i := range nums2 {
		nums2[i] = 1000 - i // strictly descending: worst case for the stack
	}
	nums1 := append([]int(nil), nums2...)
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(append([]int(nil), nums1...), append([]int(nil), nums2...))
			}
		})
	}
}
