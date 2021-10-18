package Q_80_Remove_Duplicates_from_Sorted_Array_II

func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	i, j := 0, 0
	for ; j < numsLen; j++ {
		// 在最后一个元素之前判断"跳跃"点
		// 或者在最后一个元素的位置做判断
		if (j < numsLen-1 && nums[j] != nums[j+1]) || j == numsLen-1 {
			// 一定是先判断再更新"i"指向的元素
			if j > 0 && nums[j] == nums[j-1] {
				nums[i] = nums[j]
				nums[i+1] = nums[j]
				i += 2
			} else {
				nums[i] = nums[j]
				i++
			}
		}
	}
	return i
}
