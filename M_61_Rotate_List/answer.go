package M_61_Rotate_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	root := &ListNode{
		Val:  0,
		Next: head,
	}

	p, q := head, root

	length := 0
	for p != nil {
		length += 1
		p = p.Next
	}

	if length == 0 {
		return head
	}

	k = k % length
	if k == 0 {
		return head
	}

	p = head
	for p.Next != nil {
		if k <= 1 {
			q = q.Next
		}
		k -= 1
		p = p.Next
	}

	if q != root {
		p.Next = root.Next
		root.Next = q.Next
		q.Next = nil
	}
	return root.Next
}
