package E_83_Remove_Duplicates_from_Sorted_List

import (
	"fmt"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	ret := deleteDuplicates(l)
	for ret != nil {
		fmt.Printf("%d ", ret.Val)
		ret = ret.Next
	}
}
