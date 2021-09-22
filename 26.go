package main

func removeDuplicates(nums []int) int {
	numsLen := len(nums)

	if numsLen < 2 {
		return numsLen
	}

	k, i := 0, 0
	for ; i < numsLen-1; i++ {
		if nums[i] != nums[i+1] {
			nums[k] = nums[i]
			k += 1
		}
	}
	nums[k] = nums[i]
	return k + 1
}

//func main() {
//	nums := []int{3, 3, 5, 6, 7, 7, 7, 9}
//	ret := removeDuplicates(nums)
//	println(ret)
//}
