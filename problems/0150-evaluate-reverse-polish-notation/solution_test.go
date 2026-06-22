package evaluatereversepolishnotation

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(tokens []string) int{
	"stack": evalRPN,
}

// call runs fn but turns a panic into a clean failure flag, so a buggy
// solution (e.g. divide-by-zero on a broken stack) reports per-case
// instead of crashing the whole test binary.
func call(fn func([]string) int, in []string) (got int, panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	return fn(in), false
}

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{name: "leetcode example 1", tokens: []string{"2", "1", "+", "3", "*"}, want: 9},
		// {name: "leetcode example 2", tokens: []string{"4", "13", "5", "/", "+"}, want: 6},
		// {name: "leetcode example 3", tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, want: 22},
		// {name: "single number", tokens: []string{"42"}, want: 42},
		// {name: "single negative", tokens: []string{"-5"}, want: -5},
		// {name: "subtraction order", tokens: []string{"3", "4", "-"}, want: -1},
		// {name: "multiply negatives", tokens: []string{"-3", "-4", "*"}, want: 12},
		// {name: "division truncates toward zero", tokens: []string{"7", "3", "/"}, want: 2},
		// {name: "negative division truncates toward zero", tokens: []string{"-7", "3", "/"}, want: -2},
		// {name: "division opposite signs", tokens: []string{"6", "-4", "/"}, want: -1},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, panicked := call(fn, append([]string(nil), tc.tokens...))
				if panicked {
					t.Fatalf("evalRPN(%v) panicked, want %d", tc.tokens, tc.want)
				}
				if got != tc.want {
					t.Errorf("evalRPN(%v) = %d, want %d", tc.tokens, got, tc.want)
				}
			})
		}
	}
}
