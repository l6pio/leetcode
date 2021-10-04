package main

func Permute(nums []int, tmp []int, result [][]int) [][]int {
	numsLen := len(nums)
	if numsLen == 0 {
		return append(result, tmp)
	}

	for i := 0; i < numsLen; i++ {
		result = Permute(
			append(append([]int{}, nums[:i]...), nums[i+1:]...),
			append(append([]int{}, tmp...), nums[i]),
			result,
		)
	}
	return result
}

func permute(nums []int) [][]int {
	return Permute(nums, []int{}, [][]int{})
}

//func main() {
//	ret := permute([]int{1, 2, 3})
//	fmt.Printf("%v", ret)
//}
