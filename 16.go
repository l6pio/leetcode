package main

import (
	"l6p.io/leetcode/utils"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	numsLen := len(nums)
	utils.Sort(nums, 0, numsLen-1)

	if numsLen < 3 {
		return 0
	}

	ret := nums[0] + nums[1] + nums[2]
	for i, n := range nums {
		left := i + 1
		right := numsLen - 1

		for left < right {
			sum := n + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			if sum > target {
				right -= 1
			} else {
				left += 1
			}
			if math.Abs(float64(sum)-float64(target)) < math.Abs(float64(ret)-float64(target)) {
				ret = sum
			}
		}
	}
	return ret
}

//func main() {
//	ret := threeSumClosest([]int{-1, 2, 1, -4}, 1)
//	println(ret)
//}
