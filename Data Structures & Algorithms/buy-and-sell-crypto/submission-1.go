func maxProfit(prices []int) int {
	var (
		maxProfit, lowestPrice int = 0, prices[0]
	)

	for i := 1; i < len(prices) ; i++ {
		if prices[i] < lowestPrice {
			lowestPrice = prices[i]
			continue
		}

		profit := prices[i] - lowestPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	
	return maxProfit
}
