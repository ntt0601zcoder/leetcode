package linkedlistcycle

import (
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(head *ListNode) bool{
	"twopointer": hasCycle,
}

// buildListCycle builds a linked list from vals and, if pos >= 0, links the
// tail's Next to the node at index pos to form a cycle. pos == -1 (or any
// out-of-range value) leaves the list acyclic. Returns the head (nil if empty).
func buildListCycle(vals []int, pos int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	if pos >= 0 && pos < len(nodes) {
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	return nodes[0]
}

// call runs fn with a timeout and panic guard, so a broken cycle check that
// never terminates (e.g. a pointer that runs the cycle forever without ever
// detecting it) reports a clean failure instead of hanging the whole test run.
func call(fn func(*ListNode) bool, head *ListNode) (got bool, outcome string) {
	done := make(chan bool, 1)
	panicked := make(chan struct{}, 1)
	go func() {
		defer func() {
			if recover() != nil {
				panicked <- struct{}{}
			}
		}()
		done <- fn(head)
	}()
	select {
	case got = <-done:
		return got, "ok"
	case <-panicked:
		return false, "panic"
	case <-time.After(500 * time.Millisecond):
		return false, "timeout"
	}
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int
		want bool
	}{
		{name: "leetcode example 1", vals: []int{3, 2, 0, -4}, pos: 1, want: true},
		{name: "leetcode example 2 (two-node cycle)", vals: []int{1, 2}, pos: 0, want: true},
		{name: "leetcode example 3 (single, no cycle)", vals: []int{1}, pos: -1, want: false},
		{name: "empty list", vals: []int{}, pos: -1, want: false},
		{name: "single node self-loop", vals: []int{1}, pos: 0, want: true},
		{name: "two nodes no cycle", vals: []int{1, 2}, pos: -1, want: false},
		{name: "duplicate values no cycle", vals: []int{1, 1, 1, 1}, pos: -1, want: false},
		{name: "two duplicate values no cycle", vals: []int{1, 1}, pos: -1, want: false},
		{name: "even length no cycle", vals: []int{1, 2, 3, 4, 5, 6}, pos: -1, want: false},
		{name: "odd length no cycle", vals: []int{1, 2, 3, 4, 5}, pos: -1, want: false},
		{name: "cycle back to head", vals: []int{1, 2, 3, 4}, pos: 0, want: true},
		{name: "cycle to middle", vals: []int{1, 2, 3, 4, 5}, pos: 2, want: true},
		{name: "tail self-loop", vals: []int{1, 2, 3}, pos: 2, want: true},
		{name: "long list no cycle", vals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, pos: -1, want: false},
		{name: "long list cycle near end", vals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, pos: 8, want: true},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, outcome := call(fn, buildListCycle(tc.vals, tc.pos))
				switch outcome {
				case "panic":
					t.Fatalf("hasCycle(vals=%v, pos=%d) panicked, want %v", tc.vals, tc.pos, tc.want)
				case "timeout":
					t.Fatalf("hasCycle(vals=%v, pos=%d) did not return within 500ms (likely infinite loop), want %v", tc.vals, tc.pos, tc.want)
				}
				if got != tc.want {
					t.Errorf("hasCycle(vals=%v, pos=%d) = %v, want %v", tc.vals, tc.pos, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkHasCycle(b *testing.B) {
	vals := make([]int, 10000)
	for i := range vals {
		vals[i] = i
	}
	acyclic := buildListCycle(vals, -1)
	cyclic := buildListCycle(vals, 0)
	for name, fn := range solutions {
		b.Run(name+"/acyclic", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(acyclic)
			}
		})
		b.Run(name+"/cyclic", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(cyclic)
			}
		})
	}
}
