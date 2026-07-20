package generateparentheses

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(n int) []string{
	"backtrack": generateParenthesis,
}

// normalize makes the result order-independent so the test does not depend on
// the order in which combinations are produced.
func normalize(combos []string) []string {
	out := append([]string(nil), combos...)
	sort.Strings(out)
	return out
}

// catalan returns the nth Catalan number, which is exactly the count of
// well-formed parentheses combinations for n pairs. Computed with the
// recurrence C(0)=1, C(k+1) = C(k) * 2*(2k+1) / (k+2); the numerator is always
// divisible by (k+2), so integer arithmetic stays exact.
func catalan(n int) int {
	c := 1
	for k := 0; k < n; k++ {
		c = c * 2 * (2*k + 1) / (k + 2)
	}
	return c
}

// isBalanced reports whether s consists only of '(' / ')' and is well-formed
// (never dips below zero, ends at zero).
func isBalanced(s string) bool {
	bal := 0
	for _, r := range s {
		switch r {
		case '(':
			bal++
		case ')':
			bal--
		default:
			return false
		}
		if bal < 0 {
			return false
		}
	}
	return bal == 0
}

// runGuarded runs fn(n) with a panic recover and a 500ms timeout so a buggy
// solution that crashes or recurses/loops forever fails cleanly instead of
// taking down (or OOM-ing) the whole test binary.
func runGuarded(t *testing.T, fn func(int) []string, n int) []string {
	t.Helper()
	ch := make(chan []string, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panic for n=%d: %v", n, r)
				ch <- nil
			}
		}()
		ch <- fn(n)
	}()
	select {
	case out := <-ch:
		return out
	case <-time.After(500 * time.Millisecond):
		t.Fatalf("timed out for n=%d (possible infinite recursion/loop)", n)
		return nil
	}
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string // exact expected set; nil => property checks only
	}{
		{name: "n=1 single pair", n: 1, want: []string{"()"}},
		{name: "n=2", n: 2, want: []string{"(())", "()()"}},
		{
			name: "leetcode example n=3",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "n=4 fourteen combos",
			n:    4,
			want: []string{
				"(((())))", "((()()))", "((())())", "((()))()",
				"(()(()))", "(()()())", "(()())()", "(())(())",
				"(())()()", "()((()))", "()(()())", "()(())()",
				"()()(())", "()()()()",
			},
		},
		{name: "n=5", n: 5},
		{name: "n=6", n: 6},
		{name: "n=7", n: 7},
		{name: "n=8 max constraint", n: 8},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := runGuarded(t, fn, tc.n)

				// Every combination must be well-formed and exactly 2n long,
				// with no duplicates.
				seen := make(map[string]bool, len(got))
				for _, s := range got {
					if len(s) != 2*tc.n {
						t.Errorf("combination %q has length %d, want %d", s, len(s), 2*tc.n)
					}
					if !isBalanced(s) {
						t.Errorf("combination %q is not a well-formed parentheses string", s)
					}
					if seen[s] {
						t.Errorf("duplicate combination %q", s)
					}
					seen[s] = true
				}

				// The count must equal the nth Catalan number (14 for n=4).
				if want := catalan(tc.n); len(got) != want {
					t.Errorf("generateParenthesis(%d) produced %d combos, want %d (Catalan number)", tc.n, len(got), want)
				}

				// Exact-set check for the small n where we enumerated the answer.
				if tc.want != nil {
					if g, w := normalize(got), normalize(tc.want); !reflect.DeepEqual(g, w) {
						t.Errorf("generateParenthesis(%d) = %v, want %v (order-independent)", tc.n, g, w)
					}
				}
			})
		}
	}
}

func BenchmarkGenerateParenthesis(b *testing.B) {
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(8)
			}
		})
	}
}
