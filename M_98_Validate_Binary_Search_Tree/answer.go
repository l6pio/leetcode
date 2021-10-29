package M_98_Validate_Binary_Search_Tree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return IsValidBST(root, -1<<63, 1<<63-1)
}

func IsValidBST(root *TreeNode, left int64, right int64) bool {
	if root == nil {
		return true
	}

	isRootValid := int64(root.Val) > left && int64(root.Val) < right
	isLeftValid := IsValidBST(root.Left, left, int64(math.Min(float64(right), float64(root.Val))))
	isRightValid := IsValidBST(root.Right, int64(math.Max(float64(left), float64(root.Val))), right)
	return isRootValid && isLeftValid && isRightValid
}
