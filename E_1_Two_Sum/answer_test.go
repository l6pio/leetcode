package E_1_Two_Sum

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	ret := twoSum(nums, 18)
	fmt.Printf("%+v", ret)
}
