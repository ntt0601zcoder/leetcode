package binarytreelevelordertraversal

import (
	"reflect"
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(root *TreeNode) [][]int{
	"bfs": levelOrder,
}

func ptr(v int) *int { return &v }

// buildTree builds a binary tree from a level-order slice, nil for a missing
// child, e.g. buildTree([]*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)}).
func buildTree(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// call runs fn with a timeout and panic guard, so a non-terminating solution
// (e.g. a BFS that never dequeues the queue) reports a clean failure instead
// of hanging the whole test run.
func call(fn func(*TreeNode) [][]int, root *TreeNode) (got [][]int, outcome string) {
	done := make(chan [][]int, 1)
	panicked := make(chan struct{}, 1)
	go func() {
		defer func() {
			if recover() != nil {
				panicked <- struct{}{}
			}
		}()
		done <- fn(root)
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

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		vals []*int
		want [][]int
	}{
		{name: "leetcode example 1", vals: []*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)}, want: [][]int{{3}, {9, 20}, {15, 7}}},
		{name: "single node", vals: []*int{ptr(1)}, want: [][]int{{1}}},
		{name: "empty tree", vals: nil, want: nil},
		{name: "left skewed", vals: []*int{ptr(1), ptr(2), nil, ptr(3)}, want: [][]int{{1}, {2}, {3}}},
		{name: "complete tree", vals: []*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6), ptr(7)}, want: [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
		{name: "gaps in last level", vals: []*int{ptr(1), ptr(2), ptr(3), nil, ptr(4), nil, ptr(5)}, want: [][]int{{1}, {2, 3}, {4, 5}}},
		{name: "negatives", vals: []*int{ptr(0), ptr(-1), ptr(-2)}, want: [][]int{{0}, {-1, -2}}},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, outcome := call(fn, buildTree(tc.vals))
				switch outcome {
				case "panic":
					t.Fatalf("levelOrder(%v) panicked, want %v", tc.vals, tc.want)
				case "timeout":
					t.Fatalf("levelOrder(%v) did not return within 500ms (likely infinite loop), want %v", tc.vals, tc.want)
				}
				if !reflect.DeepEqual(got, tc.want) {
					t.Errorf("levelOrder(%v) = %v, want %v", tc.vals, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkLevelOrder(b *testing.B) {
	// A balanced tree of 1..127.
	vals := make([]*int, 127)
	for i := range vals {
		vals[i] = ptr(i + 1)
	}
	root := buildTree(vals)
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(root)
			}
		})
	}
}
