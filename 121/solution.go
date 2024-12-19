package solution

func maxProfit(prices []int) int {
	var profit int
	var price = prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < price {
			price = prices[i]
		} else if prices[i]-price > profit {
			profit = prices[i] - price
		}
	}

	return profit
}
