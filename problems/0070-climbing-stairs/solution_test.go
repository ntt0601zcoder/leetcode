package climbingstairs

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "zero stair", n: 0, want: 1,
		},
		{
			name: "one stair", n: 1, want: 1,
		},
		{
			name: "two stairs", n: 2, want: 2,
		},
		{
			name: "three stairs", n: 3, want: 3,
		},
	}

	for _, tc := range tests {
		got := climbStairs(tc.n)

		if got != tc.want {
			t.Errorf("Testcase failed (nam=%s, n=%d, want=%d)", tc.name, tc.n, tc.want)
		}
	}
}
