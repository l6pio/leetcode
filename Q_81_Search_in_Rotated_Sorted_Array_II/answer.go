package Q_81_Search_in_Rotated_Sorted_Array_II

func Search(nums []int, target int, left int, right int) bool {
	if left > right {
		return false
	}

	middle := (left + right) / 2
	if nums[middle] == target {
		return true
	} else {
		if nums[left] == nums[middle] {
			// 在left的值等于middle的值的时候，无法分辨在左还是在右，例如 [1, 0, 1, 1, 1] 中查找 0
			return Search(nums, target, left, middle-1) || Search(nums, target, middle+1, right)
		} else if nums[left] < nums[middle] {
			if target >= nums[left] && target < nums[middle] {
				return Search(nums, target, left, middle-1)
			} else {
				return Search(nums, target, middle+1, right)
			}
		} else {
			if target > nums[middle] && target <= nums[right] {
				return Search(nums, target, middle+1, right)
			} else {
				return Search(nums, target, left, middle-1)
			}
		}
	}
}

func search(nums []int, target int) bool {
	return Search(nums, target, 0, len(nums)-1)
}
