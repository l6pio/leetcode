package Q_90_Subsets_II

import "sort"

// https://leetcode-cn.com/problems/subsets-ii/

func SubsetsWithDup(nums []int, numsLen int, left int, mark []bool, tmp []int, res [][]int) [][]int {
	res = append(res, tmp)
	for i := left; i < numsLen; i++ {
		if i > 0 && nums[i] == nums[i-1] && !mark[i-1] {
			continue
		}
		mark[i] = true
		res = SubsetsWithDup(nums, numsLen, i+1, mark, append(append([]int{}, tmp...), nums[i]), res)
		mark[i] = false
	}
	return res
}

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	numsLen := len(nums)
	mark := make([]bool, numsLen)
	sort.Ints(nums)
	return SubsetsWithDup(nums, numsLen, 0, mark, []int{}, res)
}
