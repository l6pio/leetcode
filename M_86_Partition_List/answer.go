package M_86_Partition_List

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	cutPoint := &ListNode{
		Val:  -1,
		Next: head,
	}
	root := cutPoint

	for cutPoint.Next != nil && cutPoint.Next.Val < x {
		cutPoint = cutPoint.Next
	}

	left := cutPoint.Next
	if left == nil {
		return root.Next
	}

	right := left.Next
	for right != nil {
		if right.Val < x {
			left.Next = right.Next
			right.Next = cutPoint.Next
			cutPoint.Next = right
			cutPoint = cutPoint.Next
			right = left.Next
		} else {
			left = left.Next
			right = right.Next
		}
	}
	return root.Next
}
