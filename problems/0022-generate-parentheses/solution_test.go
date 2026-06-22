package generateparentheses

import (
	"reflect"
	"sort"
	"testing"
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

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{name: "n=1", n: 1, want: []string{"()"}},
		{name: "n=2", n: 2, want: []string{"(())", "()()"}},
		{
			name: "leetcode example n=3",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "n=4",
			n:    4,
			want: []string{
				"(((())))", "((()()))", "((())())", "((()))()",
				"(()(()))", "(()()())", "(()())()", "(())(())",
				"(())()()", "()((()))", "()(()())", "()(())()",
				"()()(())", "()()()()",
			},
		},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := normalize(fn(tc.n))
				want := normalize(tc.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("generateParenthesis(%d) = %v, want %v (order-independent)", tc.n, got, want)
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
