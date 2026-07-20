package reverselinkedlist

import (
	"reflect"
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(head *ListNode) *ListNode{
	"iterative": reverseList,
}

// buildList builds a singly linked list from vals and returns its head.
// An empty slice yields a nil head.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

// listVals walks the list and collects its values into a slice (nil for an
// empty list). It bounds the walk so a solution that leaves a cycle in the
// reversed list can't spin forever here.
func listVals(head *ListNode) []int {
	var out []int
	for i := 0; head != nil && i < 100000; i++ {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

// call runs fn with a timeout and panic guard, so a solution that dereferences
// a nil pointer or leaves a cycle (infinite loop while collecting) reports a
// clean failure instead of hanging or crashing the whole test run.
func call(fn func(*ListNode) *ListNode, head *ListNode) (got []int, outcome string) {
	done := make(chan []int, 1)
	panicked := make(chan struct{}, 1)
	go func() {
		defer func() {
			if recover() != nil {
				panicked <- struct{}{}
			}
		}()
		done <- listVals(fn(head))
	}()
	select {
	case got = <-done:
		return got, "ok"
	case <-panicked:
		return nil, "panic"
	case <-time.After(500 * time.Millisecond):
		return nil, "timeout"
	}
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "leetcode example 1", in: []int{1, 2, 3, 4, 5}, want: []int{5, 4, 3, 2, 1}},
		{name: "leetcode example 2 (two nodes)", in: []int{1, 2}, want: []int{2, 1}},
		{name: "leetcode example 3 (empty)", in: nil, want: nil},
		{name: "single node", in: []int{42}, want: []int{42}},
		{name: "duplicates", in: []int{7, 7, 7, 7}, want: []int{7, 7, 7, 7}},
		{name: "negatives", in: []int{-1, -2, -3}, want: []int{-3, -2, -1}},
		{name: "mixed signs", in: []int{-5, 0, 5, -10, 10}, want: []int{10, -10, 5, 0, -5}},
		{name: "three nodes", in: []int{1, 2, 3}, want: []int{3, 2, 1}},
		{name: "palindrome values", in: []int{1, 2, 1}, want: []int{1, 2, 1}},
		{name: "longer run", in: []int{10, 20, 30, 40, 50, 60}, want: []int{60, 50, 40, 30, 20, 10}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, outcome := call(fn, buildList(tc.in))
				switch outcome {
				case "panic":
					t.Fatalf("reverseList(%v) panicked, want %v", tc.in, tc.want)
				case "timeout":
					t.Fatalf("reverseList(%v) did not return within 500ms (likely a cycle/infinite loop), want %v", tc.in, tc.want)
				}
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("reverseList(%v) = %v, want %v", tc.in, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkReverseList(b *testing.B) {
	vals := make([]int, 1000)
	for i := range vals {
		vals[i] = i
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				head := buildList(vals)
				b.StartTimer()
				fn(head)
			}
		})
	}
}
