package M_77_Combinations

func Combine(nums []int, k int, tmp *[]int, res *[][]int) {
	if k == 0 {
		*res = append(*res, *tmp)
		return
	}

	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		*tmp = append(*tmp, nums[i])
		Combine(append([]int{}, nums[i+1:]...), k-1, tmp, res)
		*tmp = append([]int{}, (*tmp)[:len(*tmp)-1]...)
	}
}

func combine(n int, k int) [][]int {
	nums := make([]int, 0)
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}

	tmp := make([]int, 0)
	res := make([][]int, 0)
	Combine(nums, k, &tmp, &res)
	return res
}
