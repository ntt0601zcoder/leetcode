package longestrepeatingcharacterreplacement

import (
	"testing"
	"time"
)

// solutions lists every approach; the table test runs all cases against each,
// so adding a new approach is just one line here.
var solutions = map[string]func(s string, k int) int{
	"slidingwindow": characterReplacement,
}

// call runs fn with a timeout so a non-terminating solution (e.g. a sliding
// window whose inner shrink loop never advances `left`) reports a clean
// failure instead of hanging the whole test run. A panic is also caught.
func call(fn func(string, int) int, s string, k int) (got int, bad bool) {
	done := make(chan int, 1)
	go func() {
		defer func() { _ = recover() }()
		done <- fn(s, k)
	}()
	select {
	case got = <-done:
		return got, false
	case <-time.After(500 * time.Millisecond):
		return 0, true
	}
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{name: "leetcode example 1", s: "ABAB", k: 2, want: 4},
		{name: "leetcode example 2", s: "AABABBA", k: 1, want: 4},
		{name: "example 2 with k=2", s: "AABABBA", k: 2, want: 5},
		{name: "empty string", s: "", k: 0, want: 0},
		{name: "single char k=0", s: "A", k: 0, want: 1},
		{name: "single char k=1", s: "A", k: 1, want: 1},
		{name: "all same k=0", s: "AAAA", k: 0, want: 4},
		{name: "all same k=2", s: "AAAA", k: 2, want: 4},
		{name: "k=0 longest run", s: "AABBBCC", k: 0, want: 3},
		{name: "k=0 all distinct", s: "ABC", k: 0, want: 1},
		{name: "replace leading char", s: "ABBB", k: 1, want: 4},
		{name: "replace both ends", s: "BAAAB", k: 2, want: 5},
		{name: "k exceeds replacements needed", s: "ABCDE", k: 5, want: 5},
		{name: "k=1 not enough", s: "ABCDE", k: 1, want: 2},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, bad := call(fn, tc.s, tc.k)
				if bad {
					t.Fatalf("characterReplacement(%q, %d) panicked or did not return within 500ms, want %d", tc.s, tc.k, tc.want)
				}
				if got != tc.want {
					t.Errorf("characterReplacement(%q, %d) = %d, want %d", tc.s, tc.k, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkCharacterReplacement(b *testing.B) {
	// Repeating pattern with a modest k to exercise the shrink logic.
	s := ""
	for i := 0; i < 2000; i++ {
		s += "ABCDE"
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(s, 3)
			}
		})
	}
}
