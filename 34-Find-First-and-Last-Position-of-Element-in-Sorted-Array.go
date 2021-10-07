package main

import "fmt"

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func binarySearch(nums []int, left int, right int, target int, findRightMost bool) int {
	if left > right {
		return -1
	}

	if left == right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}

	middle := (left + right) / 2
	if target == nums[middle] {
		var i int
		if findRightMost {
			if i = binarySearch(nums, middle+1, right, target, findRightMost); i > -1 {
				return i
			}
			return middle
		} else {
			if i = binarySearch(nums, left, middle-1, target, findRightMost); i > -1 {
				return i
			}
			return middle
		}
	} else if target < nums[middle] {
		return binarySearch(nums, left, middle-1, target, findRightMost)
	} else {
		return binarySearch(nums, middle+1, right, target, findRightMost)
	}
}

func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	left := binarySearch(nums, 0, numsLen-1, target, false)
	right := binarySearch(nums, 0, numsLen-1, target, true)
	return []int{left, right}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	boundary := searchRange(nums, 10)
	fmt.Printf("%+v", boundary)
}
