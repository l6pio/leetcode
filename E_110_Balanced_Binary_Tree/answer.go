package E_110_Balanced_Binary_Tree

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode, level int) (int, bool) {
	if root == nil {
		return level, true
	}

	leftLevel, leftOK := IsBalanced(root.Left, level+1)
	rightLevel, rightOK := IsBalanced(root.Right, level+1)

	if leftOK && rightOK && leftLevel-rightLevel >= -1 && leftLevel-rightLevel <= 1 {
		if leftLevel > rightLevel {
			return leftLevel, true
		} else {
			return rightLevel, true
		}
	}
	return 0, false
}

func isBalanced(root *TreeNode) bool {
	_, OK := IsBalanced(root, 0)
	return OK
}
