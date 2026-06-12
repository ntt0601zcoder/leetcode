package twosum

import (
	"reflect"
	"testing"
)

// solutions lists every approach. The table test runs all cases against
// each one, so adding a new approach is just one line here.
var solutions = map[string]func(nums []int, target int) []int{
	"brute":   twoSumBrute,
	"hashmap": twoSumHashMap,
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "basic", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "middle", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "duplicate", nums: []int{3, 3}, target: 6, want: []int{0, 1}},
		{name: "no answer", nums: []int{1, 2, 3}, target: 7, want: nil},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.nums, tt.target)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.nums, tt.target, got, tt.want)
				}
			})
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	target := 1997 // 998 + 999
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(nums, target)
			}
		})
	}
}
