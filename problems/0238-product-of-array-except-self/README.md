# 238. Product Of Array Except Self

- Difficulty: medium
- Link: https://leetcode.com/problems/product-of-array-except-self/

## Approach

Two passes, no division. Forward pass: `answer[i]` = product of every element
*to the left* of `i`. Backward pass: multiply in `suffix` = product of every
element *to the right* of `i`. The `answer` array doubles as the prefix, so it
uses only O(1) extra space (besides the output). Handles zeros correctly.

- Time:  O(n)
- Space: O(1) extra (besides the output)
