package M_86_Partition_List

import (
	"fmt"
	"testing"
)

func Test_partition(t *testing.T) {
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 2,
	//				Next: &ListNode{
	//					Val: 5,
	//					Next: &ListNode{
	//						Val:  2,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	head := &ListNode{
		Val:  1,
		Next: nil,
	}

	count := 10
	head = partition(head, 3)
	for head != nil && count > 0 {
		fmt.Printf("%d ", head.Val)
		head = head.Next
		count -= 1
	}
}
