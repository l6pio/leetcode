package M_109_Convert_Sorted_List_to_Binary_Search_Tree

import (
	"fmt"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	list := &ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  9,
						Next: nil,
					},
				},
			},
		},
	}
	tree := sortedListToBST(list)
	PrintTree(tree)
}

func PrintTree(t *TreeNode) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, t, nil)
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		if head == nil {
			println("")
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		} else {
			tag := ""
			if head.Left != nil {
				queue = append(queue, head.Left)
				tag += "<"
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
				tag += ">"
			}
			fmt.Printf("%v%s   ", head.Val, tag)
		}
	}
}
