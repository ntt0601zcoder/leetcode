package asteroidcollision

import (
	"reflect"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(asteroids []int) []int{
	"stack": asteroidCollision,
}

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		name      string
		asteroids []int
		want      []int
	}{
		{name: "leetcode example 1", asteroids: []int{5, 10, -5}, want: []int{5, 10}},
		{name: "leetcode example 2", asteroids: []int{8, -8}, want: []int{}},
		{name: "leetcode example 3", asteroids: []int{10, 2, -5}, want: []int{10}},
		{name: "leetcode example 4", asteroids: []int{-2, -1, 1, 2}, want: []int{-2, -1, 1, 2}},
		{name: "no collisions all right", asteroids: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "no collisions all left", asteroids: []int{-1, -2, -3}, want: []int{-1, -2, -3}},
		{name: "full annihilation pairs", asteroids: []int{5, -5, 5, -5}, want: []int{}},
		{name: "left then right never collide", asteroids: []int{-5, 5}, want: []int{-5, 5}},
		{name: "big right survives", asteroids: []int{10, 2, -5}, want: []int{10}},
		{name: "right destroyed by bigger left-mover chain", asteroids: []int{2, -2}, want: []int{}},
		{name: "single", asteroids: []int{7}, want: []int{7}},
		{name: "empty", asteroids: []int{}, want: []int{}},
		{name: "cascade", asteroids: []int{-2, -2, 1, -2}, want: []int{-2, -2, -2}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := fn(append([]int(nil), tc.asteroids...))
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("asteroidCollision(%v) = %v, want %v", tc.asteroids, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkAsteroidCollision(b *testing.B) {
	asteroids := []int{5, 10, -5, 8, -8, 10, 2, -5, -2, -1, 1, 2, 3, -3, 4, -4}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(append([]int(nil), asteroids...))
			}
		})
	}
}
