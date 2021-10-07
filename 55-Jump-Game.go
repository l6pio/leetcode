package main

import "math"

func canJump(nums []int) bool {
	maxPosition, end := 0, 0
	for i, n := range nums {
		maxPosition = int(math.Max(float64(maxPosition), float64(i+n)))
		if i == end {
			end = maxPosition
		}
	}
	return end >= len(nums)-1
}

func main() {
	println(canJump([]int{3, 2, 1, 0, 4}))
}
