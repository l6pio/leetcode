package main

func CombinationSum(candidates []int, target int, result [][]int, tmp []int) [][]int {
	if target == 0 {
		tmp2 := make([]int, len(tmp))
		copy(tmp2, tmp)
		return append(result, tmp2)
	}

	for i, c := range candidates {
		if target >= c {
			result = CombinationSum(candidates[i:], target-c, result, append(tmp, c))
		}
	}
	return result
}

func combinationSum(candidates []int, target int) [][]int {
	return CombinationSum(candidates, target, [][]int{}, []int{})
}

//func main() {
//	ret := combinationSum([]int{2, 7, 6, 3, 5, 1}, 9)
//	fmt.Printf("%+v", ret)
//}
