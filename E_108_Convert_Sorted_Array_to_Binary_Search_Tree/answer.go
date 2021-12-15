package E_108_Convert_Sorted_Array_to_Binary_Search_Tree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int, root *TreeNode) {
	numsLen := len(nums)
	if numsLen == 0 {
		return
	}

	middle := numsLen / 2
	root.Val = nums[middle]

	leftNums := nums[:middle]
	if len(leftNums) > 0 {
		left := &TreeNode{}
		root.Left = left
		SortedArrayToBST(leftNums, root.Left)
	}

	rightNums := nums[middle+1:]
	if len(rightNums) > 0 {
		right := &TreeNode{}
		root.Right = right
		SortedArrayToBST(rightNums, root.Right)
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	root := &TreeNode{}
	SortedArrayToBST(nums, root)
	return root
}
