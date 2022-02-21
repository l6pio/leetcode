package M_134_Gas_Station

import (
	"testing"
)

func Test(t *testing.T) {
	//ret := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	ret := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	println(ret)
}
