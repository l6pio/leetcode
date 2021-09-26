package utils

func Sort(nums []int, start int, end int) []int {
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
	nums = Sort(nums, start, left-2)
	nums = Sort(nums, left, end)
	return nums
}
