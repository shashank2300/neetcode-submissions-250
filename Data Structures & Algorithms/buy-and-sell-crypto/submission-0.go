func maxProfit(prices []int) int {
	ans := 0
	buy, sell := 0, 0
	for sell < len(prices) {
		if prices[sell] < prices[buy] {
			buy = sell
			sell++
			continue
		}
		ans = max(ans, prices[sell]-prices[buy])
		sell++
	}
	return ans
}
