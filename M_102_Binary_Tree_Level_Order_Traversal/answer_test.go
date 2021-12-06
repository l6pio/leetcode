package M_102_Binary_Tree_Level_Order_Traversal

import (
	"fmt"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	//tree := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 9,
	//	},
	//	Right: &TreeNode{
	//		Val: 20,
	//		Left: &TreeNode{
	//			Val: 15,
	//		},
	//		Right: &TreeNode{
	//			Val: 7,
	//		},
	//	},
	//}

	//[1,2,3,4,null,null,5]
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	res := levelOrder(tree)
	fmt.Printf("%v", res)
}
