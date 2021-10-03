package main

import (
	"sort"
)

func CombinationSum2(candidates []int, target int, result [][]int, tmp []int, left int, candidatesLen int) [][]int {
	if target == 0 {
		return append(result, tmp)
	}

	for left <= candidatesLen-1 {
		if target >= candidates[left] {
			result = CombinationSum2(
				candidates, target-candidates[left], result, append(tmp, candidates[left]),
				left+1, candidatesLen,
			)
		}
		left++
		for left <= candidatesLen-1 && candidates[left] == candidates[left-1] {
			left++
		}
	}
	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	candidatesLen := len(candidates)
	sort.Ints(candidates)
	return CombinationSum2(candidates, target, [][]int{}, []int{}, 0, candidatesLen)
}

//func main() {
//	ret := combinationSum2([]int{2, 2, 2}, 2)
//	fmt.Printf("%+v", ret)
//}
