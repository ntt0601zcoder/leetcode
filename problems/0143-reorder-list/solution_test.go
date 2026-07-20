package reorderlist

import (
	"reflect"
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
// reorderList mutates the list in place and returns nothing, so each case
// builds a fresh list and reads the result back with listVals.
var solutions = map[string]func(head *ListNode){
	"reorderList": reorderList,
}

// buildList builds a singly linked list from vals (nil for an empty slice).
// ListNode is declared in solution.go, so it is not redeclared here.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

// listVals reads a list back into a slice. It runs inside the guarded
// goroutine so that a cyclic list (a real failure mode of a broken reorder)
// is caught by the timeout instead of hanging forever.
func listVals(head *ListNode) []int {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}
	return vals
}

// call runs fn on head with a timeout and panic guard. reorderList mutates
// in place, so the result is read (via listVals) inside the same goroutine.
// A nil-pointer deref reports "panic"; an infinite loop / accidental cycle
// reports "timeout" instead of hanging the whole test run.
func call(fn func(*ListNode), head *ListNode) (got []int, outcome string) {
	type result struct {
		vals     []int
		didPanic bool
	}
	done := make(chan result, 1)
	go func() {
		var res result
		defer func() {
			if recover() != nil {
				res.didPanic = true
			}
			done <- res
		}()
		fn(head)
		res.vals = listVals(head)
	}()
	select {
	case r := <-done:
		if r.didPanic {
			return nil, "panic"
		}
		return r.vals, "ok"
	case <-time.After(500 * time.Millisecond):
		return nil, "timeout"
	}
}

func TestReorderList(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "leetcode example 1 (even)", in: []int{1, 2, 3, 4}, want: []int{1, 4, 2, 3}},
		{name: "leetcode example 2 (odd)", in: []int{1, 2, 3, 4, 5}, want: []int{1, 5, 2, 4, 3}},
		{name: "empty", in: []int{}, want: []int{}},
		{name: "single", in: []int{1}, want: []int{1}},
		{name: "two", in: []int{1, 2}, want: []int{1, 2}},
		{name: "three (odd)", in: []int{1, 2, 3}, want: []int{1, 3, 2}},
		{name: "six (even)", in: []int{1, 2, 3, 4, 5, 6}, want: []int{1, 6, 2, 5, 3, 4}},
		{name: "seven (odd)", in: []int{1, 2, 3, 4, 5, 6, 7}, want: []int{1, 7, 2, 6, 3, 5, 4}},
		{name: "duplicates", in: []int{2, 2, 2, 2}, want: []int{2, 2, 2, 2}},
		{name: "negatives (odd)", in: []int{-1, -2, -3, -4, -5}, want: []int{-1, -5, -2, -4, -3}},
		{name: "unsorted values (even)", in: []int{1000, 1, 999, 2}, want: []int{1000, 2, 1, 999}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, outcome := call(fn, buildList(tc.in))
				switch outcome {
				case "panic":
					t.Fatalf("reorderList(%v) panicked, want %v", tc.in, tc.want)
				case "timeout":
					t.Fatalf("reorderList(%v) did not return within 500ms (likely infinite loop / cycle), want %v", tc.in, tc.want)
				}
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("reorderList(%v) = %v, want %v", tc.in, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkReorderList(b *testing.B) {
	vals := make([]int, 2000)
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
