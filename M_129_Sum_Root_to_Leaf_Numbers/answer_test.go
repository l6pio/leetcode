package M_129_Sum_Root_to_Leaf_Numbers

import (
	"testing"
)

func Test_sumNumbers(t *testing.T) {
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	sum := sumNumbers(root)
	println(sum)
}
