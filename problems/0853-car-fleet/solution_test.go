package carfleet

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(target int, position []int, speed []int) int{
	"sort": carFleet,
}

// call runs fn but turns a panic into a clean failure flag, so a buggy
// solution reports per-case instead of crashing the whole test binary.
func call(fn func(int, []int, []int) int, target int, position, speed []int) (got int, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	// Copy inputs in case a solution sorts in place.
	p := append([]int(nil), position...)
	s := append([]int(nil), speed...)
	return fn(target, p, s), false
}

func TestCarFleet(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		want     int
	}{
		{name: "leetcode example 1", target: 12, position: []int{10, 8, 0, 5, 3}, speed: []int{2, 4, 1, 1, 3}, want: 3},
		{name: "leetcode example 2", target: 10, position: []int{3}, speed: []int{3}, want: 1},
		{name: "leetcode example 3", target: 100, position: []int{0, 2, 4}, speed: []int{4, 2, 1}, want: 1},
		{name: "single car", target: 10, position: []int{5}, speed: []int{2}, want: 1},
		{name: "no cars", target: 10, position: []int{}, speed: []int{}, want: 0},
		// All same speed and never collide: each stays its own fleet.
		{name: "all same speed never merge", target: 10, position: []int{0, 2, 4}, speed: []int{1, 1, 1}, want: 3},
		// Two cars that merge into one fleet: rear car catches the front one.
		{name: "two cars merge into one fleet", target: 10, position: []int{0, 5}, speed: []int{2, 1}, want: 1},
		// Two cars that never merge: front car is faster.
		{name: "two cars stay separate", target: 10, position: []int{0, 5}, speed: []int{1, 2}, want: 2},
		// Already ordered by position ascending; back car merges into front.
		{name: "already ordered merge", target: 10, position: []int{2, 4}, speed: []int{3, 1}, want: 1},
		// All cars merge into a single fleet at the target: the frontmost
		// (pos 8) is slowest, so every car behind it catches up by the target.
		{name: "all merge into one", target: 10, position: []int{8, 6, 4, 2, 0}, speed: []int{1, 2, 3, 5, 7}, want: 1},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, panicked := call(fn, tc.target, tc.position, tc.speed)
				if panicked {
					t.Fatalf("carFleet(%d, %v, %v) panicked, want %d", tc.target, tc.position, tc.speed, tc.want)
				}
				if got != tc.want {
					t.Errorf("carFleet(%d, %v, %v) = %d, want %d", tc.target, tc.position, tc.speed, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkCarFleet(b *testing.B) {
	target := 1000
	n := 500
	position := make([]int, n)
	speed := make([]int, n)
	for i := 0; i < n; i++ {
		position[i] = i
		speed[i] = (i % 5) + 1
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				p := append([]int(nil), position...)
				s := append([]int(nil), speed...)
				fn(target, p, s)
			}
		})
	}
}
