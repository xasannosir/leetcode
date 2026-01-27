package solution

func maxProfit(prices []int) int {
	total := 0
	n := len(prices)

	for i := 0; i < n-1; i++ {
		if prices[i] < prices[i+1] {
			total += (prices[i+1] - prices[i])
		}
	}

	return total
}
