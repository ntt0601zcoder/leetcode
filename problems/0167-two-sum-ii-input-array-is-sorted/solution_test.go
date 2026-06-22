package twosumiiinputarrayissorted

import (
	"reflect"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(numbers []int, target int) []int{
	"twopointer": twoSum,
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int // 1-based indices
	}{
		{name: "leetcode example 1", numbers: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
		{name: "leetcode example 2", numbers: []int{2, 3, 4}, target: 6, want: []int{1, 3}},
		{name: "leetcode example 3", numbers: []int{-1, 0}, target: -1, want: []int{1, 2}},
		{name: "two elements", numbers: []int{1, 2}, target: 3, want: []int{1, 2}},
		{name: "answer at ends", numbers: []int{1, 2, 3, 4, 10}, target: 11, want: []int{1, 5}},
		{name: "answer in middle", numbers: []int{1, 4, 7, 11, 15}, target: 18, want: []int{3, 4}},
		{name: "negatives", numbers: []int{-5, -3, 0, 2, 8}, target: -1, want: []int{2, 4}},
		{name: "duplicates", numbers: []int{3, 3}, target: 6, want: []int{1, 2}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := fn(append([]int(nil), tc.numbers...), tc.target)
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("twoSum(%v, %d) = %v, want %v", tc.numbers, tc.target, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	numbers := make([]int, 1000)
	for i := range numbers {
		numbers[i] = i
	}
	target := 1997 // numbers[998] + numbers[999]
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(numbers, target)
			}
		})
	}
}
