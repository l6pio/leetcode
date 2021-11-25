package E_83_Remove_Duplicates_from_Sorted_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	root, p, q := head, head, head
	for q != nil {
		if q.Val != p.Val {
			p.Next = q
			p = p.Next
		}
		q = q.Next
	}
	if p != nil {
		p.Next = nil
	}
	return root
}
