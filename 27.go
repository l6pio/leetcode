package main

func removeElement(nums []int, val int) int {
	numsLen := len(nums)
	keep := 0
	scan := 0
	for scan < numsLen {
		if nums[scan] != val {
			nums[keep] = nums[scan]
			keep += 1
		}
		scan += 1
	}
	return keep
}

//func main() {
//	arr := []int{1, 2, 3}
//	ret := removeElement(arr, 2)
//	fmt.Printf("%v, %v", arr, ret)
//}
