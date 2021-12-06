package M_107_Binary_Tree_Level_Order_Traversal_II

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	arr := []*TreeNode{root, nil}
	for idx := 0; idx < len(arr); idx++ {
		curr := arr[idx]
		if curr == nil {
			if arr[len(arr)-1] != nil {
				arr = append(arr, nil)
			}
		} else {
			if curr.Right != nil {
				arr = append(arr, curr.Right)
			}
			if curr.Left != nil {
				arr = append(arr, curr.Left)
			}
		}
	}

	var ret [][]int
	row := make([]int, 0)
	for idx := len(arr) - 1; idx >= 0; idx-- {
		if arr[idx] == nil {
			if len(row) > 0 {
				ret = append(ret, row)
				row = make([]int, 0)
			}
		} else {
			row = append(row, arr[idx].Val)
		}
	}
	return append(ret, row)
}
