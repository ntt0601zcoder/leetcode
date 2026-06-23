package removekdigits

import "testing"

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(num string, k int) string{
	"stack": removeKdigits,
}

func TestRemoveKdigits(t *testing.T) {
	tests := []struct {
		name string
		num  string
		k    int
		want string
	}{
		{name: "leetcode example 1", num: "1432219", k: 3, want: "1219"},
		{name: "leetcode example 2 (leading zero)", num: "10200", k: 1, want: "200"},
		{name: "leetcode example 3 (remove all)", num: "10", k: 2, want: "0"},
		{name: "k=0 keeps everything", num: "12345", k: 0, want: "12345"},
		{name: "increasing, trim from end", num: "12345", k: 2, want: "123"},
		{name: "decreasing", num: "54321", k: 3, want: "21"},
		{name: "multiple pops for one digit", num: "4532", k: 2, want: "32"},
		{name: "pop whole run to reach trailing zero", num: "1234567890", k: 9, want: "0"},
		{name: "result collapses to single zero", num: "100", k: 1, want: "0"},
		{name: "single digit removed", num: "9", k: 1, want: "0"},
		{name: "equal digits", num: "112", k: 1, want: "11"},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				if got := fn(tc.num, tc.k); got != tc.want {
					t.Errorf("removeKdigits(%q, %d) = %q, want %q", tc.num, tc.k, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkRemoveKdigits(b *testing.B) {
	// A long zig-zag so the stack does real popping work.
	const num = "928374650192837465019283746501928374650"
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(num, 15)
			}
		})
	}
}
