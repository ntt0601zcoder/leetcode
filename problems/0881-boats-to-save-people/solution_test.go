package boatstosavepeople

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(people []int, limit int) int{
	"two-pointer": numRescueBoats,
}

func TestNumRescueBoats(t *testing.T) {
	tests := []struct {
		name   string
		people []int
		limit  int
		want   int
	}{
		{name: "leetcode example 1", people: []int{1, 2}, limit: 3, want: 1},
		{name: "leetcode example 2", people: []int{3, 2, 2, 1}, limit: 3, want: 3},
		{name: "leetcode example 3", people: []int{3, 5, 3, 4}, limit: 5, want: 4},
		{name: "single person", people: []int{2}, limit: 3, want: 1},
		// Everyone is so heavy that nobody can share: one boat each.
		{name: "everyone fits alone only", people: []int{5, 5, 5}, limit: 5, want: 3},
		// All can be paired up: half as many boats.
		{name: "all pairs fit", people: []int{1, 1, 2, 2}, limit: 3, want: 2},
		// One heavy person at the limit goes alone, the two lighter ones pair up.
		{name: "one heavy person", people: []int{5, 1, 2}, limit: 5, want: 2},
		// A pair summing exactly to the limit shares one boat.
		{name: "pair sums exactly to limit", people: []int{2, 3}, limit: 5, want: 1},
		// Two people who each weigh more than half the limit need separate boats.
		{name: "two heavy people separate boats", people: []int{4, 4}, limit: 5, want: 2},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				// Copy input since the solution sorts in place.
				p := append([]int(nil), tc.people...)
				got := fn(p, tc.limit)
				if got != tc.want {
					t.Errorf("numRescueBoats(%v, %d) = %d, want %d", tc.people, tc.limit, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkNumRescueBoats(b *testing.B) {
	n := 1000
	people := make([]int, n)
	for i := range people {
		people[i] = (i % 30000) + 1
	}
	limit := 30000
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				p := append([]int(nil), people...)
				fn(p, limit)
			}
		})
	}
}
