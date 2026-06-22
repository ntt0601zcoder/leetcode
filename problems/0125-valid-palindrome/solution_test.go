package validpalindrome

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(s string) bool{
	"twoPointer": isPalindrome,
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "leetcode example 1", s: "A man, a plan, a canal: Panama", want: true},
		{name: "leetcode example 2", s: "race a car", want: false},
		{name: "leetcode example 3 single space", s: " ", want: true},
		{name: "empty string", s: "", want: true},
		{name: "only non-alphanumeric", s: ".,", want: true},
		{name: "mixed case palindrome", s: "Aa", want: true},
		{name: "single char", s: "a", want: true},
		{name: "single digit", s: "0", want: true},
		{name: "digits palindrome", s: "12321", want: true},
		{name: "digits not palindrome", s: "12345", want: false},
		{name: "alphanumeric mix true", s: "ab_a", want: true},
		{name: "letter digit mismatch", s: "0P", want: false},
		{name: "two same letters", s: "aa", want: true},
		{name: "two different letters", s: "ab", want: false},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := fn(tc.s)
				if got != tc.want {
					t.Errorf("isPalindrome(%q) = %v, want %v", tc.s, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	s := "A man, a plan, a canal: Panama, a man, a plan, a canal: Panama"
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(s)
			}
		})
	}
}
