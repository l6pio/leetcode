package main

import "math"

// https://leetcode-cn.com/problems/container-with-most-water/

func maxArea(height []int) int {
	ret := 0
	left := 0
	right := len(height) - 1
	for left < right {
		ret = int(math.Max(float64(ret),
			float64((right-left)*int(math.Min(float64(height[left]), float64(height[right]))))))
		if height[left] <= height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return ret
}

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
