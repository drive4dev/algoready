package max_profit

func maxProfitImp(prices []int) int {
	min := prices[0]
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		localProfit := prices[i] - min
		if localProfit > maxProfit {
			maxProfit = localProfit
		}

		if min > prices[i] {
			min = prices[i]
		}
	}

	return maxProfit
}
