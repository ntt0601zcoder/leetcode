# LeetCode (Go)

LeetCode practice in Go. Each problem is its own package under `problems/`,
named by problem number for easy lookup, and self-contained so you can paste
`solution.go` straight into leetcode.com.

## Layout

```text
.
├── Makefile                  # shortcuts: new, test, bench, ...
├── go.mod                    # root module; go test ./... runs every problem
├── scripts/new.sh            # scaffold a new problem
└── problems/
    └── 0001-two-sum/
        ├── solution.go       # the solution (paste-ready for LeetCode)
        ├── solution_test.go  # table-driven test + benchmark
        └── README.md         # statement, idea, complexity
```

## Start a new problem

```bash
# plain problem
make new ID=1 SLUG=two-sum

# linked-list problem (adds ListNode + build/flatten helpers)
make new ID=2 SLUG=add-two-numbers KIND=list DIFF=medium

# binary-tree problem (adds TreeNode + level-order builder)
make new ID=104 SLUG=maximum-depth-of-binary-tree KIND=tree DIFF=easy
```

Parameters: `ID` (problem number), `SLUG` (kebab-case), `KIND=plain|list|tree`,
`TITLE`, `DIFF=easy|medium|hard`, `URL`.

## Practice loop

```bash
make new ID=1 SLUG=two-sum   # 1. scaffold
# 2. edit solution.go + add test cases
make test P=0001             # 3. run just that problem
make test                    # run everything
```

## Multiple solutions for one problem

Keep them all in the same package (the problem's directory), one function per
approach — you can split them into several files (`bruteforce.go`,
`hashmap.go`...) if you like. Declare a `solutions` map and let the test run
*every* approach against the same cases, plus a benchmark to compare speed.
See `problems/0001-two-sum` for a full example:

```go
var solutions = map[string]func(nums []int, target int) []int{
    "brute":   twoSumBrute,
    "hashmap": twoSumHashMap,
}

for name, fn := range solutions {
    for _, tt := range tests {
        t.Run(name+"/"+tt.name, func(t *testing.T) { /* fn(...) vs tt.want */ })
    }
}
```

Adding an approach = one more function + one line in the map. The `KIND=list`
and `KIND=tree` scaffolds generate this map for you; `KIND=plain` includes a
guiding comment.

```bash
make test P=0001    # run every approach: brute/basic, hashmap/basic, ...
make bench P=0001   # compare speed across approaches
```

## Other commands

```bash
make help        # list every command
make test P=...  # test one problem (filter by number or slug)
make test-race   # test with the race detector
make bench P=... # run benchmarks
make cover P=... # coverage + open the HTML report
make verify      # fmt + vet + test everything
make list        # list solved problems
make count       # count problems
```

## Conventions

- **Package** = slug with hyphens removed (e.g. `two-sum` → `twosum`); prefixed
  with `p` if it would start with a digit (e.g. `3sum` → `p3sum`).
- **Solution function** is the lowerCamelCase of the slug (e.g. `twoSum`),
  lowercase to match LeetCode's signature exactly.
- **Tests** are white-box (same package) so they can call the unexported func.
