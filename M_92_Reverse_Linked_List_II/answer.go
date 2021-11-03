package M_92_Reverse_Linked_List_II

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	root := &ListNode{
		Val:  0,
		Next: head,
	}
	head = root

	for i := 1; i < left; i++ {
		root = root.Next
	}

	tail := root.Next
	for i := 0; i < right-left; i++ {
		moveNode := tail.Next
		tail.Next = moveNode.Next
		moveNode.Next = root.Next
		root.Next = moveNode
	}
	return head.Next
}
