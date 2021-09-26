package main

func sort(nums []int, start int, end int) []int {
	if end <= start {
		return nums
	}

	left := start + 1
	for i := start + 1; i <= end; i++ {
		if nums[i] <= nums[start] {
			tmp := nums[left]
			nums[left] = nums[i]
			nums[i] = tmp
			left += 1
		}
	}
	tmp := nums[start]
	nums[start] = nums[left-1]
	nums[left-1] = tmp
	nums = sort(nums, start, left-2)
	nums = sort(nums, left, end)
	return nums
}

func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	if numsLen <= 2 {
		return [][]int{}
	}
	nums = sort(nums, 0, len(nums)-1)

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
