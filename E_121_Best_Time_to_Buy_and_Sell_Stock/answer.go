package E_121_Best_Time_to_Buy_and_Sell_Stock

func maxProfit(prices []int) int {
	min := prices[0]
	max := -1
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			max = min
		}
		if prices[i] > max {
			max = prices[i]
		}
		if max - min > profit {
			profit = max - min
		}
	}
	return profit
}
