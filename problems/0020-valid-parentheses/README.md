# 20. Valid Parentheses

- Difficulty: easy
- Link: https://leetcode.com/problems/valid-parentheses/

## Approach

Scan the string with a stack. Push every opening bracket; on a closing bracket,
check that the stack is non-empty and its top is the matching opener, then pop
it. Any closing bracket that does not match the top (or that arrives on an empty
stack) makes the string invalid immediately. The string is valid only if the
stack is empty at the end, meaning every opener was closed in the right order.

- Time:  O(n)
- Space: O(n)
