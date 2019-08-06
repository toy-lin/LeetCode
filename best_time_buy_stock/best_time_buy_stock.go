package best_time_buy_stock

func maxProfit(prices []int) int {
	var sum int
	var n = len(prices)
	var waitToSell bool
	var indexKeepStock int
	for i := 1; i < n; i++ {
		if prices[i-1] < prices[i] {
			if !waitToSell {
				indexKeepStock = i - 1
				waitToSell = true
			}
		} else if waitToSell {
			sum += prices[i-1] - prices[indexKeepStock]
			waitToSell = false
		}
	}
	if waitToSell {
		sum += prices[n-1] - prices[indexKeepStock]
	}
	return sum
}
