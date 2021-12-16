package E_111_Minimum_Depth_of_Binary_Tree

import (
	"testing"
)

func Test_minDepth(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
		},
	}
	depth := minDepth(root)
	println(depth)
}
