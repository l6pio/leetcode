package main

func maxSubArray(nums []int) int {
	ret := nums[0]
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		if ret < tmp {
			ret = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	return ret
}

func main() {
	ret := maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})
	println(ret)
}
