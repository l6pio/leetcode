package E_112_Path_Sum

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root.Left == nil && root.Right == nil {
		return targetSum-root.Val == 0
	} else if root.Left != nil && root.Right == nil {
		return HasPathSum(root.Left, targetSum-root.Val)
	} else if root.Left == nil && root.Right != nil {
		return HasPathSum(root.Right, targetSum-root.Val)
	} else {
		return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
	}
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return HasPathSum(root, targetSum)
}
