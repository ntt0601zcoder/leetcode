package groupanagrams

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(strs []string) [][]string{
	"map": groupAnagramsMap,
}

// normalize makes [][]string order-independent so the test does not depend
// on map iteration order: sort within each group, then sort the groups.
func normalize(groups [][]string) [][]string {
	out := make([][]string, len(groups))
	for i, g := range groups {
		c := append([]string(nil), g...)
		sort.Strings(c)
		out[i] = c
	}
	sort.Slice(out, func(i, j int) bool {
		return strings.Join(out[i], "\x00") < strings.Join(out[j], "\x00")
	})
	return out
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want [][]string
	}{
		{
			name: "classic",
			in:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{name: "single empty string", in: []string{""}, want: [][]string{{""}}},
		{name: "single letter", in: []string{"a"}, want: [][]string{{"a"}}},
		{
			name: "all anagrams of each other",
			in:   []string{"abc", "bca", "cab"},
			want: [][]string{{"abc", "bca", "cab"}},
		},
		{
			name: "no anagrams",
			in:   []string{"abc", "def", "ghi"},
			want: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{name: "empty input", in: []string{}, want: [][]string{}},
		{name: "duplicate strings group together", in: []string{"a", "a"}, want: [][]string{{"a", "a"}}},
		{name: "multiple empty strings", in: []string{"", "", ""}, want: [][]string{{"", "", ""}}},
		{
			name: "mixed groups and singletons",
			in:   []string{"listen", "silent", "google", "banana"},
			want: [][]string{{"listen", "silent"}, {"google"}, {"banana"}},
		},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got := normalize(fn(tc.in))
				want := normalize(tc.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("groupAnagrams(%q) = %v, want %v (order-independent)", tc.in, got, want)
				}
			})
		}
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "abc", "cba", "xyz", "zyx"}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(strs)
			}
		})
	}
}
