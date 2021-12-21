package E_1_Two_Sum

// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, n := range nums {
		if idx, ok := m[n]; ok {
			return []int{idx, i}
		}
		m[target-n] = i
	}
	return []int{-1, -1}
}
