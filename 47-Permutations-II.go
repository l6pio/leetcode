package main

import (
	"sort"
)

func PermuteUnique(nums []int, tmp []int, result [][]int) [][]int {
	numsLen := len(nums)
	if numsLen == 0 {
		return append(result, tmp)
	}

	for i := 0; i < numsLen; {
		result = PermuteUnique(
			append(append([]int{}, nums[:i]...), nums[i+1:]...),
			append(append([]int{}, tmp...), nums[i]),
			result,
		)
		i++
		for i < numsLen && nums[i] == nums[i-1] {
			i++
		}
	}
	return result
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return PermuteUnique(nums, []int{}, [][]int{})
}

//func main() {
//	ret := permuteUnique([]int{3, 3, 0, 3})
//	fmt.Printf("%v\n", ret)
//}
