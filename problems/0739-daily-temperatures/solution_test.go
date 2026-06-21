package dailytemperatures

import (
	"reflect"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(temperatures []int) []int{
	"stack": dailyTemperatures,
	"brute": dailyTemperaturesBigOn,
}

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "leetcode example 1", in: []int{73, 74, 75, 71, 69, 72, 76, 73}, want: []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{name: "strictly increasing", in: []int{30, 40, 50, 60}, want: []int{1, 1, 1, 0}},
		{name: "leetcode example 3", in: []int{30, 60, 90}, want: []int{1, 1, 0}},
		{name: "all equal", in: []int{50, 50, 50}, want: []int{0, 0, 0}},
		{name: "strictly decreasing", in: []int{80, 70, 60}, want: []int{0, 0, 0}},
		{name: "single day", in: []int{50}, want: []int{0}},
		{name: "two increasing", in: []int{50, 60}, want: []int{1, 0}},
		{name: "two decreasing", in: []int{60, 50}, want: []int{0, 0}},
		// Equal temps must NOT count — only strictly warmer; the warmer day is far.
		{name: "equal then warmer", in: []int{70, 70, 70, 71}, want: []int{3, 2, 1, 0}},
		{name: "dip then recover", in: []int{55, 38, 53, 81}, want: []int{3, 1, 1, 0}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				in := append([]int(nil), tc.in...)
				got := fn(in)
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("dailyTemperatures(%v) = %v, want %v", tc.in, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkDailyTemperatures(b *testing.B) {
	temps := make([]int, 1000)
	for i := range temps {
		temps[i] = 100 - (i % 100) // mix of rises and falls
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(temps)
			}
		})
	}
}
