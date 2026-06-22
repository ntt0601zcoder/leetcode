package validparentheses

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(s string) bool{
	"stack": isValid,
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "leetcode example 1", s: "()", want: true},
		{name: "leetcode example 2", s: "()[]{}", want: true},
		{name: "leetcode example 3", s: "(]", want: false},
		{name: "nested valid", s: "([{}])", want: true},
		{name: "wrong order close", s: "([)]", want: false},
		{name: "single open", s: "(", want: false},
		{name: "single close", s: ")", want: false},
		{name: "empty string", s: "", want: true},
		{name: "unmatched trailing open", s: "(()", want: false},
		{name: "extra closing", s: "())", want: false},
		{name: "mismatched pair", s: "(}", want: false},
		{name: "deeply nested", s: "{[()()]}", want: true},
		{name: "close before open", s: "}{", want: false},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := fn(tc.s)
				if got != tc.want {
					t.Errorf("isValid(%q) = %v, want %v", tc.s, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	s := ""
	for i := 0; i < 500; i++ {
		s += "([{"
	}
	for i := 0; i < 500; i++ {
		s += "}])"
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(s)
			}
		})
	}
}
