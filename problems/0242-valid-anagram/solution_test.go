package validanagram

import "testing"

// solutions lists every approach; the test runs all cases against each.
// Add another approach (e.g. a [26]int counter or sort-based) by writing
// the func and adding a line here.
var solutions = map[string]func(s, t string) bool{
	"map": isAnagramMap,
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "classic anagram", s: "anagram", t: "nagaram", want: true},
		{name: "not anagram", s: "rat", t: "car", want: false},
		{name: "permutation", s: "listen", t: "silent", want: true},
		{name: "identical", s: "abc", t: "abc", want: true},
		{name: "different length", s: "a", t: "ab", want: false},
		{name: "both empty", s: "", t: "", want: true},
		{name: "one empty", s: "a", t: "", want: false},
		{name: "same letters different counts", s: "aabb", t: "abbb", want: false},
		{name: "repeated letters", s: "aacc", t: "ccaa", want: true},
		{name: "superset", s: "aab", t: "aabb", want: false},
		// Unicode follow-up — passes because the solution counts runes.
		{name: "unicode anagram", s: "你好", t: "好你", want: true},
		{name: "unicode not anagram", s: "你好", t: "你你", want: false},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				if got := fn(tc.s, tc.t); got != tc.want {
					t.Errorf("isAnagram(%q, %q) = %v, want %v", tc.s, tc.t, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	const s, t = "anagramanagramanagram", "nagaramnagaramnagaram"
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(s, t)
			}
		})
	}
}
