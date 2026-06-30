package capacitytoshippackageswithinddays

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(weights []int, days int) int{
	"binarysearch": shipWithinDays,
}

func TestShipWithinDays(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		days    int
		want    int
	}{
		{name: "leetcode example 1", weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, days: 5, want: 15},
		{name: "leetcode example 2", weights: []int{3, 2, 2, 4, 1, 4}, days: 3, want: 6},
		{name: "leetcode example 3", weights: []int{1, 2, 3, 1, 1}, days: 4, want: 3},
		// One day => capacity must hold everything (the sum).
		{name: "one day (sum)", weights: []int{1, 2, 3, 4, 5}, days: 1, want: 15},
		{name: "one day big", weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, days: 1, want: 55},
		// days == len => each package alone => capacity is the max weight.
		{name: "days equals count (max weight)", weights: []int{1, 2, 3, 4, 5}, days: 5, want: 5},
		// days > len => still bounded below by the max weight.
		{name: "days exceed count (max weight)", weights: []int{3, 2, 2, 4, 1, 4}, days: 6, want: 4},
		{name: "single package one day", weights: []int{10}, days: 1, want: 10},
		// Many days but one package: capacity still bounded by the max weight.
		{name: "single package many days", weights: []int{2}, days: 5, want: 2},
		{name: "all equal split evenly", weights: []int{5, 5, 5, 5}, days: 2, want: 10},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := fn(append([]int(nil), tc.weights...), tc.days)
				if got != tc.want {
					t.Errorf("shipWithinDays(%v, %d) = %d, want %d", tc.weights, tc.days, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkShipWithinDays(b *testing.B) {
	weights := make([]int, 1000)
	for i := range weights {
		weights[i] = (i % 500) + 1
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(weights, 100)
			}
		})
	}
}
