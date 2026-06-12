# LeetCode (Go)

Luyện LeetCode bằng Go. Mỗi bài là một package riêng trong `problems/`,
đặt tên theo số bài để dễ tra cứu, và tự chứa (self-contained) nên có thể
copy thẳng `solution.go` lên leetcode.com.

## Cấu trúc

```text
.
├── Makefile                  # các shortcut: new, test, bench, ...
├── go.mod                    # module gốc, go test ./... chạy mọi bài
├── scripts/new.sh            # sinh khung bài mới
└── problems/
    └── 0001-two-sum/
        ├── solution.go       # lời giải (paste thẳng lên LeetCode được)
        ├── solution_test.go  # table-driven test + benchmark
        └── README.md         # đề bài, ý tưởng, độ phức tạp
```

## Bắt đầu một bài mới

```bash
# bài thường
make new ID=1 SLUG=two-sum

# bài dùng linked list (tự thêm ListNode + helper build/flatten)
make new ID=2 SLUG=add-two-numbers KIND=list DIFF=medium

# bài dùng cây nhị phân (tự thêm TreeNode + builder level-order)
make new ID=104 SLUG=maximum-depth-of-binary-tree KIND=tree DIFF=easy
```

Tham số: `ID` (số bài), `SLUG` (kebab-case), `KIND=plain|list|tree`,
`TITLE`, `DIFF=easy|medium|hard`, `URL`.

## Vòng lặp luyện tập

```bash
make new ID=1 SLUG=two-sum   # 1. sinh khung
# 2. sửa solution.go + thêm test case
make test P=0001             # 3. chạy test riêng bài đó
make test                    # chạy toàn bộ
```

## Nhiều lời giải cho cùng 1 bài

Cứ giữ tất cả trong cùng package (cùng thư mục bài), mỗi cách một hàm — có
thể tách thành nhiều file (`bruteforce.go`, `hashmap.go`...) cho gọn. Khai
báo một map `solutions` rồi cho test chạy *mọi cách* qua cùng bộ case, và
benchmark so sánh tốc độ. Xem `problems/0001-two-sum` làm mẫu:

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

Thêm cách mới = viết thêm 1 hàm + thêm 1 dòng vào map. Scaffold `KIND=list`
và `KIND=tree` đã sinh sẵn map này; `KIND=plain` có comment hướng dẫn.

```bash
make test P=0001    # chạy tất cả các cách: brute/basic, hashmap/basic, ...
make bench P=0001   # so sánh tốc độ giữa các cách
```

## Các lệnh khác

```bash
make help        # liệt kê mọi lệnh
make test P=...  # test 1 bài (lọc theo số hoặc slug)
make test-race   # test với race detector
make bench P=... # chạy benchmark
make cover P=... # coverage + mở báo cáo HTML
make verify      # fmt + vet + test toàn bộ
make list        # liệt kê các bài đã làm
make count       # đếm số bài
```

## Quy ước

- **Package** = slug bỏ dấu gạch (vd `two-sum` → `twosum`); thêm tiền tố
  `p` nếu bắt đầu bằng số (vd `3sum` → `p3sum`).
- **Hàm giải** đặt theo camelCase của slug (vd `twoSum`), viết thường để
  khớp đúng chữ ký LeetCode.
- **Test** là white-box (cùng package) nên gọi được hàm viết thường.
