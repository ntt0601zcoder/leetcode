#!/usr/bin/env bash
# Scaffold a new LeetCode problem directory.
#
# Usage (via Makefile):
#   make new ID=1 SLUG=two-sum
#   make new ID=2 SLUG=add-two-numbers KIND=list TITLE="Add Two Numbers" DIFF=medium
#   make new ID=104 SLUG=maximum-depth-of-binary-tree KIND=tree DIFF=easy
#
# KIND controls the boilerplate injected so each problem stays
# self-contained and paste-ready on leetcode.com:
#   plain (default) — just the solution stub + table-driven test
#   list            — adds a ListNode type + linked-list test helpers
#   tree            — adds a TreeNode type + level-order test helpers
set -euo pipefail

ID="${ID:-}"
SLUG="${SLUG:-}"
KIND="${KIND:-plain}"
TITLE="${TITLE:-}"
DIFF="${DIFF:-}"
URL="${URL:-}"

die() { echo "error: $*" >&2; exit 1; }

[ -n "$ID" ]   || die "ID is required (e.g. ID=1)"
[ -n "$SLUG" ] || die "SLUG is required (e.g. SLUG=two-sum)"
[[ "$ID" =~ ^[0-9]+$ ]] || die "ID must be a number, got '$ID'"
[[ "$SLUG" =~ ^[a-z0-9]+(-[a-z0-9]+)*$ ]] || die "SLUG must be kebab-case, got '$SLUG'"
case "$KIND" in plain|list|tree) ;; *) die "KIND must be plain|list|tree, got '$KIND'";; esac

NNNN=$(printf '%04d' "$ID")
DIR="problems/${NNNN}-${SLUG}"
[ -e "$DIR" ] && die "$DIR already exists"

# Package name: hyphens stripped; prefixed if it would start with a digit.
PKG=$(echo "$SLUG" | tr -d '-')
[[ "$PKG" =~ ^[0-9] ]] && PKG="p${PKG}"

# camelCase func name and PascalCase test name from the slug.
camel() {
  local out="" first=1 part
  IFS='-' read -ra parts <<< "$1"
  for part in "${parts[@]}"; do
    if [ "$first" -eq 1 ]; then out+="$part"; first=0
    else out+="$(tr '[:lower:]' '[:upper:]' <<< "${part:0:1}")${part:1}"; fi
  done
  printf '%s' "$out"
}
FUNC=$(camel "$SLUG")
# Go identifiers can't start with a digit (e.g. slug "3sum"); the real
# LeetCode name won't either, so prefix a valid stub for the user to rename.
[[ "$FUNC" =~ ^[0-9] ]] && FUNC="solve${FUNC}"
TESTNAME="Test$(tr '[:lower:]' '[:upper:]' <<< "${FUNC:0:1}")${FUNC:1}"

[ -n "$TITLE" ] || TITLE=$(echo "$SLUG" | tr '-' ' ' | awk '{for(i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2)}1')
[ -n "$URL" ]   || URL="https://leetcode.com/problems/${SLUG}/"

mkdir -p "$DIR"

# --- README ---
cat > "$DIR/README.md" <<EOF
# ${ID}. ${TITLE}

- Difficulty: ${DIFF:-?}
- Link: ${URL}

## Approach

_TODO: jot down the idea, time/space complexity, and any edge cases._

- Time:  O(?)
- Space: O(?)
EOF

# --- solution.go ---
{
  echo "// Package ${PKG} solves LeetCode ${ID}. ${TITLE}."
  echo "// ${URL}"
  echo "package ${PKG}"
  echo
  if [ "$KIND" = "list" ]; then
    cat <<EOF
// ListNode is the singly-linked list node used by LeetCode.
type ListNode struct {
	Val  int
	Next *ListNode
}

func ${FUNC}(head *ListNode) *ListNode {
	// TODO: implement — adjust the signature to match the problem.
	return head
}
EOF
  elif [ "$KIND" = "tree" ]; then
    cat <<EOF
// TreeNode is the binary tree node used by LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ${FUNC}(root *TreeNode) int {
	// TODO: implement — adjust the signature to match the problem.
	return 0
}
EOF
  else
    cat <<EOF
func ${FUNC}() {
	// TODO: implement — adjust the signature to match the problem.
}
EOF
  fi
} > "$DIR/solution.go"

# --- solution_test.go ---
{
  echo "package ${PKG}"
  echo
  if [ "$KIND" = "list" ]; then
    cat <<EOF
import (
	"reflect"
	"testing"
)

// buildList builds a linked list from a slice of values.
func buildList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// listVals flattens a linked list back into a slice.
func listVals(n *ListNode) []int {
	var out []int
	for ; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

// solutions lists every approach; the test runs all cases against each.
// Add another approach by writing the func and adding a line here.
var solutions = map[string]func(head *ListNode) *ListNode{
	"v1": ${FUNC},
}

func ${TESTNAME}(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "example", in: []int{1, 2, 3}, want: []int{1, 2, 3}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := listVals(fn(buildList(tt.in)))
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.in, got, tt.want)
				}
			})
		}
	}
}
EOF
  elif [ "$KIND" = "tree" ]; then
    cat <<EOF
import "testing"

// buildTree builds a binary tree from a level-order slice. Use a nil
// pointer (via the helper below) for missing children, e.g.
//   buildTree([]*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)})
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

func ptr(v int) *int { return &v }

// solutions lists every approach; the test runs all cases against each.
// Add another approach by writing the func and adding a line here.
var solutions = map[string]func(root *TreeNode) int{
	"v1": ${FUNC},
}

func ${TESTNAME}(t *testing.T) {
	tests := []struct {
		name string
		in   []*int
		want int
	}{
		{name: "example", in: []*int{ptr(3), ptr(9), ptr(20), nil, nil, ptr(15), ptr(7)}, want: 0},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(buildTree(tt.in))
				if got != tt.want {
					t.Errorf("%s = %v, want %v", name, got, tt.want)
				}
			})
		}
	}
}
EOF
  else
    cat <<EOF
import "testing"

// To compare multiple approaches, declare them in a map and range over it,
// e.g. (see problems/0001-two-sum for a full example):
//   var solutions = map[string]func(/* args */) /* ret */{
//       "v1": ${FUNC},
//   }
//   for name, fn := range solutions { for _, tt := range tests { ... } }

func ${TESTNAME}(t *testing.T) {
	tests := []struct {
		name string
		// TODO: input fields
		want int
	}{
		// TODO: add cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: call ${FUNC} and compare against tt.want
			_ = tt
		})
	}
}
EOF
  fi
} > "$DIR/solution_test.go"

gofmt -w "$DIR/solution.go" "$DIR/solution_test.go" 2>/dev/null || true

echo "created $DIR"
echo "  $DIR/solution.go      (func ${FUNC}, package ${PKG})"
echo "  $DIR/solution_test.go (${TESTNAME})"
echo "  $DIR/README.md"
echo
echo "run it:  make test P=${NNNN}"
