package M_61_Rotate_List

import (
	"fmt"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	//l := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: &ListNode{
	//				Val:  4,
	//				Next: &ListNode{
	//					Val:  5,
	//					Next: &ListNode{
	//						Val:  6,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}

	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	ret := rotateRight(l, 2)
	for ret != nil {
		fmt.Printf("%d ", ret.Val)
		ret = ret.Next
	}
}
