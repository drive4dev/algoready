package max_profit

import "math"

// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
// You must buy before you sell.
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	mins := make([]int, len(prices))
	maxs := make([]int, len(prices))

	mins[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		mins[i] = int(math.Min(float64(mins[i-1]), float64(prices[i])))
	}

	maxs[len(prices)-1] = prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		maxs[i] = int(math.Max(float64(maxs[i+1]), float64(prices[i])))
	}

	maxDiff := maxs[0] - mins[0]
	for i := 1; i < len(prices); i++ {
		maxDiff = int(math.Max(float64(maxDiff), float64(maxs[i]-mins[i])))
	}

	return maxDiff
}
