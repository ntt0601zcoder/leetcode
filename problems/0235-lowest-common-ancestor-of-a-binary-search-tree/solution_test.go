package lowestcommonancestorofabinarysearchtree

import (
	"testing"
	"time"
)

// solutions lists every approach; the test runs all cases against each.
var solutions = map[string]func(root, p, q *TreeNode) *TreeNode{
	"iterative": lowestCommonAncestor,
}

func ptr(v int) *int { return &v }

// buildTree builds a binary tree from a level-order slice, nil for a missing
// child, e.g. buildTree([]*int{ptr(6), ptr(2), ptr(8)}).
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

// find returns the pointer to the node holding val by walking the BST, so tests
// can hand the solution pointers to existing nodes (p and q are guaranteed to
// exist in the tree by the problem constraints).
func find(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil {
		switch {
		case val < node.Val:
			node = node.Left
		case val > node.Val:
			node = node.Right
		default:
			return node
		}
	}
	return nil
}

// call runs fn with a timeout and panic guard so a non-terminating walk or a
// nil dereference reports a clean failure instead of hanging or crashing the run.
func call(fn func(root, p, q *TreeNode) *TreeNode, root, p, q *TreeNode) (got *TreeNode, outcome string) {
	done := make(chan *TreeNode, 1)
	panicked := make(chan struct{}, 1)
	go func() {
		defer func() {
			if recover() != nil {
				panicked <- struct{}{}
			}
		}()
		done <- fn(root, p, q)
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

// leetcodeTree is {6,2,8,0,4,7,9,null,null,3,5}.
var leetcodeTree = []*int{ptr(6), ptr(2), ptr(8), ptr(0), ptr(4), ptr(7), ptr(9), nil, nil, ptr(3), ptr(5)}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name string
		vals []*int
		p    int
		q    int
		want int // LCA value (LCA is unique)
	}{
		{name: "leetcode example 1 split at root", vals: leetcodeTree, p: 2, q: 8, want: 6},
		{name: "leetcode example 2 ancestor is p", vals: leetcodeTree, p: 2, q: 4, want: 2},
		{name: "p equals q", vals: leetcodeTree, p: 2, q: 2, want: 2},
		{name: "both deep in left subtree", vals: leetcodeTree, p: 0, q: 5, want: 2},
		{name: "siblings under 4", vals: leetcodeTree, p: 3, q: 5, want: 4},
		{name: "siblings under 8", vals: leetcodeTree, p: 7, q: 9, want: 8},
		{name: "root is one of the nodes", vals: leetcodeTree, p: 6, q: 3, want: 6},
		{name: "ancestor is 4 for descendant", vals: leetcodeTree, p: 4, q: 3, want: 4},
		{name: "single node tree", vals: []*int{ptr(1)}, p: 1, q: 1, want: 1},
		{name: "two node tree lca is root", vals: []*int{ptr(2), ptr(1)}, p: 1, q: 2, want: 2},
		{name: "right skewed", vals: []*int{ptr(1), nil, ptr(2), nil, ptr(3)}, p: 2, q: 3, want: 2},
		{name: "negatives split at root", vals: []*int{ptr(0), ptr(-5), ptr(5), ptr(-10), ptr(-2)}, p: -10, q: 5, want: 0},
	}
	for name, fn := range solutions {
		for _, tc := range tests {
			t.Run(name+"/"+tc.name, func(t *testing.T) {
				root := buildTree(tc.vals)
				p := find(root, tc.p)
				q := find(root, tc.q)
				if p == nil || q == nil {
					t.Fatalf("test setup: p=%d or q=%d not present in tree %v", tc.p, tc.q, tc.vals)
				}
				got, outcome := call(fn, root, p, q)
				switch outcome {
				case "panic":
					t.Fatalf("lowestCommonAncestor(tree=%v, p=%d, q=%d) panicked, want %d", tc.vals, tc.p, tc.q, tc.want)
				case "timeout":
					t.Fatalf("lowestCommonAncestor(tree=%v, p=%d, q=%d) did not return within 500ms (likely infinite loop), want %d", tc.vals, tc.p, tc.q, tc.want)
				}
				if got == nil {
					t.Fatalf("lowestCommonAncestor(tree=%v, p=%d, q=%d) = nil, want node %d", tc.vals, tc.p, tc.q, tc.want)
				}
				if got.Val != tc.want {
					t.Errorf("lowestCommonAncestor(tree=%v, p=%d, q=%d) = %d, want %d", tc.vals, tc.p, tc.q, got.Val, tc.want)
				}
			})
		}
	}
}

func BenchmarkLowestCommonAncestor(b *testing.B) {
	// A balanced BST of 1..127 laid out in level order.
	// For a balanced BST built level-order, index 0 holds 64, etc. Simpler:
	// build a right-leaning BST and query two leaves.
	root := buildTree(leetcodeTree)
	p := find(root, 3)
	q := find(root, 9)
	for name, fn := range solutions {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fn(root, p, q)
			}
		})
	}
}
