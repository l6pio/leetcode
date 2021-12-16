package E_110_Balanced_Binary_Tree

import (
	"testing"
)

func Test_isBalanced(t *testing.T) {
	root := &TreeNode{
		Val:  3,
		Left: nil,
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	balanced := isBalanced(root)
	println(balanced)
}
