package utils

import "testing"

func Test_PrintTree(t *testing.T) {
	PrintTree(MakeTree([]int{2, 1, 3}))
}
