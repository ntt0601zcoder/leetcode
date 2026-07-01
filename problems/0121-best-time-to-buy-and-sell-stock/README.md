# 121. Best Time To Buy And Sell Stock

- Difficulty: easy
- Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

## Approach

One pass. Track the lowest price seen so far (the best day to have bought) and
the best profit `price - minSoFar` when selling on the current day. Update both
as you scan. Profit is 0 if prices never rise (never sell at a loss).

A brute-force O(n²) variant (`maxProfitGreedy`) checks every buy/sell pair; the
one-pass version above is the intended O(n) solution.

- Time:  O(n) one-pass / O(n²) brute
- Space: O(1)
