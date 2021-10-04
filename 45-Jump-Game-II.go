package main

import "math"

//func jump(nums []int) int {
//	numsLen := len(nums)
//	minSteps := make([]int, numsLen)
//	minSteps[0] = 0
//	for i := 1; i < numsLen; i++ {
//		minSteps[i] = numsLen
//	}
//
//	for i := 1; i < numsLen; i++ {
//		minStep := numsLen
//		for j := 0; j < i; j++ {
//			if minSteps[j] < minStep && j+nums[j] >= i {
//				minStep = minSteps[j] + 1
//			}
//		}
//		minSteps[i] = minStep
//	}
//	return minSteps[numsLen-1]
//}

func jump(nums []int) int {
	step := 0
	end := 0
	maxPosition := 0
	for i := 0; i < len(nums) - 1; i++ {
		maxPosition = int(math.Max(float64(maxPosition), float64(i+nums[i])))
		if i == end {
			end = maxPosition
			step++
		}
	}
	return step
}

//func main() {
//	ret := jump([]int{2, 3, 1, 1, 4})
//	//ret := jump([]int{2, 3, 0, 1, 4})
//	println(ret)
//}
