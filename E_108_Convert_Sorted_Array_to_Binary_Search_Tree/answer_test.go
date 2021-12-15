package E_108_Convert_Sorted_Array_to_Binary_Search_Tree

import (
	"fmt"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	PrintTree(root)
}

func PrintTree(t *TreeNode) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, t)
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
		fmt.Printf("%v ", head.Val)
	}
}
