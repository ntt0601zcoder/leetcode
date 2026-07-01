package besttimetobuyandsellstock

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(prices []int) int{
	"onepass": maxProfit,
	"brute":   maxProfitGreedy,
}

// call guards against a panic (e.g. an out-of-range index) so the suite
// reports a clean failure instead of crashing the whole test binary.
func call(fn func([]int) int, prices []int) (got int, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	return fn(prices), false
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "leetcode example 1", prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{name: "leetcode example 2 (no profit)", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "increasing", prices: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "single day", prices: []int{5}, want: 0},
		{name: "two increasing", prices: []int{1, 5}, want: 4},
		{name: "two decreasing", prices: []int{5, 1}, want: 0},
		{name: "all equal", prices: []int{3, 3, 3}, want: 0},
		{name: "best sell at end", prices: []int{2, 4, 1, 7}, want: 6},
		{name: "min then later max", prices: []int{3, 2, 6, 1, 4}, want: 4},
		{name: "strictly decreasing", prices: []int{5, 4, 3, 2, 1}, want: 0},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, panicked := call(fn, append([]int(nil), tc.prices...))
				if panicked {
					t.Fatalf("maxProfit(%v) panicked, want %d", tc.prices, tc.want)
				}
				if got != tc.want {
					t.Errorf("maxProfit(%v) = %d, want %d", tc.prices, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	prices := make([]int, 1000)
	for i := range prices {
		prices[i] = (i*7 + 13) % 100
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(prices)
			}
		})
	}
}
