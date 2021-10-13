package main

import "fmt"

func sortColors(nums []int) {
	i, j, k := 0, 0, 0
	for n := 0; n < len(nums); n++ {
		if nums[n] == 0 {
			nums[k] = 2
			nums[j] = 1
			nums[i] = 0
			k++
			j++
			i++
		} else if nums[n] == 1 {
			nums[k] = 2
			nums[j] = 1
			k++
			j++
		} else if nums[n] == 2 {
			nums[k] = 2
			k++
		}
	}
}

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Printf("%+v\n", arr)
}
