package M_99_Recover_Binary_Search_Tree

import (
	"testing"
)

func Test_recoverTree(t *testing.T) {
	tree := &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 10,
		},
	}

	recoverTree(tree)
}
