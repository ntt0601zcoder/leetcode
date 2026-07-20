package minimumwindowsubstring

import (
	"testing"
	"time"
)

// solutions lists every approach; the table test runs all cases against each.
var solutions = map[string]func(s string, t string) string{
	"slidingwindow": minWindow,
}

// call runs fn with a timeout so a non-terminating solution (e.g. a shrink
// loop whose left pointer never advances) reports a clean failure instead of
// hanging the whole test run. A panic is also caught and surfaced as a failure.
func call(fn func(string, string) string, s, t string) (got string, timedOut bool) {
	done := make(chan string, 1)
	go func() {
		defer func() { _ = recover() }()
		done <- fn(s, t)
	}()
	select {
	case got = <-done:
		return got, false
	case <-time.After(500 * time.Millisecond):
		return "", true
	}
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		// LeetCode examples.
		{name: "leetcode example 1", s: "ADOBECODEBANC", t: "ABC", want: "BANC"},
		{name: "leetcode example 2 (single)", s: "a", t: "a", want: "a"},
		{name: "leetcode example 3 (t longer than s)", s: "a", t: "aa", want: ""},
		// Multiplicity: t requires two 'a's.
		{name: "duplicates in t exact", s: "aa", t: "aa", want: "aa"},
		{name: "duplicates in t within longer s", s: "baacd", t: "aa", want: "aa"},
		{name: "multiplicity picks wider window", s: "aaflslflsldkalskaaa", t: "aaa", want: "aaa"},
		// No valid window.
		{name: "char not present", s: "abc", t: "d", want: ""},
		{name: "empty s", s: "", t: "a", want: ""},
		// Whole string / positional windows.
		{name: "whole string reordered", s: "abc", t: "cba", want: "abc"},
		{name: "window at the end", s: "abcdxyz", t: "xyz", want: "xyz"},
		{name: "later shorter window wins", s: "cabwefgewcwaefgcf", t: "cae", want: "cwae"},
		// Larger case with a unique minimal window.
		{name: "single minimal ab window", s: "aaab", t: "ab", want: "ab"},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, timedOut := call(fn, tc.s, tc.t)
				if timedOut {
					t.Fatalf("minWindow(%q, %q) did not return within 500ms (likely infinite loop), want %q", tc.s, tc.t, tc.want)
				}
				if got != tc.want {
					t.Errorf("minWindow(%q, %q) = %q, want %q", tc.s, tc.t, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkMinWindow(b *testing.B) {
	// Long string where the only full match spans nearly the whole input, so the
	// window has to grow and shrink across the entire length.
	s := ""
	for i := 0; i < 5000; i++ {
		s += "xyzxyzxyz"
	}
	s = "a" + s + "bc"
	t := "abc"
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(s, t)
			}
		})
	}
}
