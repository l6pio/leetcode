package M_122_Best_Time_to_Buy_and_Sell_Stock_II

func maxProfit(prices []int) int {
	profit := 0
	min, max := prices[0], prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] >= prices[i-1] {
			max = prices[i]
		} else {
			if max > min {
				profit += max - min
			}
			min, max = prices[i], prices[i]
		}
	}
	if max > min {
		profit += max - min
	}
	return profit
}
