package maxProfit

func maxProfit(prices []int) int {
	buy := 0
	sell := 0
	profit := 0
	n := len(prices)
	for (buy < n && sell < n) {
		for (buy + 1 < n && prices[buy + 1] < prices[buy]) {
			buy++
		}
		sell = buy
		for (sell + 1 < n && prices[sell + 1] > prices[sell]) {
			sell++
		}

		profit += prices[sell] - prices[buy]
		buy = sell + 1
	}
	return profit

}