package longestsubstringwithoutrepeatingcharacters

import (
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(s string) int{
	"slidingwindow": lengthOfLongestSubstring,
}

// call runs fn with a timeout and a panic guard, so a non-terminating loop or
// an out-of-range index reports a clean failure instead of hanging/crashing.
func call(fn func(string) int, s string) (got int, outcome string) {
	done := make(chan int, 1)
	panicked := make(chan struct{}, 1)
	go func() {
		defer func() {
			if recover() != nil {
				panicked <- struct{}{}
			}
		}()
		done <- fn(s)
	}()
	select {
	case got = <-done:
		return got, "ok"
	case <-panicked:
		return 0, "panic"
	case <-time.After(time.Second):
		return 0, "timeout"
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "leetcode example 1", s: "abcabcbb", want: 3},
		{name: "leetcode example 2 (all same)", s: "bbbbb", want: 1},
		{name: "leetcode example 3", s: "pwwkew", want: 3},
		{name: "empty", s: "", want: 0},
		{name: "single char", s: "a", want: 1},
		{name: "all distinct", s: "abcdef", want: 6},
		{name: "repeat forces shrink", s: "dvdf", want: 3},
		{name: "mirror", s: "abba", want: 2},
		{name: "single space", s: " ", want: 1},
		{name: "two distinct", s: "au", want: 2},
		{name: "longest is a suffix", s: "tmmzuxt", want: 5},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, outcome := call(fn, tc.s)
				switch outcome {
				case "panic":
					t.Fatalf("lengthOfLongestSubstring(%q) panicked, want %d", tc.s, tc.want)
				case "timeout":
					t.Fatalf("lengthOfLongestSubstring(%q) did not return within 1s (likely infinite loop), want %d", tc.s, tc.want)
				}
				if got != tc.want {
					t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tc.s, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	const s = "abcdefghijklmnopqrstuvwxyzabcdefghij"
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(s)
			}
		})
	}
}
