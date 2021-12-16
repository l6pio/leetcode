package M_113_Path_Sum_II

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int, tmp []int, res *[][]int) {
	tmp = append(append([]int{}, tmp...), root.Val)

	if root.Left == nil && root.Right == nil {
		if targetSum-root.Val == 0 {
			*res = append(*res, tmp)
		}
	}

	if root.Left != nil {
		PathSum(root.Left, targetSum-root.Val, tmp, res)
	}

	if root.Right != nil {
		PathSum(root.Right, targetSum-root.Val, tmp, res)
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	PathSum(root, targetSum, []int{}, &res)
	return res
}
