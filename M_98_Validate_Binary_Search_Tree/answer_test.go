package M_98_Validate_Binary_Search_Tree

import (
	"testing"
)

func Test_isValidBST(t *testing.T) {
	//tree := &TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}

	tree := &TreeNode{
		Val: 2147483647,
	}
	ret := isValidBST(tree)
	print(ret, 1<<31-1)
}
