package main

import "l6p.io/leetcode/utils"

func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	if numsLen <= 2 {
		return [][]int{}
	}
	nums = utils.Sort(nums, 0, len(nums)-1)

	var ret [][]int
	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}

		left := i + 1
		right := numsLen - 1
		for left < right {
			if n+nums[left]+nums[right] == 0 {
				ret = append(ret, []int{n, nums[left], nums[right]})
			}
			if n+nums[left]+nums[right] < 0 {
				left += 1
				for left < numsLen && nums[left-1] == nums[left] {
					left += 1
				}
			} else {
				right -= 1
				for right >= 0 && nums[right] == nums[right+1] {
					right -= 1
				}
			}
		}
	}
	return ret
}

//func main() {
//	ret := threeSum([]int{0, 0, 0})
//	fmt.Printf("%+v", ret)
//}
