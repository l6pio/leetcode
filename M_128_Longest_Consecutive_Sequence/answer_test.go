package M_128_Longest_Consecutive_Sequence

import (
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	ret := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	//ret := longestConsecutive([]int{1, 2, 0, 1})
	println(ret)
}
