package taskscheduler

import (
	"testing"
	"time"
)

// solutions lists every approach; the table test runs all cases against each.
var solutions = map[string]func(tasks []byte, n int) int{
	"heap": leastInterval,
	"math": leastIntervalMath,
}

// call runs fn in a goroutine guarded by a timeout so a non-terminating
// solution (the heap variant re-pushes leftovers each frame; a bad idle/pop
// condition could loop forever with growing memory) reports a clean failure
// instead of hanging or OOM-ing the whole test run. A panic is recovered and
// surfaced too.
func call(fn func([]byte, int) int, tasks []byte, n int) (got int, panicked, timedOut bool) {
	type result struct {
		val      int
		panicked bool
	}
	done := make(chan result, 1)
	go func() {
		var r result
		defer func() {
			if recover() != nil {
				r.panicked = true
			}
			done <- r
		}()
		r.val = fn(tasks, n)
	}()
	select {
	case r := <-done:
		return r.val, r.panicked, false
	case <-time.After(500 * time.Millisecond):
		return 0, false, true
	}
}

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		name  string
		tasks string
		n     int
		want  int
	}{
		{name: "leetcode example 1", tasks: "AAABBB", n: 2, want: 8},
		{name: "n=0 no cooldown", tasks: "AAABBB", n: 0, want: 6},
		{name: "n=1 forces idles", tasks: "AAABBB", n: 1, want: 6},
		{name: "all distinct small n", tasks: "ABCDEF", n: 2, want: 6},
		{name: "all distinct large n", tasks: "ABCD", n: 100, want: 4},
		{name: "single type triple", tasks: "AAA", n: 2, want: 7},
		{name: "single type no cooldown", tasks: "AAA", n: 0, want: 3},
		{name: "single task", tasks: "A", n: 5, want: 1},
		{name: "two same large n", tasks: "AA", n: 2, want: 4},
		{name: "one dominant frequency", tasks: "AAAAABC", n: 2, want: 13},
		{name: "perfect fill no idle", tasks: "AABBCC", n: 2, want: 6},
		{name: "three equal triples", tasks: "AAABBBCCC", n: 2, want: 9},
		{name: "idles at tail mid-frame", tasks: "AAABBCD", n: 3, want: 9},
		{name: "mixed counts n=1", tasks: "ACABDB", n: 1, want: 6},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				// copy input in case a solution mutates it in place
				in := append([]byte(nil), []byte(tc.tasks)...)
				got, panicked, timedOut := call(fn, in, tc.n)
				if timedOut {
					t.Fatalf("leastInterval(%q, %d) did not return within 500ms (likely infinite loop), want %d", tc.tasks, tc.n, tc.want)
				}
				if panicked {
					t.Fatalf("leastInterval(%q, %d) panicked, want %d", tc.tasks, tc.n, tc.want)
				}
				if got != tc.want {
					t.Errorf("leastInterval(%q, %d) = %d, want %d", tc.tasks, tc.n, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkLeastInterval(b *testing.B) {
	// modest, valid input: 26 task types with skewed frequencies.
	var tasks []byte
	for c := byte('A'); c <= 'Z'; c++ {
		reps := int('Z'-c) + 1
		for r := 0; r < reps; r++ {
			tasks = append(tasks, c)
		}
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				in := append([]byte(nil), tasks...)
				fn(in, 10)
			}
		})
	}
}
