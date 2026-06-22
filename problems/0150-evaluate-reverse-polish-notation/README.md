# 150. Evaluate Reverse Polish Notation

- Difficulty: medium
- Link: https://leetcode.com/problems/evaluate-reverse-polish-notation/

## Approach

Scan the tokens with a stack: push every **number**; on an **operator**, pop
the top two values (`num1` is the deeper one, `num2` is on top), compute
`num1 op num2`, and push the result back. The final stack top is the answer.

Note: operand order matters for `-` and `/` (`num1 op num2`, not the reverse);
division **truncates toward zero**, which matches Go's integer `/`, e.g.
`-7 / 3 == -2`.

- Time:  O(n)
- Space: O(n)
