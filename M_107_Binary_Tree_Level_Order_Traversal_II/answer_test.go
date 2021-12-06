package M_107_Binary_Tree_Level_Order_Traversal_II

import (
	"fmt"
	"testing"
)

func Test_levelOrderBottom(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	ret := levelOrderBottom(tree)
	fmt.Printf("%v", ret)
}
