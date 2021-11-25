package E_100_Same_Tree

import (
	"testing"
)

func Test_isSameTree(t *testing.T) {
	tree1 := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 20,
		},
	}

	tree2 := &TreeNode{
		Val: 12,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 20,
		},
	}

	println(isSameTree(tree1, tree2))
}
