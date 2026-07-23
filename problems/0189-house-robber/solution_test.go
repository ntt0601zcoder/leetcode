package houserobber

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name   string
		houses []int
		want   int
	}{
		{
			name:   "one house - zero money",
			houses: []int{0},
			want:   0,
		},
		{
			name:   "one house - non-zero money",
			houses: []int{10},
			want:   10,
		},
		{
			name:   "multiple houses - zero money",
			houses: []int{0, 0, 0},
			want:   0,
		},
		{
			name:   "multiple houses - tc 1",
			houses: []int{1, 0, 3},
			want:   4,
		},
		{
			name:   "multiple houses - tc 2",
			houses: []int{1, 5, 3},
			want:   5,
		},
		{
			name:   "multiple houses - tc 3",
			houses: []int{6, 4, 0, 7},
			want:   13,
		},
		{
			name:   "multiple houses - tc 4",
			houses: []int{3, 0, 4, 3},
			want:   7,
		},
		{
			name:   "multiple houses - tc 5",
			houses: []int{1, 5, 3, 4},
			want:   9,
		},
		{
			name:   "multiple houses - tc 6",
			houses: []int{2, 1, 3, 2, 7},
			want:   12,
		},
	}

	for _, tc := range tests {
		got := rob(tc.houses)

		if got != tc.want {
			t.Errorf("TC failed (name: %s, houses: %x, want: %d)", tc.name, tc.houses, tc.want)
		}
	}
}
