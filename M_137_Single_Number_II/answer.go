package M_137_Single_Number_II

func singleNumber(nums []int) int {
	numsLen := len(nums)

	// 要指定 int32 类型，因为需要规定符号位的位置，要不然会把负数表示成大的正数
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := 0
		for j := 0; j < numsLen; j++ {
			total += (nums[j] >> i) & 1
		}
		if total % 3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
