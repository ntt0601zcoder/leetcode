package maximumdepthofbinarytree

import (
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(root *TreeNode) int{
	"recursive": maxDepth,
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

// call runs fn with a timeout and panic guard, so a non-terminating or panicking
// solution reports a clean failure instead of hanging the whole test run.
func call(fn func(*TreeNode) int, root *TreeNode) (got int, outcome string) {
	done := make(chan int, 1)
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
		return 0, "panic"
	case <-time.After(500 * time.Millisecond):
		return 0, "timeout"
	}
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		vals []*int
		want int
	}{
		{name: "leetcode example 1", vals: []*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)}, want: 3},
		{name: "leetcode example 2", vals: []*int{ptr(1), nil, ptr(2)}, want: 2},
		{name: "empty tree", vals: nil, want: 0},
		{name: "single node", vals: []*int{ptr(1)}, want: 1},
		{name: "balanced full", vals: []*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6), ptr(7)}, want: 3},
		{name: "left skewed", vals: []*int{ptr(1), ptr(2), nil, ptr(3), nil, ptr(4)}, want: 4},
		{name: "right skewed", vals: []*int{ptr(1), nil, ptr(2), nil, ptr(3), nil, ptr(4)}, want: 4},
		{name: "unbalanced deeper left", vals: []*int{ptr(1), ptr(2), ptr(3), ptr(4), nil, nil, nil, ptr(5)}, want: 4},
		{name: "unbalanced deeper right", vals: []*int{ptr(1), ptr(2), ptr(3), nil, nil, nil, ptr(4), nil, ptr(5)}, want: 4},
		{name: "two levels partial", vals: []*int{ptr(1), ptr(2)}, want: 2},
		{name: "negatives", vals: []*int{ptr(-3), ptr(-9), ptr(-20), nil, nil, ptr(-15), ptr(-7)}, want: 3},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				got, outcome := call(fn, buildTree(tc.vals))
				switch outcome {
				case "panic":
					t.Fatalf("maxDepth(%v) panicked, want %d", tc.vals, tc.want)
				case "timeout":
					t.Fatalf("maxDepth(%v) did not return within 500ms (likely infinite loop), want %d", tc.vals, tc.want)
				}
				if got != tc.want {
					t.Errorf("maxDepth(%v) = %d, want %d", tc.vals, got, tc.want)
				}
			})
		}
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	// A balanced tree of 1..127 (depth 7).
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
