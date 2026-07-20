package mergetwosortedlists

import (
	"reflect"
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(list1, list2 *ListNode) *ListNode{
	"iterative": mergeTwoLists,
}

// buildList builds a linked list from a slice; nil for an empty slice.
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

// listVals walks the list into a slice. It stops after a generous bound so a
// solution that accidentally splices a cycle reports a clean failure via the
// timeout guard rather than looping forever here.
func listVals(head *ListNode) []int {
	var out []int
	for n, guard := head, 0; n != nil && guard < 100000; n, guard = n.Next, guard+1 {
		out = append(out, n.Val)
	}
	return out
}

// call runs fn with a timeout and panic guard, so a solution that builds a
// cycle (infinite loop while reading the result) or dereferences nil reports a
// clean failure instead of hanging or crashing the whole test run.
func call(fn func(*ListNode, *ListNode) *ListNode, l1, l2 *ListNode) (got []int, outcome string) {
	done := make(chan []int, 1)
	panicked := make(chan struct{}, 1)
	go func() {
		defer func() {
			if recover() != nil {
				panicked <- struct{}{}
			}
		}()
		done <- listVals(fn(l1, l2))
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

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{name: "leetcode example 1", list1: []int{1, 2, 4}, list2: []int{1, 3, 4}, want: []int{1, 1, 2, 3, 4, 4}},
		{name: "both empty", list1: nil, list2: nil, want: nil},
		{name: "list1 empty", list1: nil, list2: []int{0}, want: []int{0}},
		{name: "list2 empty", list1: []int{0}, list2: nil, want: []int{0}},
		{name: "single each", list1: []int{2}, list2: []int{1}, want: []int{1, 2}},
		{name: "single each equal", list1: []int{5}, list2: []int{5}, want: []int{5, 5}},
		{name: "interleaved", list1: []int{1, 3, 5}, list2: []int{2, 4, 6}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "list1 all smaller", list1: []int{1, 2, 3}, list2: []int{4, 5, 6}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "list1 much longer", list1: []int{1, 2, 3, 4, 5}, list2: []int{6}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "list2 much longer", list1: []int{6}, list2: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "duplicates across", list1: []int{1, 1, 1}, list2: []int{1, 1}, want: []int{1, 1, 1, 1, 1}},
		{name: "negatives", list1: []int{-5, -1, 0}, list2: []int{-3, -2, 4}, want: []int{-5, -3, -2, -1, 0, 4}},
		{name: "stable on equal keys", list1: []int{2, 2}, list2: []int{2, 2}, want: []int{2, 2, 2, 2}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, outcome := call(fn, buildList(tc.list1), buildList(tc.list2))
				switch outcome {
				case "panic":
					t.Fatalf("mergeTwoLists(%v, %v) panicked, want %v", tc.list1, tc.list2, tc.want)
				case "timeout":
					t.Fatalf("mergeTwoLists(%v, %v) did not return within 500ms (likely a cycle), want %v", tc.list1, tc.list2, tc.want)
				}
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("mergeTwoLists(%v, %v) = %v, want %v", tc.list1, tc.list2, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkMergeTwoLists(b *testing.B) {
	odds := make([]int, 1000)
	evens := make([]int, 1000)
	for i := range odds {
		odds[i] = 2*i + 1
		evens[i] = 2 * i
	}
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Rebuild each iteration: mergeTwoLists splices nodes in place.
				fn(buildList(odds), buildList(evens))
			}
		})
	}
}
