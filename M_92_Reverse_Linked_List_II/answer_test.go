package M_92_Reverse_Linked_List_II

import (
	"fmt"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	//head := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val: 5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}
	//ret := reverseBetween(head, 2, 4)

	//head := &ListNode{
	//	Val:  1,
	//	Next: nil,
	//}
	//ret := reverseBetween(head, 1, 1)

	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}
	ret := reverseBetween(head, 1, 2)

	for {
		fmt.Printf("%v ", ret.Val)
		ret = ret.Next
		if ret == nil {
			break
		}
	}
}
