package E_111_Minimum_Depth_of_Binary_Tree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MinDepth(root *TreeNode, level int) int {
	if root.Left == nil && root.Right == nil {
		return level
	} else if root.Left != nil && root.Right == nil {
		return MinDepth(root.Left, level+1)
	} else if root.Left == nil && root.Right != nil {
		return MinDepth(root.Right, level+1)
	} else {
		leftLevel := MinDepth(root.Left, level+1)
		rightLevel := MinDepth(root.Right, level+1)
		if leftLevel < rightLevel {
			return leftLevel
		} else {
			return rightLevel
		}
	}
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return MinDepth(root, 1)
}
