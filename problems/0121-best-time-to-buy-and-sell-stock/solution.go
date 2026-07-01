// Package besttimetobuyandsellstock solves LeetCode 121. Best Time To Buy And Sell Stock.
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	buy := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else if prices[i]-buy > profit {
			profit = prices[i] - buy
		}
	}

	return profit
}

func maxProfitGreedy(prices []int) int {
	maxProfit := 0

	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]

			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
