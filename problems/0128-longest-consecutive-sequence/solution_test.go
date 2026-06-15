package longestconsecutivesequence

import "testing"

// To compare multiple approaches, declare them in a map and range over it,
// e.g. (see problems/0001-two-sum for a full example):
//   var solutions = map[string]func(/* args */) /* ret */{
//       "v1": longestConsecutiveSequence,
//   }
//   for name, fn := range solutions { for _, tt := range tests { ... } }

func TestLongestConsecutiveSequence(t *testing.T) {
	tests := []struct {
		name string
		// TODO: input fields
		want int
	}{
		// TODO: add cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: call longestConsecutiveSequence and compare against tt.want
			_ = tt
		})
	}
}
