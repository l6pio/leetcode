package E_121_Best_Time_to_Buy_and_Sell_Stock

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	profit := maxProfit([]int{7, 1, 5, 3, 6, 4})
	//profit := maxProfit([]int{7, 6, 4, 3, 1})
	println(profit)
}
