package M_128_Longest_Consecutive_Sequence

//func longestConsecutive(nums []int) int {
//	m := make(map[int]int)
//	maxLen := 0
//	for _, n := range nums {
//		_, ok := m[n]
//		if ok {
//			continue
//		}
//
//		k := n
//		length, ok := m[n-1]
//		if ok {
//			m[n] = length + 1
//		} else {
//			m[n] = 1
//		}
//
//		for ; ; k++ {
//			if _, ok := m[k+1]; ok {
//				m[k+1] += m[n]
//			} else {
//				break
//			}
//		}
//
//		if m[k] > maxLen {
//			maxLen = m[k]
//		}
//	}
//	return maxLen
//}

/*
 * 改题的主要点在于连接左右两个"块"之后，只要确保"块头"和"块尾"的值是"正确的"即可。
 */
func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	maxLen := 0
	for _, n := range nums {
		_, ok := m[n]
		if ok {
			continue
		}

		left := m[n-1]
		right := m[n+1]
		m[n] = left + right + 1
		m[n-left] = m[n]
		m[n+right] = m[n]

		if m[n] > maxLen {
			maxLen = m[n]
		}
	}
	return maxLen
}
