package main

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

func Search(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2

	if nums[middle] == target {
		return middle
	}

	if left == right {
		return -1
	}

	if nums[left] < nums[middle] {
		if target >= nums[left] && target < nums[middle] {
			return Search(nums, left, middle-1, target)
		} else {
			return Search(nums, middle+1, right, target)
		}
	} else {
		if target >= nums[middle+1] && target <= nums[right] {
			return Search(nums, middle+1, right, target)
		} else {
			return Search(nums, left, middle-1, target)
		}
	}
}

func search(nums []int, target int) int {
	return Search(nums, 0, len(nums)-1, target)
}

func main() {
	nums := []int{5}
	ret := search(nums, 2)
	println(ret)
}
