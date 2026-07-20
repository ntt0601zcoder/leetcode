package mergeksortedlists

import (
	"reflect"
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
// ListNode is declared in solution.go, so it is not redeclared here.
var solutions = map[string]func(lists []*ListNode) *ListNode{
	"divideconquer": mergeKLists,
}

// buildList builds a singly linked list from vals. An empty slice yields nil.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

// buildLists turns a slice of value slices into a slice of list heads. A nil
// outer slice stays nil (the "no lists" case); an empty inner slice becomes a
// nil list head (the "empty list" case).
func buildLists(lists [][]int) []*ListNode {
	if lists == nil {
		return nil
	}
	out := make([]*ListNode, len(lists))
	for i, l := range lists {
		out[i] = buildList(l)
	}
	return out
}

// listVals collects the values of a list. It is bounded so a cycle-producing
// bug (easy to introduce in a linked-list merge) reports a failure via the
// timeout guard instead of growing memory without limit and OOMing the runner.
func listVals(head *ListNode) []int {
	var vals []int
	for n := head; n != nil; n = n.Next {
		vals = append(vals, n.Val)
		if len(vals) > 1<<20 {
			break
		}
	}
	return vals
}

// call runs fn with a timeout and panic guard. A merge that panics (e.g. nil
// deref) or never terminates (e.g. a cycle in the merged list) reports a clean
// failure instead of hanging or crashing the whole test run.
func call(fn func([]*ListNode) *ListNode, lists []*ListNode) (got []int, outcome string) {
	done := make(chan []int, 1)
	panicked := make(chan struct{}, 1)
	go func() {
		defer func() {
			if recover() != nil {
				panicked <- struct{}{}
			}
		}()
		done <- listVals(fn(lists))
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

func TestMergeKSortedLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{name: "leetcode example 1", lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, want: []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{name: "leetcode example 2 (no lists)", lists: [][]int{}, want: nil},
		{name: "leetcode example 3 (single empty list)", lists: [][]int{{}}, want: nil},
		{name: "nil slice of lists", lists: nil, want: nil},
		{name: "all empty lists", lists: [][]int{{}, {}, {}}, want: nil},
		{name: "single non-empty list", lists: [][]int{{1, 2, 3}}, want: []int{1, 2, 3}},
		{name: "two interleaved lists", lists: [][]int{{1, 3, 5}, {2, 4, 6}}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "empty lists mixed in", lists: [][]int{{}, {1, 5}, {}, {2, 3}}, want: []int{1, 2, 3, 5}},
		{name: "leading empty lists", lists: [][]int{{}, {}, {7, 8, 9}}, want: []int{7, 8, 9}},
		{name: "duplicates across lists", lists: [][]int{{1, 1, 2}, {1, 2, 2}, {2}}, want: []int{1, 1, 1, 2, 2, 2, 2}},
		{name: "negatives", lists: [][]int{{-3, -1, 2}, {-2, 0, 5}}, want: []int{-3, -2, -1, 0, 2, 5}},
		{name: "odd count, single-element lists", lists: [][]int{{5}, {1}, {3}, {2}, {4}}, want: []int{1, 2, 3, 4, 5}},
		{name: "three single-element lists", lists: [][]int{{1}, {2}, {3}}, want: []int{1, 2, 3}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, outcome := call(fn, buildLists(tc.lists))
				switch outcome {
				case "panic":
					t.Fatalf("mergeKLists(%v) panicked, want %v", tc.lists, tc.want)
				case "timeout":
					t.Fatalf("mergeKLists(%v) did not return within 500ms (likely a cycle/infinite loop), want %v", tc.lists, tc.want)
				}
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("mergeKLists(%v) = %v, want %v", tc.lists, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkMergeKSortedLists(b *testing.B) {
	const k, n = 50, 200 // k lists, n nodes each; values interleaved across lists.
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Rebuild each iteration: the merge relinks nodes in place, so a
				// reused input would be mutated (and could alias into cycles).
				b.StopTimer()
				lists := make([]*ListNode, k)
				for j := 0; j < k; j++ {
					vals := make([]int, n)
					for x := 0; x < n; x++ {
						vals[x] = x*k + j
					}
					lists[j] = buildList(vals)
				}
				b.StartTimer()
				fn(lists)
			}
		})
	}
}
