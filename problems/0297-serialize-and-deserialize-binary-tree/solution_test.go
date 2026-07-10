package serializeanddeserializebinarytree

import (
	"strconv"
	"testing"
	"time"
)

func ptr(v int) *int { return &v }

// buildTree builds a binary tree from a level-order slice (nil for a missing
// child), e.g. buildTree([]*int{ptr(1), ptr(2), ptr(3), nil, nil, ptr(4), ptr(5)}).
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

// treeEqual reports whether two trees have identical structure and values.
func treeEqual(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && treeEqual(a.Left, b.Left) && treeEqual(a.Right, b.Right)
}

// treeString is a readable preorder rendering (# for nil) for error messages.
func treeString(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "(" + treeString(root.Left) + " " + treeString(root.Right) + ")"
}

// roundTrip serializes with one Codec and deserializes with another (as the
// LeetCode harness does, enforcing a stateless codec). A timeout + panic guard
// keeps a broken codec from hanging or crashing the whole test run.
func roundTrip(root *TreeNode) (got *TreeNode, s string, outcome string) {
	type res struct {
		tree *TreeNode
		str  string
	}
	done := make(chan res, 1)
	panicked := make(chan struct{}, 1)
	go func() {
		defer func() {
			if recover() != nil {
				panicked <- struct{}{}
			}
		}()
		ser := Constructor()
		deser := Constructor()
		str := ser.serialize(root)
		done <- res{deser.deserialize(str), str}
	}()
	select {
	case r := <-done:
		return r.tree, r.str, "ok"
	case <-panicked:
		return nil, "", "panic"
	case <-time.After(time.Second):
		return nil, "", "timeout"
	}
}

func TestCodecRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		vals []*int
	}{
		{name: "leetcode example", vals: []*int{ptr(1), ptr(2), ptr(3), nil, nil, ptr(4), ptr(5)}},
		{name: "single node", vals: []*int{ptr(1)}},
		{name: "empty tree", vals: nil},
		{name: "left skewed", vals: []*int{ptr(1), ptr(2), nil, ptr(3)}},
		{name: "right skewed", vals: []*int{ptr(1), nil, ptr(2), nil, ptr(3)}},
		{name: "complete tree", vals: []*int{ptr(1), ptr(2), ptr(3), ptr(4), ptr(5), ptr(6), ptr(7)}},
		{name: "negative values", vals: []*int{ptr(-1), ptr(-2), ptr(-3)}},
		{name: "value zero", vals: []*int{ptr(5), ptr(0), nil}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := buildTree(tc.vals)
			got, s, outcome := roundTrip(root)
			switch outcome {
			case "panic":
				t.Fatalf("serialize/deserialize panicked for tree %s", treeString(root))
			case "timeout":
				t.Fatalf("serialize/deserialize did not finish within 1s for tree %s", treeString(root))
			}
			if !treeEqual(got, root) {
				t.Errorf("round-trip mismatch\n  tree:      %s\n  serialize: %q\n  got back:  %s", treeString(root), s, treeString(got))
			}
		})
	}
}

func BenchmarkCodecRoundTrip(b *testing.B) {
	vals := make([]*int, 127)
	for i := range vals {
		vals[i] = ptr(i + 1)
	}
	root := buildTree(vals)
	for i := 0; i < b.N; i++ {
		ser := Constructor()
		deser := Constructor()
		deser.deserialize(ser.serialize(root))
	}
}
