package minstack

import "testing"

// MinStack is a data-structure design problem, so instead of the approach
// registry used elsewhere we drive a sequence of method calls and assert that
// Top() and GetMin() return the expected value after each step.

// op is a single operation in a scenario. kind is one of "push", "pop",
// "top", "getMin". For "push" arg is the value pushed; for "top"/"getMin"
// want is the expected return value.
type op struct {
	kind string
	arg  int
	want int
}

func TestMinStack(t *testing.T) {
	scenarios := []struct {
		name string
		ops  []op
	}{
		{
			name: "leetcode example",
			ops: []op{
				{kind: "push", arg: -2},
				{kind: "push", arg: 0},
				{kind: "push", arg: -3},
				{kind: "getMin", want: -3},
				{kind: "pop"},
				{kind: "top", want: 0},
				{kind: "getMin", want: -2},
			},
		},
		{
			name: "single element",
			ops: []op{
				{kind: "push", arg: 42},
				{kind: "top", want: 42},
				{kind: "getMin", want: 42},
			},
		},
		{
			name: "duplicate of the min",
			ops: []op{
				{kind: "push", arg: 1},
				{kind: "push", arg: 1},
				{kind: "getMin", want: 1},
				{kind: "pop"},
				{kind: "getMin", want: 1},
				{kind: "top", want: 1},
			},
		},
		{
			name: "pop exposes previous min",
			ops: []op{
				{kind: "push", arg: 5},
				{kind: "push", arg: 3},
				{kind: "push", arg: 7},
				{kind: "getMin", want: 3},
				{kind: "pop"}, // remove 7
				{kind: "getMin", want: 3},
				{kind: "pop"}, // remove 3, min back to 5
				{kind: "getMin", want: 5},
				{kind: "top", want: 5},
			},
		},
		{
			name: "descending then ascending pops",
			ops: []op{
				{kind: "push", arg: 3},
				{kind: "push", arg: 2},
				{kind: "push", arg: 1},
				{kind: "getMin", want: 1},
				{kind: "pop"},
				{kind: "getMin", want: 2},
				{kind: "pop"},
				{kind: "getMin", want: 3},
			},
		},
		{
			name: "negatives and re-push",
			ops: []op{
				{kind: "push", arg: 0},
				{kind: "push", arg: -1},
				{kind: "getMin", want: -1},
				{kind: "pop"},
				{kind: "getMin", want: 0},
				{kind: "push", arg: -5},
				{kind: "getMin", want: -5},
				{kind: "top", want: -5},
			},
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			st := Constructor()
			for i, o := range sc.ops {
				switch o.kind {
				case "push":
					st.Push(o.arg)
				case "pop":
					st.Pop()
				case "top":
					if got := st.Top(); got != o.want {
						t.Errorf("op %d Top() = %d, want %d", i, got, o.want)
					}
				case "getMin":
					if got := st.GetMin(); got != o.want {
						t.Errorf("op %d GetMin() = %d, want %d", i, got, o.want)
					}
				default:
					t.Fatalf("op %d: unknown kind %q", i, o.kind)
				}
			}
		})
	}
}

func BenchmarkMinStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		st := Constructor()
		for v := 0; v < 100; v++ {
			st.Push(100 - v)
		}
		for v := 0; v < 100; v++ {
			st.GetMin()
			st.Top()
			st.Pop()
		}
	}
}
