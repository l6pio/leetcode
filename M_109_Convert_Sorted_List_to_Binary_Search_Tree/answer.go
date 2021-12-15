package M_109_Convert_Sorted_List_to_Binary_Search_Tree

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedListToBST(head *ListNode, end *ListNode, root *TreeNode) {
	if head == end {
		return
	}

	middle, step2 := head, head
	for step2 != end && step2.Next != end {
		middle = middle.Next
		step2 = step2.Next.Next
	}

	root.Val = middle.Val
	if head != middle {
		leftNode := &TreeNode{}
		root.Left = leftNode
		SortedListToBST(head, middle, root.Left)
	}
	if middle.Next != end {
		rightNode := &TreeNode{}
		root.Right = rightNode
		SortedListToBST(middle.Next, end, root.Right)
	}
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	root := &TreeNode{}
	SortedListToBST(head, nil, root)
	return root
}
