package main

import "fmt"

// https://leetcode-cn.com/problems/next-permutation/

func reverse(nums []int, left int, right int) {
	for left < right {
		tmp := nums[left]
		nums[left] = nums[right]
		nums[right] = tmp
		left += 1
		right -= 1
	}
}

func nextPermutation(nums []int) {
	numsLen := len(nums)
	if numsLen < 2 {
		return
	}

	left := numsLen - 1
	for left > 0 {
		if nums[left] > nums[left-1] {
			break
		}
		left -= 1
	}

	if left > 0 {
		right := numsLen - 1
		for right > left-1 {
			if nums[right] > nums[left-1] {
				break
			}
			right -= 1
		}

		if right > left-1 {
			tmp := nums[left-1]
			nums[left-1] = nums[right]
			nums[right] = tmp
		}
	}

	reverse(nums, left, numsLen-1)
}

func main() {
	//ret := []int{1, 2, 3, 4, 5}
	//ret := []int{5, 4, 3, 2, 1}
	ret := []int{1, 5, 4, 3, 2}
	nextPermutation(ret)
	fmt.Printf("%v\n", ret)
}
