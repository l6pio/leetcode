package Q_80_Remove_Duplicates_from_Sorted_Array_II

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	//arr := []int{1, 1, 1, 2, 2, 3}
	//arr := []int{1}
	//arr := []int{1, 1}
	//arr := []int{1, 1, 1}
	//arr := []int{1, 1, 1, 2, 3}
	newLen := removeDuplicates(arr)
	fmt.Printf("%+v, %d\n", arr, newLen)
}
